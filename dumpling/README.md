🥟 Dumpling
============

[![Build Status](https://travis-ci.org/pingcap/dumpling.svg?branch=master)](https://travis-ci.org/pingcap/dumpling)
[![codecov](https://codecov.io/gh/pingcap/dumpling/branch/master/graph/badge.svg)](https://codecov.io/gh/pingcap/dumpling)
[![API Docs](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/pingcap/dumpling)
[![Go Report Card](https://goreportcard.com/badge/github.com/pingcap/dumpling)](https://goreportcard.com/report/github.com/pingcap/dumpling)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fpingcap%2Fdumpling.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fpingcap%2Fdumpling?ref=badge_shield)

**Dumpling** 是一个工具和 Go 库，用于从 MySQL 兼容数据库创建 SQL 转储。
它旨在替代 `mysqldump` 和 `mydumper`，特别是针对 TiDB。

您可以阅读[设计文档](https://github.com/pingcap/community/blob/master/rfc/2019-12-06-dumpling.md)、[英文用户指南](docs/en/user-guide.md) 和 [中文使用手册](docs/cn/user-guide.md) 了解详细信息。

功能
--------

> Dumpling 目前处于早期开发阶段，大多数功能尚未完成。欢迎贡献！

- [x] SQL 转储被拆分成多个文件（类似 `mydumper`），便于管理。
- [x] 并行导出多个表以加快执行速度。
- [x] 多种输出格式：SQL、CSV 等。
- [ ] 原生写入云存储（S3、GCS）
- [x] 高级表过滤

有任何问题？让我们在 [TiDB Internals 论坛](https://internals.tidb.io/) 讨论！

构建
--------

0. 在 `tidb` 目录下
1. 安装 Go 1.16 或更高版本
2. 运行 `make build_dumpling` 进行编译。输出在 `bin/dumpling`。
3. 运行 `make dumpling_unit_test` 进行单元测试。
4. 运行 `make dumpling_integration_test` 进行集成测试。对于集成测试：
  - 以下可执行文件必须复制、生成或链接到这些位置：
    * `bin/sync_diff_inspector`（从 [tidb-enterprise-tools-latest-linux-amd64](http://download.pingcap.org/tidb-enterprise-tools-latest-linux-amd64.tar.gz) 下载）
    * `bin/tidb-server`（从 [tidb-master-linux-amd64](https://download.pingcap.org/tidb-master-linux-amd64.tar.gz) 下载）
    * `bin/tidb-lightning`（从 [tidb-toolkit-latest-linux-amd64](https://download.pingcap.org/tidb-toolkit-latest-linux-amd64.tar.gz) 下载）
    * `bin/minio`（从 <https://min.io/download> 下载）
    * 现在，您可以运行 `sh ./dumpling/install.sh` 获取上述二进制文件。
  - 必须安装以下程序：
    * `mysql`（CLI 客户端）
  - 必须有一个本地 mysql 服务器监听 `127.0.0.1:3306`，并且有一个无密码的活动用户可以通过此 TCP 地址连接。

许可证
-------

Dumpling 使用 Apache 2.0 许可证。详情请参见 [LICENSE](./LICENSE) 文件。

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fpingcap%2Fdumpling.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fpingcap%2Fdumpling?ref=badge_large)
