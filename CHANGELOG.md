# CHANGELOG

所有对本项目及子模块的重要变更，均会记录在此文件中。

本文件遵循 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/) 规范，项目版本遵循 [语](https://semver.org/lang/zh-CN/)[义](https://semver.org/lang/zh-CN/)[化](https://semver.org/lang/zh-CN/)[版](https://semver.org/lang/zh-CN/)[本](https://semver.org/lang/zh-CN/)[ 2](https://semver.org/lang/zh-CN/)[.](https://semver.org/lang/zh-CN/)[0](https://semver.org/lang/zh-CN/)[.](https://semver.org/lang/zh-CN/)[0](https://semver.org/lang/zh-CN/)（格式：主版本号。次版本号。修订号）。

## [Unreleased] - 待发布

> 说明：记录当前开发中、尚未发布的变更，发布新版本时需将内容迁移至对应版本下。

### 主项目变更

* **Added**: 新增 `version` 子目录，提供一组用于版本管理的工具函数，包括获取当前版本、设置动态版本等。详情请参考 [version 子目录](./version/README.md)
* **Added**: 新增 `server` 子目录，提供一组用于启动 HTTP 服务器的工具函数，包括设置路由、处理请求、启动服务器等。详情请参考 [server 子目录](./server/README.md)
* **Added**: 新增 `errorsx` 子目录，提供一组用于处理错误的工具函数，包括自定义错误类型、错误转换、错误处理等。详情请参考 [errorsx 子目录](./errorsx/README.md)
* **Added**: 新增 `app` 子目录，提供一组用于应用程序初始化、配置加载、路由设置等的工具函数。详情请参考 [app 子目录](./app/README.md)
* **Added**: 新增 `core` 子目录，提供一组用于 Gin 框架的核心工具函数，包括路由设置、中间件设置、请求处理等。详情请参考 [core 子目录](./core/README.md)
* **Added**: 新增 `idx` 子目录，提供一组用于生成唯一索引的工具函数，包括基于 Snowflake 算法的索引生成器等。详情请参考 [idx 子目录](./idx/README.md)
* **Added**: 新增 `store` 子目录，提供一组用于操作数据库的工具函数，包括查询、插入、更新、删除等操作。详情请参考 [store 子目录](./store/README.md)
* **Added**: 新增 `validation` 子目录，提供一组用于验证数据的工具函数，包括结构体验证、字段验证、错误处理等。详情请参考 [validation 子目录](./validation/README.md)
* **Added**: 新增 `rid` 子目录，提供一组用于生成唯一资源标识符的工具函数，包括基于 Snowflake 算法的标识符生成器等。详情请参考 [rid 子目录](./rid/README.md)
* **Added**: 新增 `i18n` 子目录，提供一组用于国际化的工具函数，包括加载语言资源、翻译文本等。详情请参考 [i18n 子目录](./i18n/README.md)


### 子模块变更

none

## [v0.6.0] - 2025-10-16

### 主项目变更


* **Added**: 新增 `version` 子目录，提供一组用于版本管理的工具函数，包括获取当前版本、设置动态版本等。详情请参考 [version 子目录](./version/README.md)
* **Added**: 新增 `server` 子目录，提供一组用于启动 HTTP 服务器的工具函数，包括设置路由、处理请求、启动服务器等。详情请参考 [server 子目录](./server/README.md)
* **Added**: 新增 `errorsx` 子目录，提供一组用于处理错误的工具函数，包括自定义错误类型、错误转换、错误处理等。详情请参考 [errorsx 子目录](./errorsx/README.md)
* **Added**: 新增 `app` 子目录，提供一组用于应用程序初始化、配置加载、路由设置等的工具函数。详情请参考 [app 子目录](./app/README.md)
* **Added**: 新增 `core` 子目录，提供一组用于 Gin 框架的核心工具函数，包括路由设置、中间件设置、请求处理等。详情请参考 [core 子目录](./core/README.md)
* **Added**: 新增 `idx` 子目录，提供一组用于生成唯一索引的工具函数，包括基于 Snowflake 算法的索引生成器等。详情请参考 [idx 子目录](./idx/README.md)
* **Added**: 新增 `store` 子目录，提供一组用于操作数据库的工具函数，包括查询、插入、更新、删除等操作。详情请参考 [store 子目录](./store/README.md)
* **Added**: 新增 `validation` 子目录，提供一组用于验证数据的工具函数，包括结构体验证、字段验证、错误处理等。详情请参考 [validation 子目录](./validation/README.md)
* **Added**: 新增 `rid` 子目录，提供一组用于生成唯一资源标识符的工具函数，包括基于 Snowflake 算法的标识符生成器等。详情请参考 [rid 子目录](./rid/README.md)
* **Added**: 新增 `i18n` 子目录，提供一组用于国际化的工具函数，包括加载语言资源、翻译文本等。详情请参考 [i18n 子目录](./i18n/README.md)


### 子模块变更

none

## [v0.5.0] - 2025-10-16

### 主项目变更

* **Added**: 新增 `log` 子目录，提供一组用于日志记录的工具函数，包括日志级别、日志格式、日志输出等。详情请参考 [log 子目录](./log/README.md)
* **Added**：新增 `cache` 子目录，提供一组用于操作缓存的工具函数，包括设置缓存、获取缓存、删除缓存等。详情请参考 [cache 子目录](./cache/README.md)
* **Added**：新增 `gormx` 子目录, 提供一组用于操作 GORM 数据库的工具函数。详情请参考 [gormx 子目录](./gormx/README.md)
* **Added**：新增 `options` 子目录，提供一组用于配置数据库连接的选项函数，包括 MySQL、PostgreSQL 等数据库的连接选项。详情请参考 [options 子目录](./options/README.md)
* **Changed**：`ip` 子目录已替换为 `net` 子目录，提供一组用于获取出站 IP 地址的工具函数，包括使用默认 DNS 服务器和指定 DNS 服务器的方法。详情请参考 [net 子目录](./net/README.md)

### 子模块变更

none

## [v0.4.1] - 2025-10-15

### 主项目变更

none

### 子模块变更

#### entx

* **Fixed**：修复 `entx/query` 代码风格。

## [v0.4.0] - 2025-10-15

### 主项目变更

* **Added**：新增 `ip` 子目录，提供一组用于获取出站 IP 地址的工具函数，包括使用默认 DNS 服务器和指定 DNS 服务器的方法。详情请参考 [ip 子目录](./ip/README.md)
* **Added**：新增 `pagination` 子目录，提供一组用于分页查询的工具函数，包括构建分页查询参数、解析分页查询结果等。详情请参考 [pagination 子目录](./pagination/README.md)

### 子模块变更

#### entx

* **Added**：新增 `entx/query` 子目录，提供一组用于构建和执行 SQL 查询操作的工具函数，特别适用于处理 NULL 值和 JSON 字段的查询场景。详情请参考 [query 子目录](./entx/query/README.md)


## [v0.3.3] - 2025-10-15

### 主项目变更

无

### 子模块变更

#### id

* **Fixed**：修复 `GenerateOrderIdWithTenantId` 函数在商户 ID 不足 5 位时，右侧补零的问题。详情请参考 [id 子目录](./id/README.md)


## [v0.3.0] - 2025-10-15

### 主项目变更

* **Added**：新增 `trans` 子目录，提供一组通用的转换函数，包括字符串大小写转换、UUID 转换、时间转换等。详情请参考 [trans 子目录](./trans/README.md)

### 子模块变更

#### id

* **Added**：新增 `id` 子目录，提供一组用于生成和解析 UUID 的工具函数，包括生成随机 UUID、解析 UUID 字符串等。详情请参考 [id 子目录](./id/README.md)

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