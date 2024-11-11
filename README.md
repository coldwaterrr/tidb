<div align="center">
<a href='https://www.pingcap.com/?utm_source=github&utm_medium=tidb'>
<img src="docs/tidb-logo.png" alt="TiDB, a distributed SQL database" height=100></img>
</a>

---

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://github.com/pingcap/tidb/blob/master/LICENSE)
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![Build Status](https://prow.tidb.net/badge.svg?jobs=pingcap/tidb/merged_build)](https://prow.tidb.net/?repo=pingcap%2Ftidb&type=postsubmit&job=pingcap%2Ftidb%2Fmerged_build)
[![Go Report Card](https://goreportcard.com/badge/github.com/pingcap/tidb)](https://goreportcard.com/report/github.com/pingcap/tidb)
[![GitHub release](https://img.shields.io/github/tag/pingcap/tidb.svg?label=release)](https://github.com/pingcap/tidb/releases)
</div>

# TiDB

TiDB (/’taɪdiːbi:/, "Ti" 代表钛) 是一个开源的、云原生的、分布式 SQL 数据库，设计用于高可用性、水平和垂直可扩展性、强一致性和高性能。

- [关键特性](#关键特性)
- [快速开始](#快速开始)
- [需要帮助？](#需要帮助)
- [架构](#架构)
- [贡献](#贡献)
- [许可证](#许可证)
- [另见](#另见)
- [致谢](#致谢)

## 关键特性

- **[分布式事务](https://www.pingcap.com/blog/distributed-transactions-tidb?utm_source=github&utm_medium=tidb)**: TiDB 使用两阶段提交协议确保 ACID 合规性，提供强一致性。事务跨多个节点，TiDB 的分布式特性确保即使在网络分区或节点故障的情况下数据也能正确。

- **[水平和垂直可扩展性](https://docs.pingcap.com/tidb/stable/scale-tidb-using-tiup?utm_source=github&utm_medium=tidb)**: TiDB 可以通过增加节点水平扩展，或通过增加现有节点的资源垂直扩展，且无需停机。TiDB 的架构将计算与存储分离，使您可以根据需要独立调整两者以实现灵活性和增长。

- **[高可用性](https://docs.pingcap.com/tidbcloud/high-availability-with-multi-az?utm_source=github&utm_medium=tidb)**: 内置的 Raft 共识协议确保可靠性和自动故障转移。数据存储在多个副本中，事务只有在写入大多数副本后才会提交，保证强一致性和可用性，即使某些副本失败。可以配置副本的地理位置以实现不同的灾难容忍级别。

- **[混合事务/分析处理 (HTAP)](https://www.pingcap.com/blog/htap-demystified-defining-modern-data-architecture-tidb?utm_source=github&utm_medium=tidb)**: TiDB 提供两种存储引擎：TiKV（行存储引擎）和 TiFlash（列存储引擎）。TiFlash 使用 Multi-Raft Learner 协议实时从 TiKV 复制数据，确保 TiKV 行存储引擎和 TiFlash 列存储引擎之间的数据一致性。TiDB 服务器协调跨 TiKV 和 TiFlash 的查询执行以优化性能。

- **[云原生](https://www.pingcap.com/cloud-native?utm_source=github&utm_medium=tidb)**: TiDB 可以部署在公有云、本地或 Kubernetes 中。 [TiDB Operator](https://docs.pingcap.com/tidb-in-kubernetes/stable/tidb-operator-overview/?utm_source=github&utm_medium=tidb) 帮助在 Kubernetes 上管理 TiDB，自动化集群操作，而 [TiDB Cloud](https://tidbcloud.com/?utm_source=github&utm_medium=tidb) 提供了一个完全托管的服务，允许用户只需几次点击即可轻松经济地部署集群。

- **[MySQL 兼容性](https://docs.pingcap.com/tidb/stable/mysql-compatibility?utm_source=github&utm_medium=tidb)**: TiDB 兼容 MySQL 8.0，允许您使用熟悉的协议、框架和工具。您可以在不更改任何代码或仅进行最小修改的情况下将应用程序迁移到 TiDB。此外，TiDB 提供了一套 [数据迁移工具](https://docs.pingcap.com/tidb/stable/ecosystem-tool-user-guide?utm_source=github&utm_medium=tidb) 以帮助轻松将应用数据迁移到 TiDB。

- **[开源承诺](https://www.pingcap.com/blog/open-source-is-in-our-dna-reaffirming-tidb-commitment?utm_source=github&utm_medium=tidb)**: 开源是 TiDB 身份的核心。所有源代码都在 Apache 2.0 许可证下在 GitHub 上公开，包括企业级功能。TiDB 的构建理念是开源能够实现透明、创新和协作。我们积极鼓励社区的贡献，以帮助建立一个充满活力和包容性的生态系统，重申我们对开放开发和可访问性的承诺。

## 快速开始

> [!提示]  
> 作为我们对开源承诺的一部分，我们希望奖励所有 GitHub 用户。除了免费层，您还可以获得高达 $2000 的 TiDB Cloud Serverless 积分用于您的开源贡献 - [在此领取](https://ossinsight.io/open-source-heroes/?utm_source=ossinsight&utm_medium=referral&utm_campaign=plg_OSScontribution_credit_05)。

1. 启动 TiDB 集群

  - **在本地 Playground**。要启动本地测试集群，请参阅 [TiDB 快速入门指南](https://docs.pingcap.com/tidb/stable/quick-start-with-tidb#deploy-a-local-test-cluster?utm_source=github&utm_medium=tidb)。

  - **在 Kubernetes 上**。TiDB 可以轻松部署在自管理的 Kubernetes 环境或公有云上的 Kubernetes 服务中，使用 TiDB Operator。更多详情请参阅 [TiDB on Kubernetes 快速入门指南](https://docs.pingcap.com/tidb-in-kubernetes/stable/get-started?utm_source=github&utm_medium=tidb)。

  - **使用 TiDB Cloud（推荐）**。TiDB Cloud 提供了一个完全托管的 TiDB 版本，具有免费层，无需信用卡，因此您可以在几秒钟内获得一个免费集群并轻松开始：[注册 TiDB Cloud](https://tidbcloud.com/free-trial?utm_source=github&utm_medium=tidb)。

2. 了解 TiDB SQL：要探索 TiDB 的 SQL 功能，请参阅 [TiDB SQL 文档](https://docs.pingcap.com/tidb/stable/sql-statement-overview?utm_source=github&utm_medium=tidb)。

3. 使用 MySQL 驱动或 ORM [构建一个使用 TiDB 的应用](https://docs.pingcap.com/tidbcloud/dev-guide-overview?utm_source=github&utm_medium=tidb)。

4. 探索关键特性，如 [数据迁移](https://docs.pingcap.com/tidbcloud/tidb-cloud-migration-overview?utm_source=github&utm_medium=tidb)、[变更数据捕获](https://docs.pingcap.com/tidbcloud/changefeed-overview?utm_source=github&utm_medium=tidb)、[向量搜索](https://docs.pingcap.com/tidbcloud/vector-search-overview?utm_source=github&utm_medium=tidb)、[HTAP](https://docs.pingcap.com/tidbcloud/tidb-cloud-htap-quickstart?utm_source=github&utm_medium=tidb)、[灾难恢复](https://docs.pingcap.com/tidb/stable/dr-solution-introduction?utm_source=github&utm_medium=tidb) 等。

## 需要帮助？

- 您可以在我们的社区平台上与 TiDB 用户联系、提问、寻找答案并帮助他人：[Discord](https://discord.gg/KVRZBR2DrG?utm_source=github)、Slack（[英文](https://slack.tidb.io/invite?team=tidb-community&channel=everyone&ref=pingcap-tidb)、[日文](https://slack.tidb.io/invite?team=tidb-community&channel=tidb-japan&ref=github-tidb)）、[Stack Overflow](https://stackoverflow.com/questions/tagged/tidb)、TiDB 论坛（[英文](https://ask.pingcap.com/)、[中文](https://asktug.com)）、X [@PingCAP](https://twitter.com/PingCAP)

- 要提交错误报告、建议改进或请求新功能，请使用 [Github Issues](https://github.com/pingcap/tidb/issues) 或加入 [Github Discussions](https://github.com/orgs/pingcap/discussions) 讨论。

- 要排除 TiDB 故障，请参阅 [故障排除文档](https://docs.pingcap.com/tidb/stable/tidb-troubleshooting-map?utm_source=github&utm_medium=tidb)。

## 架构

![TiDB 架构](./docs/tidb-architecture.png)

在我们的 [文档](https://docs.pingcap.com/tidb/stable/tidb-architecture?utm_source=github&utm_medium=tidb) 中了解更多关于 TiDB 架构的详细信息。

## 贡献

TiDB 建立在对开源的承诺之上，我们欢迎每个人的贡献。无论您是对改进文档、修复错误还是开发新功能感兴趣，我们都邀请您共同塑造 TiDB 的未来。

- 请参阅我们的 [贡献者指南](https://github.com/pingcap/community/blob/master/contributors/README.md#how-to-contribute) 和 [TiDB 开发指南](https://pingcap.github.io/tidb-dev-guide/index.html) 以开始。

- 如果您正在寻找要处理的问题，可以查看 [good first issues](https://github.com/pingcap/tidb/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22) 或 [help wanted issues](https://github.com/pingcap/tidb/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)。

- [贡献地图](https://github.com/pingcap/tidb-map/blob/master/maps/contribution-map.md#a-map-that-guides-what-and-how-contributors-can-contribute) 列出了您可以贡献的所有内容。

- [社区仓库](https://github.com/pingcap/community) 包含您需要的其他所有内容。

- 不要忘记通过填写并提交此 [表单](https://forms.pingcap.com/f/tidb-contribution-swag) 领取您的贡献纪念品。

<a href="https://next.ossinsight.io/widgets/official/compose-recent-active-contributors?repo_id=41986369&limit=30" target="_blank" style="display: block" align="center">
  <picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://next.ossinsight.io/widgets/official/compose-recent-active-contributors/thumbnail.png?repo_id=41986369&limit=30&image_size=auto&color_scheme=dark" width="655" height="auto">
  <img alt="Active Contributors of pingcap/tidb - Last 28 days" src="https://next.ossinsight.io/widgets/official/compose-recent-active-contributors/thumbnail.png?repo_id=41986369&limit=30&image_size=auto&color_scheme=light" width="655" height="auto">
  </picture>
</a>

## 许可证

TiDB 采用 Apache 2.0 许可证。详情请参阅 [LICENSE](./LICENSE) 文件。

## 另见

- [TiDB 在线 Playground](https://play.tidbcloud.com/?utm_source=github&utm_medium=tidb_readme)
- TiDB 案例研究：[TiDB 客户](https://www.pingcap.com/customers/?utm_source=github&utm_medium=tidb)、[TiDB 事例記事](https://pingcap.co.jp/case-study/?utm_source=github&utm_medium=tidb)、[TiDB 中文用户案例](https://cn.pingcap.com/case/?utm_source=github&utm_medium=tidb)
- [TiDB 用户文档](https://docs.pingcap.com/tidb/stable?utm_source=github&utm_medium=tidb)
- [TiDB 设计文档](/docs/design)
- [TiDB 发布说明](https://docs.pingcap.com/tidb/dev/release-notes?utm_source=github&utm_medium=tidb)
- [TiDB 博客](https://www.pingcap.com/blog/?utm_source=github&utm_medium=tidb)
- [TiDB 路线图](roadmap.md)

## 致谢

- 感谢 [cznic](https://github.com/cznic) 提供一些很棒的开源工具。
- 感谢 [GolevelDB](https://github.com/syndtr/goleveldb)、[BoltDB](https://github.com/boltdb/bolt) 和 [RocksDB](https://github.com/facebook/rocksdb) 提供强大的存储引擎。
