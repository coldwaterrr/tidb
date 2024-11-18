# 集成测试

此文件夹包含所有依赖于外部进程（如 TiDB）的测试。

## 准备工作

1. 必须将以下可执行文件复制或链接到这些位置：

    * `bin/tidb-server`
    * `bin/tikv-server`
    * `bin/pd-server`
    * `bin/tiflash`
    * `bin/minio`
    * `bin/mc`

    版本必须 ≥2.1.0。

    **只有部分测试需要 `minio`/`mc`**，可以从官方网站下载，如果不需要运行这些测试，可以跳过。

    你可以使用 `tiup` 下载与 TiDB 集群相关的二进制文件，然后将它们链接到 `bin` 目录：
    ```shell
    cluster_version=v8.1.0 # 更改为你需要的版本
    tiup install tidb:$cluster_version tikv:$cluster_version pd:$cluster_version tiflash:$cluster_version
    ln -s ~/.tiup/components/tidb/$cluster_version/tidb-server bin/tidb-server
    ln -s ~/.tiup/components/tikv/$cluster_version/tikv-server bin/tikv-server
    ln -s ~/.tiup/components/pd/$cluster_version/pd-server bin/pd-server
    ln -s ~/.tiup/components/tiflash/$cluster_version/tiflash/tiflash bin/tiflash
    ```

2. `make build_for_lightning_integration_test`
   
    如果你的测试需要最新的 TiDB 服务器，运行 `make server` 来构建。

3. 必须安装以下程序：

    * `mysql`（CLI 客户端）
    * `curl`
    * `openssl`
    * `wget`
    * `lsof`

4. 执行测试的用户必须有权限创建文件夹 `/tmp/lightning_test`。所有测试工件将写入此文件夹。

## 运行

运行 `make lightning_integration_test` 来执行所有集成测试。
- 日志将写入 `/tmp/lightning_test` 目录。

运行 `tests/run.sh --debug` 在所有服务器启动后立即暂停。

如果你只想运行某些测试，可以使用：
```shell
TEST_NAME="lightning_gcs lightning_view" lightning/tests/run.sh
```

测试用例名称用空格分隔。

## 编写新测试

1. 新的集成测试可以作为 shell 脚本编写在 `tests/TEST_NAME/run.sh` 中。
    - `TEST_NAME` 应以 `lightning_` 开头。
    - 脚本在失败时应以非零错误代码退出。
2. 将 TEST_NAME 添加到 [run_group_lightning_tests.sh](./run_group_lightning_tests.sh) 中的现有组（推荐），或为其添加一个新组。
3. 如果你添加了一个新组，必须将新组的名称添加到 CI [lightning-integration-test](https://github.com/PingCAP-QE/ci/blob/main/pipelines/pingcap/tidb/latest/pull_lightning_integration_test.groovy)。

在 [utils](../../tests/_utils/) 中提供了几个方便的命令：

* `run_sql <SQL>` — 在 TiDB 数据库上执行 SQL 查询
* `run_lightning [CONFIG]` — 使用 `tests/TEST_NAME/CONFIG.toml` 启动 `tidb-lightning`
* `check_contains <TEXT>` — 检查之前的 `run_sql` 结果是否包含给定文本（以 `-E` 格式）
* `check_not_contains <TEXT>` — 检查之前的 `run_sql` 结果是否不包含给定文本（以 `-E` 格式）
