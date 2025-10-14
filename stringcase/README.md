# stringcase 包

本包提供各种编程命名法之间的转换功能，支持多种常见命名风格的相互转换，是处理字符串命名格式的实用工具。

## 功能特点

- 支持多种常见命名法之间的相互转换
- 自动识别多种分隔符（空格、下划线、连字符等）
- 智能处理缩写词（如HTTP -> Http）
- 处理数字与字母的边界
- 移除特殊字符
- 处理多余的空白字符和分隔符

## 命名法说明

### LowerCamelCase（小驼峰命名法）

小驼峰命名法，首字母小写，后续单词首字母大写。例如：`myVariableName`

### UpperCamelCase/PascalCase（大驼峰/帕斯卡命名法）

大驼峰命名法，每个单词首字母都大写。例如：`MyVariableName`

### snake_case（蛇形命名法）

使用下划线分隔单词的命名方式。例如：`my_variable_name`

### kebab-case（烤肉串命名法）

使用连字符分隔单词的命名方式。例如：`my-variable-name`

## 安装

使用Go模块安装：

```bash
go get github.com/moweilong/mo/stringcase
```

## 使用示例

```go
package main

import (
    "fmt"
    "github.com/moweilong/mo/stringcase"
)

func main() {
    // 驼峰命名法转换
    fmt.Println(stringcase.UpperCamelCase("hello world")) // 输出: HelloWorld
    fmt.Println(stringcase.LowerCamelCase("hello_world")) // 输出: helloWorld
    fmt.Println(stringcase.ToPascalCase("hello-world"))   // 输出: HelloWorld
    
    // 蛇形命名法转换
    fmt.Println(stringcase.ToSnakeCase("HelloWorld"))     // 输出: hello_world
    fmt.Println(stringcase.UpperSnakeCase("HelloWorld"))  // 输出: HELLO_WORLD
    
    // 烤肉串命名法转换
    fmt.Println(stringcase.KebabCase("HelloWorld"))       // 输出: hello-world
    fmt.Println(stringcase.UpperKebabCase("HelloWorld"))  // 输出: HELLO-WORLD
    
    // 处理包含缩写和数字的复杂字符串
    fmt.Println(stringcase.ToSnakeCase("HTTPStatusCode200")) // 输出: http_status_code_200
    fmt.Println(stringcase.KebabCase("UserID123"))           // 输出: user-id-123
}
```

## API参考

### 驼峰命名法相关函数

#### UpperCamelCase(input string) string

将输入字符串转换为大驼峰命名法（首字母大写）。

**参数**: 
- input: 输入字符串

**返回值**: 
- 转换后的大驼峰命名字符串

**示例**: 
- `UpperCamelCase("hello world")` -> `"HelloWorld"`
- `UpperCamelCase("hello_world")` -> `"HelloWorld"`

#### LowerCamelCase(input string) string

将输入字符串转换为小驼峰命名法（首字母小写）。

**参数**: 
- input: 输入字符串

**返回值**: 
- 转换后的小驼峰命名字符串

**示例**: 
- `LowerCamelCase("hello world")` -> `"helloWorld"`
- `LowerCamelCase("HELLO_WORLD")` -> `"helloWorld"`

#### ToPascalCase(input string) string

将输入字符串转换为帕斯卡命名法（等同于大驼峰）。

**参数**: 
- input: 输入字符串

**返回值**: 
- 转换后的帕斯卡命名字符串

**示例**: 
- `ToPascalCase("hello world")` -> `"HelloWorld"`
- `ToPascalCase("hello-world")` -> `"HelloWorld"`

### 蛇形命名法相关函数

#### ToSnakeCase(input string) string

将输入字符串转换为小写蛇形命名法（使用下划线分隔单词）。

**参数**: 
- input: 输入字符串

**返回值**: 
- 转换后的小写蛇形命名字符串

**示例**: 
- `ToSnakeCase("HelloWorld")` -> `"hello_world"`
- `ToSnakeCase("HTTPStatusCode")` -> `"http_status_code"`

#### SnakeCase(s string) string

ToSnakeCase的别名函数。

**参数**: 
- s: 输入字符串

**返回值**: 
- 转换后的小写蛇形命名字符串

#### UpperSnakeCase(s string) string

将输入字符串转换为大写蛇形命名法。

**参数**: 
- s: 输入字符串

**返回值**: 
- 转换后的大写蛇形命名字符串

**示例**: 
- `UpperSnakeCase("HelloWorld")` -> `"HELLO_WORLD"`
- `UpperSnakeCase("hello-world")` -> `"HELLO_WORLD"`

### 烤肉串命名法相关函数

#### KebabCase(s string) string

将输入字符串转换为小写烤肉串命名法（使用连字符分隔单词）。

**参数**: 
- s: 输入字符串

**返回值**: 
- 转换后的小写烤肉串命名字符串

**示例**: 
- `KebabCase("HelloWorld")` -> `"hello-world"`
- `KebabCase("user ID")` -> `"user-id"`

#### UpperKebabCase(s string) string

将输入字符串转换为大写烤肉串命名法。

**参数**: 
- s: 输入字符串

**返回值**: 
- 转换后的大写烤肉串命名字符串

**示例**: 
- `UpperKebabCase("HelloWorld")` -> `"HELLO-WORLD"`
- `UpperKebabCase("hello world!")` -> `"HELLO-WORLD"`

## 智能处理特性

本包具有以下智能处理特性：

1. **自动识别分隔符**：能自动识别空格、下划线、连字符等常见分隔符
2. **缩写词处理**：对常见缩写（如HTTP、URL、ID等）进行智能处理
3. **边界识别**：能够识别数字与字母之间的边界
4. **特殊字符处理**：自动过滤掉特殊字符
5. **多余分隔符处理**：自动合并多余的分隔符

## 测试用例

本包中的测试用例确保了函数的正确性和稳定性。测试用例覆盖了各种常见场景和边界情况，包括但不限于：

- 基本字符串转换
- 包含各种分隔符的字符串
- 包含缩写词的字符串
- 包含数字的字符串
- 包含特殊字符的字符串
- 空字符串和单字符处理
- 多余分隔符和空格处理

## 许可证

本包采用 MIT 许可证。
