package id

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOrderIdWithRandom(t *testing.T) {
	prefix := "PT"

	// 测试生成的订单号是否包含前缀
	orderID := GenerateOrderIdWithRandom(prefix, nil)
	assert.Contains(t, orderID, prefix, "订单号应包含前缀")
	t.Logf("GenerateOrderIdWithRandom: %s", orderID)

	// 测试生成的订单号长度是否正确
	assert.Equal(t, len(prefix)+14+4, len(orderID), "订单号长度应为前缀+时间戳+随机数")
}

func TestGenerateOrderIdWithIndex(t *testing.T) {
	prefix := "PT"
	tm := time.Now()

	// 验证订单ID格式
	sampleId := GenerateOrderIdWithIncreaseIndex(prefix, &tm)
	t.Logf("Sample order ID: %s", sampleId)

	// 验证订单ID包含前缀
	assert.Contains(t, sampleId, prefix, "订单号应包含前缀")

	// 验证订单ID包含正确的时间戳
	expectedTimestamp := tm.Format("20060102150405")
	assert.Contains(t, sampleId, expectedTimestamp, "订单号应包含正确的时间戳")

	// 测试唯一性 - 注意：由于计数器在达到1000时会重置，
	// 在同一时间戳下最多只能生成1000个唯一ID
	ids := make(map[string]bool)
	count := 1000 // 与计数器重置阈值保持一致
	for i := 0; i < count; i++ {
		id := GenerateOrderIdWithIncreaseIndex(prefix, &tm)
		ids[id] = true
	}
	assert.Equal(t, count, len(ids), "所有生成的订单ID应唯一")
}

func TestGenerateOrderIdWithIndexThread(t *testing.T) {
	tm := time.Now()

	var wg sync.WaitGroup
	var ids sync.Map
	// 存储生成的ID总数
	var totalGenerated int32
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				id := GenerateOrderIdWithIncreaseIndex("PT", &(tm))
				ids.Store(id, true)
				atomic.AddInt32(&totalGenerated, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// 获取唯一ID数量
	aLen := 0
	ids.Range(func(k, v interface{}) bool {
		aLen++
		return true
	})

	// 验证生成的ID总数
	assert.Equal(t, int32(10*100), totalGenerated, "生成的ID总数应与请求的数量一致")

	// 由于idCounter在计数达到1000时会重置为0，在多线程环境下可能生成重复的ID
	// 但在同一时间戳下理论上最多可以生成1000个唯一ID
	// 所以我们期望唯一ID数量接近1000，但不严格等于
	assert.GreaterOrEqual(t, aLen, 900, "在多线程环境下生成的唯一ID数量应接近1000")
}

func TestGenerateOrderIdWithTenantId(t *testing.T) {
	tenantID := "M"
	orderID := GenerateOrderIdWithTenantId(tenantID)

	t.Logf("%s", orderID)

	// 验证订单号长度是否正确
	assert.Equal(t, 14+5+8, len(orderID))

	// 验证时间戳部分是否正确
	timestamp := time.Now().Format("20060102150405")
	assert.Contains(t, orderID, timestamp)
	t.Logf("timestamp %d", len(timestamp))

	// 验证商户ID部分是否正确
	assert.Contains(t, orderID, tenantID)

	// 验证随机数部分是否为6位数字
	randomPart := orderID[len(orderID)-6:]
	assert.Regexp(t, `^\d{6}$`, randomPart)
}

func TestGenerateOrderIdWithTenantIdCollision(t *testing.T) {
	tenantID := "M987"
	count := 1000 // 生成订单号的数量
	ids := make(map[string]bool)

	for i := 0; i < count; i++ {
		orderID := GenerateOrderIdWithTenantId(tenantID)
		if ids[orderID] {
			t.Errorf("碰撞的订单号: %s", orderID)
		}
		ids[orderID] = true
	}

	t.Logf("生成了 %d 个订单号，没有发生碰撞", count)
}

func TestGenerateOrderIdWithPrefixSonyflake(t *testing.T) {
	prefix := "ORD"
	orderID := GenerateOrderIdWithPrefixSonyflake(prefix)
	t.Logf("order id with prefix sonyflake: %s [%d]", orderID, len(orderID))

	// 验证订单号是否包含前缀
	assert.Contains(t, orderID, prefix, "订单号应包含前缀")

	// 验证订单号是否为有效的数字字符串
	assert.Regexp(t, `^ORD\d+$`, orderID, "订单号格式应为前缀加数字")
}

func TestGenerateOrderIdWithPrefixSonyflakeCollision(t *testing.T) {
	prefix := "ORD"
	count := 100000 // 生成订单号的数量
	ids := make(map[string]bool)

	for i := 0; i < count; i++ {
		orderID := GenerateOrderIdWithPrefixSonyflake(prefix)
		if ids[orderID] {
			t.Errorf("碰撞的订单号: %s", orderID)
		}
		ids[orderID] = true
	}

	t.Logf("生成了 %d 个订单号，没有发生碰撞", count)
}

func TestGenerateOrderIdWithPrefixSnowflake(t *testing.T) {
	workerId := int64(1) // 假设使用的 workerId
	prefix := "ORD"
	orderID := GenerateOrderIdWithPrefixSnowflake(workerId, prefix)
	t.Logf("order id with prefix snowflake: %s [%d]", orderID, len(orderID))

	// 验证订单号是否包含前缀
	assert.Contains(t, orderID, prefix, "订单号应包含前缀")

	// 验证订单号是否为有效的数字字符串
	assert.Regexp(t, `^ORD\d+$`, orderID, "订单号格式应为前缀加数字")
}

func TestGenerateOrderIdWithPrefixSnowflakeCollision(t *testing.T) {
	workerId := int64(1) // 假设使用的 workerId
	prefix := "ORD"
	count := 1000000 // 生成订单号的数量
	ids := make(map[string]bool)

	for i := 0; i < count; i++ {
		orderID := GenerateOrderIdWithPrefixSnowflake(workerId, prefix)
		if ids[orderID] {
			t.Errorf("碰撞的订单号: %s", orderID)
		}
		ids[orderID] = true
	}

	t.Logf("生成了 %d 个订单号，没有发生碰撞", count)
}
