package id

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/moweilong/mo/trans"
)

type idCounter uint32

func (c *idCounter) Increase() uint32 {
	cur := *c
	atomic.AddUint32((*uint32)(c), 1)
	atomic.CompareAndSwapUint32((*uint32)(c), 1000, 0)
	return uint32(cur)
}

var orderIdIndex idCounter

// GenerateOrderIdWithRandom 生成20位订单号，前缀 + 时间戳 + 随机数
func GenerateOrderIdWithRandom(prefix string, tm *time.Time) string {
	// 前缀 + 时间戳（14位） + 随机数（4位）

	if tm == nil {
		tm = trans.ToPtr(time.Now())
	}

	timestamp := tm.Format("20060102150405")

	randNum := rand.Intn(10000) // 生成0-9999之间的随机数

	// 使用%04d确保随机数部分为4位
	return fmt.Sprintf("%s%s%04d", prefix, timestamp, randNum)
}

// GenerateOrderIdWithIncreaseIndex 生成订单号，格式为：前缀+时间戳(14位)+自增长索引(1-3位)
// 注意：订单号总长度不固定，取决于前缀长度和索引值大小
func GenerateOrderIdWithIncreaseIndex(prefix string, tm *time.Time) string {
	if tm == nil {
		tm = trans.ToPtr(time.Now())
	}

	timestamp := tm.Format("20060102150405")

	index := orderIdIndex.Increase()

	return fmt.Sprintf("%s%s%d", prefix, timestamp, index)
}

// GenerateOrderIdWithTenantId 带商户ID的订单ID生成器：20250604123456M789012345678
// 格式：时间戳(14位) + 商户ID(固定5位) + 随机数(8位)
func GenerateOrderIdWithTenantId(tenantID string) string {
	tenantPart := tenantID
	if len(tenantPart) > 5 {
		tenantPart = tenantPart[:5]
	} else if len(tenantPart) < 5 {
		// 右侧补零
		padding := "00000"[:5-len(tenantPart)]
		tenantPart = tenantPart + padding
	}

	// 生成包含时间和随机数的订单号
	// 使用精确到纳秒的时间和更长的随机数来减少碰撞概率
	now := time.Now()
	timestamp := now.Format("20060102150405")
	// 结合纳秒部分和随机数，增加唯一性
	// 纳秒部分取前4位，随机数生成4位，总共8位随机部分
	nanoPart := fmt.Sprintf("%06d", now.UnixNano()%1000000)
	randomPart := fmt.Sprintf("%04d", rand.Intn(10000)) // 生成0-9999的随机数

	return timestamp + tenantPart + nanoPart[:4] + randomPart
}

func GenerateOrderIdWithPrefixSonyflake(prefix string) string {
	id, _ := NewSonyflakeID()
	return fmt.Sprintf("%s%d", prefix, id)
}

func GenerateOrderIdWithPrefixSnowflake(workerId int64, prefix string) string {
	id, _ := NewSnowflakeID(workerId)
	return fmt.Sprintf("%s%d", prefix, id)
}
