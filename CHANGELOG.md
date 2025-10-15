# CHANGELOG

所有对本项目及子模块的重要变更，均会记录在此文件中。

本文件遵循 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/) 规范，项目版本遵循 [语](https://semver.org/lang/zh-CN/)[义](https://semver.org/lang/zh-CN/)[化](https://semver.org/lang/zh-CN/)[版](https://semver.org/lang/zh-CN/)[本](https://semver.org/lang/zh-CN/)[ 2](https://semver.org/lang/zh-CN/)[.](https://semver.org/lang/zh-CN/)[0](https://semver.org/lang/zh-CN/)[.](https://semver.org/lang/zh-CN/)[0](https://semver.org/lang/zh-CN/)（格式：主版本号。次版本号。修订号）。

## [Unreleased] - 待发布

> 说明：记录当前开发中、尚未发布的变更，发布新版本时需将内容迁移至对应版本下。

### 主项目变更

* **Added**：新增 stringcase 子目录，提供字符串大小写转换功能：ToCamelCase(将字符串转换为驼峰式命名法)、ToPascalCase(将字符串转换为帕斯卡命名法)、ToSnakeCase(将字符串转换为蛇形命名法)、ToKebabCase(将字符串转换为烤肉串命名法)，详情请参考 [stringcase 子目录](./stringcase/README.md)


### 子模块变更

#### entx/client

* **Added**：升级依赖的 `grpc-client` 版本至 v1.6.2

## [v0.3.0] - 2025-10-15

### 主项目变更

* **Added**：none

### 子模块变更

#### entx

* **Added**：新增 `entx/mixin` 子目录，提供一组用于扩展 ent 实体的 mixin 实现，包括操作人字段、创建时间字段、更新时间字段、软删除字段等。详情请参考 [mixin 子目录](./entx/mixin/README.md)

## [v0.2.0] - 2025-10-14

### 主项目变更

* **Added**：none

### 子模块变更

#### entx

* **Added**：新增 `entx/update` 子目录，提供一组用于构建和执行 SQL 更新操作的工具函数，特别适用于处理 NULL 值和 JSON 字段的更新场景。详情请参考 [update 子目录](./entx/update/README.md)

## [v0.1.0] - 2025-10-14

> 说明：次版本（首次引入子模块），对应标签：v0.1.0

### 主项目变更

* **Added**：新增stringcase和fieldmaskutil包，提供字符串命名法转换功能（驼峰式、蛇形、烤肉串等），以及用于处理 Protocol Buffers 字段掩码（FieldMask）的工具包。



### 子模块变更

#### entx（首次引入）


* **Added**：新增 `entx` 子目录，提供 ent 客户端封装连接数据库


## [v0.0.1] - 2025-10-13

> 说明：初始版本（无副标题模块），对应标签：v0.0.1



* **Added**：byteutil 子目录，提供字节和整数转换以及字符大小写转换功能：IntToBytes(将int转为[]byte)、BytesToInt(将[]byte转为int)、ByteToLower(字节转小写)、ByteToUpper(字节转大写)

* **Fixed**：none

* **Changed**：none




## 编写规则说明



1. **版本排序**：按时间倒序排列，最新版本在最上方；`[Unreleased]` 固定在最顶部。

2. **版本标题格式**：`[版本号] - 发布日期`，若未发布则省略日期（如 `[Unreleased]`）。

3. **变更类型**：仅使用 6 种固定类型，且首字母大写（`Added`/`Changed`/`Fixed`/`Deprecated`/`Removed`/`Security`），避免自定义类型导致混乱。

4. **子模块记录**：

* 子模块需按「目录层级」拆分（如 `entx/model`/`entx/client`），使用 `####` 二级标题区分。

* 每个子模块的变更记录，需明确关联具体功能（避免模糊描述，例：不说 “优化模型”，而说 “优化 `User` 模型的查询效率”）。

1. **链接配置**：版本标题后可添加对比链接（格式：`[版本号](``https://github.com/``用户名/仓库名/compare/上一版本号...当前版本号)`），方便查看版本间差异。

2. **兼容性说明**：若变更涉及不兼容（主版本号升级），需在版本标题下添加 `> 注意：本版本不兼容 vX.X.X 及以下版本，升级前需参考《升级指南》` 提示。