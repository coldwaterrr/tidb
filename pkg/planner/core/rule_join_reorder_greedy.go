// Copyright 2018 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"cmp"
	"math"
	"slices"

	"github.com/pingcap/tidb/pkg/expression"
	"github.com/pingcap/tidb/pkg/planner/core/base"
	"github.com/pingcap/tidb/pkg/util/intest"
)

type joinReorderGreedySolver struct {
	*baseSingleGroupJoinOrderSolver
}

// solve 基于贪心算法重新排序组中的连接节点。
//
// 对于组中与当前连接树具有连接等值条件的每个节点，计算该节点与连接树的累积连接成本，选择累积成本最小的节点与当前连接树连接。
//
// 累积连接成本 = CumCount(lhs) + CumCount(rhs) + RowCount(join)
//
//	对于基础节点，其 CumCount 等于其子树计数的总和。
//	有关更多详细信息，请参见 baseNodeCumCost。
//
// TODO: 这个公式将来可以改为实际的物理成本。
//
// 对于没有连接等值条件来连接它们的节点和连接树，我们最终会生成一个灌木连接树来进行笛卡尔连接。
func (s *joinReorderGreedySolver) solve(joinNodePlans []base.LogicalPlan, tracer *joinReorderTrace) (base.LogicalPlan, error) {
	var err error
	s.curJoinGroup, err = s.generateJoinOrderNode(joinNodePlans, tracer)
	if err != nil {
		return nil, err
	}
	var leadingJoinNodes []*jrNode
	if s.leadingJoinGroup != nil {
		// 我们有一个 leading hint 来让一些表先连接。结果存储在 s.leadingJoinGroup 中。
		// 我们为它单独生成 jrNode。
		leadingJoinNodes, err = s.generateJoinOrderNode([]base.LogicalPlan{s.leadingJoinGroup}, tracer)
		if err != nil {
			return nil, err
		}
	}
	// Sort plans by cost
	slices.SortStableFunc(s.curJoinGroup, func(i, j *jrNode) int {
		return cmp.Compare(i.cumCost, j.cumCost)
	})

	// joinNodeNum indicates the number of join nodes except leading join nodes in the current join group
	joinNodeNum := len(s.curJoinGroup)
	if leadingJoinNodes != nil {
		// The leadingJoinNodes should be the first element in the s.curJoinGroup.
		// So it can be joined first.
		leadingJoinNodes := append(leadingJoinNodes, s.curJoinGroup...)
		s.curJoinGroup = leadingJoinNodes
	}
	var cartesianGroup []base.LogicalPlan
	for len(s.curJoinGroup) > 0 {
		newNode, err := s.constructConnectedJoinTree(tracer)
		if err != nil {
			return nil, err
		}
		if joinNodeNum > 0 && len(s.curJoinGroup) == joinNodeNum {
			// 到这里意味着 leading hint 中使用的表与其他表之间没有连接条件
			// 例如: select /*+ leading(t3) */ * from t1 join t2 on t1.a=t2.a cross join t3
			// 我们不能让表 t3 先连接。
			// TODO(hawkingrei): 我们在 TestHint 中发现了这个问题。
			// 	`select * from t1, t2, t3 union all select /*+ leading(t3, t2) */ * from t1, t2, t3 union all select * from t1, t2, t3`
			//  这个 SQL 不应该返回警告，但它不会影响结果，所以我们会尽快修复它。
			s.ctx.GetSessionVars().StmtCtx.SetHintWarning("leading hint 无效，请检查 leading hint 表是否与其他表有连接条件")
		}
		cartesianGroup = append(cartesianGroup, newNode.p)
	}

	return s.makeBushyJoin(cartesianGroup), nil
}

func (s *joinReorderGreedySolver) constructConnectedJoinTree(tracer *joinReorderTrace) (*jrNode, error) {
	curJoinTree := s.curJoinGroup[0]
	s.curJoinGroup = s.curJoinGroup[1:]
	for {
		bestCost := math.MaxFloat64
		bestIdx, whateverValidOneIdx := -1, -1
		var finalRemainOthers, remainOthersOfWhateverValidOne []expression.Expression
		var bestJoin, whateverValidOne base.LogicalPlan
		for i, node := range s.curJoinGroup {
			newJoin, remainOthers := s.checkConnectionAndMakeJoin(curJoinTree.p, node.p)
			if newJoin == nil {
				continue
			}
			_, err := newJoin.RecursiveDeriveStats(nil)
			if err != nil {
				return nil, err
			}
			whateverValidOne = newJoin
			whateverValidOneIdx = i
			remainOthersOfWhateverValidOne = remainOthers
			curCost := s.calcJoinCumCost(newJoin, curJoinTree, node)
			tracer.appendLogicalJoinCost(newJoin, curCost)
			if bestCost > curCost {
				bestCost = curCost
				bestJoin = newJoin
				bestIdx = i
				finalRemainOthers = remainOthers
			}
		}
		// If we could find more join node, meaning that the sub connected graph have been totally explored.
		if bestJoin == nil {
			if whateverValidOne == nil {
				break
			}
			// This branch is for the unexpected case.
			// We throw assertion in test env. And create a valid join to avoid wrong result in the production env.
			intest.Assert(false, "Join reorder should find one valid join but failed.")
			bestJoin = whateverValidOne
			bestCost = math.MaxFloat64
			bestIdx = whateverValidOneIdx
			finalRemainOthers = remainOthersOfWhateverValidOne
		}
		curJoinTree = &jrNode{
			p:       bestJoin,
			cumCost: bestCost,
		}
		s.curJoinGroup = append(s.curJoinGroup[:bestIdx], s.curJoinGroup[bestIdx+1:]...)
		s.otherConds = finalRemainOthers
	}
	return curJoinTree, nil
}

func (s *joinReorderGreedySolver) checkConnectionAndMakeJoin(leftPlan, rightPlan base.LogicalPlan) (base.LogicalPlan, []expression.Expression) {
	leftPlan, rightPlan, usedEdges, joinType := s.checkConnection(leftPlan, rightPlan)
	if len(usedEdges) == 0 {
		return nil, nil
	}
	return s.makeJoin(leftPlan, rightPlan, usedEdges, joinType)
}
