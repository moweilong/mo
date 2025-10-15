# 分页工具包

## 功能概述

pagination 是一个提供分页相关功能的工具包，主要用于计算分页偏移量和提供默认的分页参数值。

## 常量定义

```go
const (
	DefaultPage     = 1  // 默认页数
	DefaultPageSize = 10 // 默认每页行数
)
```

## 主要函数

### GetPageOffset
```go
func GetPageOffset(pageNum, pageSize int32) int
```

根据页码和每页行数计算查询偏移量。在数据库查询中，常用于SQL语句的OFFSET子句。

参数：
- `pageNum`：当前页码，从1开始计数
- `pageSize`：每页显示的记录数量

返回值：
- `int`：计算得到的偏移量，表示从第几条记录开始查询

## 安装

```bash
go get github.com/moweilong/mo/pagination
```

## 使用示例

### 基本使用

```go
package main

import (
	"fmt"
	"github.com/moweilong/mo/pagination"
)

func main() {
	// 使用默认页码和页大小计算偏移量
	pageNum := int32(pagination.DefaultPage)
	pageSize := int32(pagination.DefaultPageSize)
	offset := pagination.GetPageOffset(pageNum, pageSize)
	fmt.Printf("第 %d 页，每页 %d 条记录，偏移量为: %d\n", pageNum, pageSize, offset)
	// 输出: 第 1 页，每页 10 条记录，偏移量为: 0

	// 使用自定义页码和页大小计算偏移量
	pageNum = 3
	pageSize = 20
	offset = pagination.GetPageOffset(pageNum, pageSize)
	fmt.Printf("第 %d 页，每页 %d 条记录，偏移量为: %d\n", pageNum, pageSize, offset)
	// 输出: 第 3 页，每页 20 条记录，偏移量为: 40
}
```

### 结合数据库查询

```go
package main

import (
	"database/sql"
	"fmt"
	"github.com/moweilong/mo/pagination"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 假设这些是从请求参数中获取的
	pageNum := int32(2)
	pageSize := int32(15)

	// 确保页码和页大小是有效的
	if pageNum <= 0 {
		pageNum = int32(pagination.DefaultPage)
	}
	if pageSize <= 0 {
		pageSize = int32(pagination.DefaultPageSize)
	}

	// 计算偏移量
	offset := pagination.GetPageOffset(pageNum, pageSize)

	// 构建SQL查询语句
	sqlQuery := fmt.Sprintf("SELECT * FROM users LIMIT %d OFFSET %d", pageSize, offset)
	fmt.Printf("生成的SQL查询: %s\n", sqlQuery)
	// 输出: 生成的SQL查询: SELECT * FROM users LIMIT 15 OFFSET 15

	// 实际执行查询（这里省略具体的数据库连接和查询代码）
	// db, _ := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	// rows, _ := db.Query(sqlQuery)
	// ...
}
```

## 注意事项

- 页码从1开始计数，而不是从0开始
- 当传入的页码小于等于0时，建议使用DefaultPage作为默认值
- 当传入的页大小小于等于0时，建议使用DefaultPageSize作为默认值
- GetPageOffset函数返回的是int类型，在与不同数据库交互时，可能需要根据具体驱动进行类型转换
- 该包设计简洁，仅提供分页计算的基础功能，不包含复杂的分页逻辑或与特定ORM框架的集成