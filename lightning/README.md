# TiDB-Lightning

**TiDB-Lightning** 是一个用于将 TB 级数据导入 TiDB 集群的工具。

## 文档

[中文文档](https://docs.pingcap.com/zh/tidb/stable/tidb-lightning-overview)

[英文文档](https://docs.pingcap.com/tidb/stable/tidb-lightning-overview)

[Import Into 语句](https://docs.pingcap.com/tidbcloud/sql-statement-import-into) 用于通过 TiDB Lightning 的物理导入模式将数据文件或 select 语句的结果导入 TiDB 中的空表。

## 构建

构建二进制文件：

```bash
$ cd ../tidb
$ make build_lightning
```

注意 TiDB-Lightning 支持使用 Go 版本 `Go >= 1.16` 进行构建。

当 TiDB-Lightning 构建成功后，你可以在 `bin` 目录中找到二进制文件。

## 运行测试

请参阅[此文档](../lightning/tests/README.md)了解如何运行集成测试。

## 快速开始

请参阅[TiDB Lightning 快速开始](https://docs.pingcap.com/tidb/stable/get-started-with-tidb-lightning)。

## 贡献

欢迎并非常感谢您的贡献。有关提交补丁和贡献工作流程的详细信息，请参阅[贡献指南](../CONTRIBUTING.md)。

## 许可证

TiDB-Lightning 使用 Apache 2.0 许可证。有关详细信息，请参阅[许可证](../LICENSE)文件。
