# id 包

id 包提供了各种 ID 生成和处理工具，包括 UUID、Snowflake、Sonyflake 和订单 ID 等多种 ID 生成方案，可以满足不同场景下的唯一标识符需求。

## 功能特性

- **UUID 系列**：支持标准 UUIDv4、短 UUID、KSUID、XID 和 MongoDB ObjectID
- **分布式 ID**：支持 Snowflake 和 Sonyflake 分布式 ID 生成算法
- **订单 ID**：提供多种订单号生成方式，包括随机数、自增索引、商户 ID 等方案
- **线程安全**：所有 ID 生成器均支持并发环境使用

## 安装方法

```bash
go get github.com/moweilong/mo/id
```

## 使用示例

### UUID 相关功能

```go
import "github.com/moweilong/mo/id"

// 生成带连字符的 UUIDv4
guid := id.NewGUIDv4(true) // 输出类似: 123e4567-e89b-12d3-a456-426614174000

// 生成不带连字符的 UUIDv4
guidNoHyphen := id.NewGUIDv4(false) // 输出类似: 123e4567e89b12d3a456426614174000

// 生成短 UUID
shortUUID := id.NewShortUUID() // 输出类似: 2spJN8kCx35Kp2b165YQkM

// 生成 KSUID
ksuid := id.NewKSUID() // 输出类似: 23Y7c9Hf5V9fK9XvDnP79V7pJ5b

// 生成 XID
xid := id.NewXID() // 输出类似: c8u6n03z6r1234567890

// 生成 MongoDB ObjectID
mongoId := id.NewMongoObjectID() // 输出类似: 648a3f73e1a8b1234567890a
```

### Snowflake 分布式 ID

```go
import "github.com/moweilong/mo/id"

// 直接生成 Snowflake ID
id, err := id.NewSnowflakeID(1) // workerId 为 1
if err != nil {
    // 处理错误
}

// 忽略错误的简便方法
id := id.GenerateSnowflakeID(1)

// 使用 SnowflakeNode 结构体
node, err := id.NewSnowflakeNode(1)
if err != nil {
    // 处理错误
}
idInt64 := node.Generate() // 获取 int64 类型的 ID
idStr := node.GenerateString() // 获取字符串类型的 ID
```

### Sonyflake 分布式 ID

```go
import "github.com/moweilong/mo/id"

// 直接生成 Sonyflake ID
id, err := id.NewSonyflakeID()
if err != nil {
    // 处理错误
}

// 忽略错误的简便方法
id := id.GenerateSonyflakeID()
```

### 订单 ID 生成

```go
import (
    "github.com/moweilong/mo/id"
    "time"
)

// 生成带前缀和随机数的订单号
orderId := id.GenerateOrderIdWithRandom("ORD", nil) // 输出类似: ORD202506041234567890

// 自定义时间
customTime := time.Date(2025, 6, 4, 12, 34, 56, 0, time.UTC)
orderId := id.GenerateOrderIdWithRandom("ORD", &customTime)

// 生成带自增长索引的订单号
orderId := id.GenerateOrderIdWithIncreaseIndex("ORD", nil)

// 生成带商户 ID 的订单号
orderId := id.GenerateOrderIdWithTenantId("MERCH") // 输出类似: 20250604123456MERCH7890

// 生成带前缀的 Sonyflake 订单号
orderId := id.GenerateOrderIdWithPrefixSonyflake("ORD-")

// 生成带前缀的 Snowflake 订单号
orderId := id.GenerateOrderIdWithPrefixSnowflake(1, "ORD-")
```

## 订单ID格式参考

- 电商平台：202506041234567890（时间戳 + 随机数，19-20 位）。
- 支付系统：PAY20250604123456789（业务前缀 + 时间戳 + 序号）。
- 微信支付：1589123456789012345（类似 Snowflake 的纯数字 ID）。
- 美团订单：202506041234567890123（时间戳 + 商户 ID + 随机数）。

## ID 类型比较

| 特性    | GUID/UUID    | KSUID      | ShortUUID | XID      | Snowflake      |
|-------|--------------|------------|-----------|----------|----------------|
| 长度    | 36/32字符（不含-） | 27字符       | 22字符      | 20字符     | 19（数字位数）       |
| 有序性   | 无序（UUIDv4）   | 严格时序       | 无序        | 趋势有序     | 严格时序           |
| 时间精度  | 无（UUIDv4）    | 毫秒级        | 无         | 秒级       | 毫秒级            |
| 分布式安全 | 高（随机数）       | 高          | 高         | 高        | 高（需配置WorkerID） |
| 性能    | 中等           | 中等         | 较低（编码开销）  | 极高       | 极高             |
| 时钟依赖  | 无            | 有（需处理时钟回拨） | 无         | 有（但影响较小） | 强依赖（需严格同步）     |
| 适用场景  | 跨系统兼容        | 时序索引       | 短ID、URL   | 高并发、短ID  | 分布式时序ID        |

## 选择建议

- **GUID/UUID**: 适用于需要跨系统兼容的场景，特别是当不需要有序性时。
- **KSUID**: 适合需要严格时序的应用，如事件日志、时间序列数据。
- **ShortUUID**: 当需要短ID且不关心有序性时的理想选择，适用于URL、短链接等。
- **XID**: 高并发场景下的短ID选择，适合需要一定有序性的应用。
- **Snowflake**: 适合分布式系统，特别是需要严格时序和高性能的场景，如大规模分布式应用。
- **Sonyflake**: Snowflake的改进版，更适用于大规模分布式系统。
- **订单ID**: 特定业务场景下的ID生成方案，根据具体需求选择合适的方法。

## API 参考

### UUID 相关函数

- **NewGUIDv4(withHyphen bool) string**: 生成UUIDv4，参数控制是否包含连字符
- **NewShortUUID() string**: 生成短UUID
- **NewKSUID() string**: 生成KSUID
- **NewXID() string**: 生成XID
- **NewMongoObjectID() string**: 生成MongoDB ObjectID

### Snowflake 相关函数和结构体

- **NewSnowflakeNode(workerId int64) (*SnowflakeNode, error)**: 创建新的Snowflake节点
- **NewSnowflakeID(workerId int64) (int64, error)**: 生成Snowflake ID
- **GenerateSnowflakeID(workerId int64) int64**: 生成Snowflake ID（忽略错误）
- **SnowflakeNode.Generate() int64**: 生成int64类型的Snowflake ID
- **SnowflakeNode.GenerateString() string**: 生成字符串类型的Snowflake ID

### Sonyflake 相关函数

- **NewSonyflakeID() (uint64, error)**: 生成Sonyflake ID
- **GenerateSonyflakeID() uint64**: 生成Sonyflake ID（忽略错误）

### 订单ID 相关函数

- **GenerateOrderIdWithRandom(prefix string, tm *time.Time) string**: 生成带随机数的订单号
- **GenerateOrderIdWithIncreaseIndex(prefix string, tm *time.Time) string**: 生成带自增长索引的订单号
- **GenerateOrderIdWithTenantId(tenantID string) string**: 生成带商户ID的订单号
- **GenerateOrderIdWithPrefixSonyflake(prefix string) string**: 生成带前缀的Sonyflake订单号
- **GenerateOrderIdWithPrefixSnowflake(workerId int64, prefix string) string**: 生成带前缀的Snowflake订单号

## 实现原理

### Snowflake ID 结构

64位ID = 41位时间戳 + 10位工作节点ID + 12位序列号

### Sonyflake ID 结构

64位ID = 39位时间戳 + 8位机器ID + 16位序列号

### 订单ID 结构

- 随机数订单ID：前缀 + 14位时间戳 + 4位随机数
- 自增索引订单ID：前缀 + 14位时间戳 + 自增索引
- 商户ID订单ID：14位时间戳 + 5位商户ID + 4位随机数

## 版本要求

- Go 1.16+（推荐 Go 1.18+ 以获得更好的性能和稳定性）

## 依赖项

- github.com/google/uuid
- github.com/lithammer/shortuuid/v4
- github.com/rs/xid
- github.com/segmentio/ksuid
- go.mongodb.org/mongo-driver/bson/primitive
- github.com/bwmarrin/snowflake
- github.com/sony/sonyflake
- github.com/tx7do/go-utils/trans
