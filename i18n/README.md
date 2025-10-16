# I18n - 国际化支持工具包

## 功能概述

i18n 是一个提供国际化（Internationalization）支持的Go工具包，基于 [go-i18n](https://github.com/nicksnyder/go-i18n) 库实现。该包提供了简单易用的API，支持多语言切换、文本翻译和错误信息本地化，适用于构建国际化的应用程序。

## 目录结构

```
i18n/
├── README.md       # 文档
├── i18n.go         # 核心实现
├── context.go      # 上下文集成
├── options.go      # 配置选项
└── i18n_test.go    # 测试文件
```

## 核心类型

### I18n 结构体

```go
// I18n is used to store the options and configurations for internationalization.
type I18n struct {
    ops       Options
    bundle    *i18n.Bundle
    localizer *i18n.Localizer
    lang      language.Tag
}
```

`I18n` 是包的核心结构体，存储国际化相关的选项和配置，包括语言包、本地化器和当前语言。

### Options 结构体

```go
type Options struct {
    format   string         // 语言文件格式（yaml、json、toml）
    language language.Tag   // 默认语言
    files    []string       // 语言文件路径列表
    fs       embed.FS       // 嵌入的文件系统
}
```

`Options` 结构体定义了国际化配置选项，用于自定义i18n实例的行为。

## 主要函数和方法

### 1. 创建和配置

#### New

```go
// New creates a new instance of the I18n struct with the given options.
func New(options ...func(*Options)) *I18n
```

创建一个新的I18n实例，支持通过函数选项模式进行配置。

**参数**：
- `options` - 可选的配置函数列表

**返回值**：
- 初始化后的 `I18n` 实例

#### 配置选项函数

```go
// WithFormat 设置语言文件格式（yaml、json、toml）
func WithFormat(format string) func(*Options)

// WithLanguage 设置默认语言
func WithLanguage(lang language.Tag) func(*Options)

// WithFile 添加语言文件
func WithFile(f string) func(*Options)

// WithFS 添加嵌入的文件系统
func WithFS(fs embed.FS) func(*Options)
```

这些函数用于配置I18n实例的各种选项。

### 2. 语言切换

#### Select

```go
// Select can change language.
func (i I18n) Select(lang language.Tag) *I18n
```

切换当前使用的语言。

**参数**：
- `lang` - 目标语言标签

**返回值**：
- 切换语言后的新 `I18n` 实例

#### Language

```go
// Language get current language.
func (i I18n) Language() language.Tag
```

获取当前使用的语言。

**返回值**：
- 当前语言标签

### 3. 翻译功能

#### T

```go
// T localizes the message with the given ID and returns the localized string.
func (i I18n) T(id string) string
```

根据消息ID翻译文本。

**参数**：
- `id` - 消息ID

**返回值**：
- 翻译后的文本，如果无法翻译则返回消息ID

#### E

```go
// E is a wrapper for T that converts the localized string to an error type and returns it.
func (i I18n) E(id string) error
```

根据消息ID翻译错误信息。

**参数**：
- `id` - 消息ID

**返回值**：
- 包含翻译后文本的错误对象

#### LocalizeT

```go
// LocalizeT localizes the given message and returns the localized string.
func (i I18n) LocalizeT(message *i18n.Message) string
```

本地化给定的消息对象。

**参数**：
- `message` - 消息对象

**返回值**：
- 翻译后的文本，如果无法翻译则返回消息ID

#### LocalizeE

```go
// LocalizeE is a wrapper for LocalizeT method that converts the localized string to an error type and returns it.
func (i I18n) LocalizeE(message *i18n.Message) error
```

本地化给定的消息对象并返回错误类型。

**参数**：
- `message` - 消息对象

**返回值**：
- 包含翻译后文本的错误对象

### 4. 语言文件管理

#### Add

```go
// Add is add language file or dir(auto get language by filename).
func (i *I18n) Add(f string)
```

添加语言文件或目录。如果是目录，会自动遍历并加载所有文件。

**参数**：
- `f` - 文件或目录路径

#### AddFS

```go
// AddFS is add language embed files.
func (i *I18n) AddFS(fs embed.FS)
```

添加嵌入的语言文件系统。

**参数**：
- `fs` - 嵌入的文件系统

### 5. 上下文集成

#### WithContext

```go
func WithContext(ctx context.Context, i *I18n) context.Context
```

将I18n实例存储到上下文中。

**参数**：
- `ctx` - 上下文对象
- `i` - I18n实例

**返回值**：
- 包含I18n实例的新上下文

#### FromContext

```go
func FromContext(ctx context.Context) *I18n
```

从上下文中获取I18n实例。如果上下文中没有，则返回默认实例。

**参数**：
- `ctx` - 上下文对象

**返回值**：
- I18n实例

## 使用示例

### 1. 基本使用

```go
import (
    "fmt"
    "github.com/moweilong/mo/i18n"
    "golang.org/x/text/language"
)

func main() {
    // 创建默认I18n实例（英文）
    t := i18n.New()
    
    // 添加语言文件目录
    t.Add("./locales")
    
    // 翻译文本
    hello := t.T("common.hello")
    fmt.Println(hello) // 输出英文翻译
    
    // 切换到中文
    tChinese := t.Select(language.Chinese)
    helloZh := tChinese.T("common.hello")
    fmt.Println(helloZh) // 输出中文翻译
    
    // 创建翻译错误
    err := t.E("error.not_found")
    fmt.Println(err.Error()) // 输出翻译后的错误信息
}
```

### 2. 自定义配置

```go
import (
    "embed"
    "github.com/moweilong/mo/i18n"
    "golang.org/x/text/language"
)

//go:embed locales
var localeFS embed.FS

func main() {
    // 创建自定义配置的I18n实例
    t := i18n.New(
        i18n.WithLanguage(language.Chinese),  // 设置默认语言为中文
        i18n.WithFormat("json"),              // 设置语言文件格式为JSON
        i18n.WithFile("./locales/zh.json"),   // 添加单个语言文件
        i18n.WithFS(localeFS),                // 添加嵌入的文件系统
    )
    
    // 使用翻译器
    fmt.Println(t.T("common.welcome"))
}
```

### 3. 上下文集成

```go
import (
    "context"
    "fmt"
    "github.com/moweilong/mo/i18n"
    "golang.org/x/text/language"
)

func main() {
    // 创建I18n实例
    t := i18n.New(
        i18n.WithLanguage(language.Chinese),
        i18n.WithFile("./locales/zh.yml"),
    )
    
    // 将I18n实例存储到上下文中
    ctx := i18n.WithContext(context.Background(), t)
    
    // 在其他函数中从上下文中获取I18n实例
    useTranslator(ctx)
}

func useTranslator(ctx context.Context) {
    // 从上下文中获取I18n实例
    t := i18n.FromContext(ctx)
    
    // 使用翻译器
    fmt.Println(t.T("common.goodbye"))
}
```

### 4. 使用Message对象

```go
import (
    "fmt"
    "github.com/moweilong/mo/i18n"
    "github.com/nicksnyder/go-i18n/v2/i18n"
)

func main() {
    t := i18n.New()
    t.Add("./locales")
    
    // 创建Message对象
    message := &i18n.Message{
        ID:    "greeting",
        Other: "Hello, {{.Name}}!",
    }
    
    // 翻译并格式化消息
    greeting := t.LocalizeT(message)
    fmt.Println(greeting) // 输出翻译后的问候语
    
    // 翻译错误消息
    errMessage := &i18n.Message{
        ID:    "error.invalid_input",
        Other: "Invalid input provided",
    }
    err := t.LocalizeE(errMessage)
    fmt.Println(err.Error()) // 输出翻译后的错误消息
}
```

## 语言文件格式

支持三种格式的语言文件：YAML（默认）、JSON和TOML。以下是一个YAML格式的示例：

```yaml
# en.yml
i18n:
  common:
    hello: "Hello"
    welcome: "Welcome to our application"
  error:
    not_found: "Resource not found"
    invalid_input: "Invalid input provided"
```

```yaml
# zh.yml
i18n:
  common:
    hello: "你好"
    welcome: "欢迎使用我们的应用"
  error:
    not_found: "未找到资源"
    invalid_input: "提供的输入无效"
```

## 特性和优势

1. **简单易用**：提供简洁的API，易于集成到各种项目中
2. **多格式支持**：支持YAML、JSON和TOML三种语言文件格式
3. **灵活配置**：通过函数选项模式支持灵活的配置
4. **上下文集成**：提供上下文集成功能，方便在请求处理中传递翻译器
5. **错误本地化**：支持错误消息的本地化
6. **嵌入文件系统支持**：支持从嵌入的文件系统加载语言文件

## 注意事项

1. 语言文件的命名应遵循go-i18n的约定，包含语言标签（如en、zh、zh-CN等）
2. 如果无法找到对应的翻译，会返回消息ID作为默认值
3. 使用嵌入文件系统时，需要确保文件路径正确
4. 切换语言时，Select方法会返回一个新的I18n实例，而不是修改原实例


i18n of different languages based [go-i18n](https://github.com/nicksnyder/go-i18n).


## Usage


```bash
go get -u github.com/onexstack/onexstack/pkg/i18n
```

add language files

```bash
mkdir locales

cat <<EOF > locales/en.yml
hello.world: Hello world!
EOF

cat <<EOF > locales/zh.yml
hello.world: 你好, 世界!
EOF
```

```go
import (
	"embed"
	"fmt"
	"golang.org/x/text/language"
  "github.com/onexstack/onexstack/pkg/i18n"
)

//go:embed locales
var locales embed.FS

func main() {
	i := i18n.New(
		i18n.WithFormat("yml"),
		// with absolute files
		i18n.WithFile("locales/en.yml"),
		i18n.WithFile("locales/zh.yml"),
		// with go embed files
		// i18n.WithFs(locales),
		i18n.WithLanguage(language.Chinese),
	)

	// print string
	fmt.Println(i.T("hello.world"))
	// 你好, 世界!

	// print error
	fmt.Println(i.E("hello.world").Error() == "你好, 世界!")
	// true

	// override default language
	fmt.Println(i.Select(language.English).T("hello.world"))
	// Hello world!
}
```


## Options


- `WithFormat` - language file format, default yml
- `WithLanguage` - set default language file format, default en
- `WithFile` - set language files by file system
- `WithFs` - set language files by go embed file
