# Id
# idx - ID 编码与生成库

idx 是一个提供 ID 编码和分布式 ID 生成功能的库，专注于将数字 ID 转换为更友好、更安全的字符串形式，以及基于 Sonyflake 算法生成分布式唯一 ID。

## 功能概述

- **ID 编码**：将 uint64 类型的数字 ID 转换为可读性更好、更安全的字符串形式
- **Sonyflake 分布式 ID**：基于 Sonyflake 算法实现的高并发、分布式环境下的唯一 ID 生成器

## 目录结构

```
idx/
├── code.go       // ID 编码功能实现
├── code_test.go  // 编码功能测试
├── doc.go        // 包文档
├── options.go    // 配置选项定义
└── sonyflake.go  // Sonyflake ID 生成器实现
```

## ID 编码功能

### 实现原理

ID 编码功能通过 `NewCode` 函数实现，主要采用以下步骤将数字 ID 转换为字符串：

1. **数值放大与加盐**：将输入的 ID 乘以一个系数并加上盐值，增强安全性
2. **扩散算法**：将数字分解并通过扩散算法使每一位数字相互影响
3. **混淆算法**：通过置换盒（Permutation Box）对字符位置进行混淆，进一步增强安全性
4. **字符映射**：将处理后的数字映射到预定义的字符集上

### 使用示例

```go
import (
	"fmt"
	"github.com/moweilong/mo/idx"
)

func main() {
	// 使用默认配置生成编码
	code1 := idx.NewCode(1)  // 输出: "VHB4JX86"
	fmt.Println(code1)

	// 使用自定义配置生成编码
	code2 := idx.NewCode(
		1,
		idx.WithCodeChars([]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}),
		idx.WithCodeN1(9),
		idx.WithCodeN2(3),
		idx.WithCodeL(5),
		idx.WithCodeSalt(56789),
	)  // 输出: "80773"
	fmt.Println(code2)

	// 确保不同ID生成不同编码
	fmt.Println(idx.NewCode(2))  // 输出: "98CST47F"
}
```

### 配置选项

| 选项 | 说明 | 默认值 |
|------|------|--------|
| `WithCodeChars` | 编码字符集 | 移除了易混淆字符的30个字符：'2','3','4','5','6','7','8','9','A','B','C','D','E','F','G','H','J','K','L','M','N','P','Q','R','S','T','V','W','X','Y' |
| `WithCodeL` | 编码长度 | 8 |
| `WithCodeN1` | 与字符集大小互质的系数 | 17 |
| `WithCodeN2` | 与编码长度互质的系数 | 5 |
| `WithCodeSalt` | 编码盐值 | 123567369 |

## Sonyflake 分布式 ID 生成器

### 实现原理

Sonyflake 是一种分布式 ID 生成算法，生成的 ID 结构为：
- 39 位时间戳（自自定义起始时间起的毫秒数）
- 8 位机器 ID
- 16 位序列号

该算法比传统的 Snowflake 算法更适合于分布式系统，尤其是在容器化环境中。

### 使用示例

```go
import (
	"context"
	"fmt"
	"time"
	
	"github.com/moweilong/mo/idx"
)

func main() {
	// 创建 Sonyflake 实例，使用默认配置
	sf1 := idx.NewSonyflake()
	if sf1.Error != nil {
		fmt.Printf("创建 Sonyflake 失败: %v\n", sf1.Error)
		return
	}
	// 生成 ID
	id1 := sf1.Id(context.Background())
	fmt.Printf("生成的 ID: %d\n", id1)

	// 使用自定义配置创建 Sonyflake 实例
	sf2 := idx.NewSonyflake(
		idx.WithSonyflakeMachineId(42),
		idx.WithSonyflakeStartTime(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
	)
	if sf2.Error != nil {
		fmt.Printf("创建 Sonyflake 失败: %v\n", sf2.Error)
		return
	}
	id2 := sf2.Id(context.Background())
	fmt.Printf("生成的 ID: %d\n", id2)
}
```

### 配置选项

| 选项 | 说明 | 默认值 |
|------|------|--------|
| `WithSonyflakeMachineId` | 机器 ID | 1 |
| `WithSonyflakeStartTime` | ID 生成的起始时间 | 2022-10-10 00:00:00 UTC |

## 特性说明

1. **安全性**：通过扩散和混淆算法增强 ID 编码的安全性
2. **可读性**：默认字符集移除了易混淆字符（0,1,I,O,U,Z），提高了编码的可读性
3. **唯一性**：确保不同的原始 ID 生成不同的字符串编码
4. **可配置性**：提供丰富的配置选项，可以根据需要自定义编码规则
5. **高可用性**：Sonyflake 生成器在 ID 生成失败时会自动重试

## 注意事项

1. 使用 Sonyflake 时，建议为不同的服务实例分配不同的机器 ID，以避免 ID 冲突
2. Sonyflake 的起始时间设置后不应随意更改，否则可能导致 ID 重复
3. 编码算法的参数（如 N1、N2、字符集大小）需要保持互质关系，以确保编码的唯一性
4. 对于需要长期稳定运行的系统，建议根据实际需求调整 Sonyflake 的配置参数

## 测试

```bash
# 运行测试
go test -v ./idx

# 运行性能测试
go test -bench=. ./idx
```
