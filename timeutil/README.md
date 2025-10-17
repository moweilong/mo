# timeutil - 时间工具包

## 功能概述

`timeutil` 是一个提供丰富时间处理功能的Go语言工具包，包含时间常量定义、时间格式化与解析、时间差计算、时间范围处理以及时间类型转换等功能。该包简化了Go标准库中时间处理的复杂性，提供了更加便捷、直观的API。

## 目录结构

```
timeutil/
├── consts.go      # 时间格式常量定义
├── diff.go        # 时间差计算
├── diff_test.go   # 时间差计算测试
├── format.go      # 时间格式化与解析
├── format_test.go # 时间格式化与解析测试
├── range.go       # 时间范围处理
├── range_test.go  # 时间范围处理测试
├── trans.go       # 时间类型转换
└── trans_test.go  # 时间类型转换测试
```

## 核心功能

### 1. 时间格式常量

**consts.go** 文件中定义了多种常用的时间格式常量：

```go
// 基本时间格式
const (
    DateLayout  = "2006-01-02"       // 日期格式
    ClockLayout = "15:04:05"         // 时间格式
    TimeLayout  = DateLayout + " " + ClockLayout // 完整日期时间格式
    
    DefaultTimeLocationName = "Asia/Shanghai" // 默认时区名称
)
```

支持的其他格式包括：RFC3339、ISO8601、SQLTimestamp、DIN5008等多种标准格式。

### 2. 时间差计算

**diff.go** 文件提供了计算时间间隔的函数：

- `DayDifferenceHours(startDate, endDate string) float64` - 计算两个日期字符串之间的小时差
- `StringDifferenceDays(startDate, endDate string) int` - 计算两个日期字符串之间的天数差
- `TimeDifferenceDays(startDate, endDate time.Time) int` - 计算两个time.Time对象之间的天数差
- `SecondsDifferenceDays(startSecond, endSecond int64) int` - 计算两个秒级时间戳之间的天数差

### 3. 时间格式化与解析

**format.go** 文件提供了丰富的时间格式化和解析功能：

- `FormatTimer(d time.Duration) string` - 将时间间隔格式化为 [HH:]MM:SS 格式
- `FromTo2(fromLayout, toLayout, value string) (string, error)` - 在不同格式之间转换时间字符串
- `ParseFirst(layouts []string, value string) (time.Time, error)` - 尝试使用多个格式解析时间字符串
- `ParseOrZero(layout, value string) time.Time` - 解析时间，如果失败则返回零值时间

该文件还定义了特殊的时间类型，用于JSON序列化和反序列化：

- `RFC3339YMDTime` - 支持仅日期格式的JSON序列化和反序列化
- `ISO8601NoTzMilliTime` - 支持带毫秒但无时区的ISO8601格式JSON序列化和反序列化

### 4. 时间范围处理

**range.go** 文件提供了获取常见时间范围的功能：

- 今天/昨天时间范围：`GetTodayRangeTime()` / `GetYesterdayRangeTime()`
- 本月/上月时间范围：`GetCurrentMonthRangeTime()` / `GetLastMonthRangeTime()`
- 今年/去年时间范围：`GetCurrentYearRangeTime()` / `GetLastYearRangeTime()`

同时提供了获取日期字符串范围和时间字符串范围的函数，例如：
- `GetTodayRangeDateString()` - 获取今天的日期字符串范围
- `GetCurrentMonthRangeTimeString()` - 获取本月的时间字符串范围

### 5. 时间类型转换

**trans.go** 文件提供了不同时间表示形式之间的转换：

- 毫秒时间戳与字符串之间的转换：
  - `UnixMilliToStringPtr(milli *int64) *string`
  - `StringToUnixMilliInt64Ptr(tm *string) *int64`

- 秒时间戳与时间对象之间的转换：
  - `UnixSecondToTimePtr(second *int64) *time.Time`
  - `TimeToUnixSecondInt64Ptr(tm *time.Time) *int64`

- 时间字符串与时间对象之间的转换：
  - `StringTimeToTime(str *string) *time.Time`
  - `TimeToTimeString(tm *time.Time) *string`
  - `StringDateToTime(str *string) *time.Time`
  - `TimeToDateString(tm *time.Time) *string`

- 时区相关函数：
  - `GetDefaultTimeLocation() *time.Location`
  - `RefreshDefaultTimeLocation(name string) *time.Location`

## 使用示例

### 1. 时间格式转换

```go
import (
    "fmt"
    "github.com/moweilong/mo/timeutil"
)

func main() {
    // 将时间字符串从一种格式转换为另一种格式
    value := "2023-12-25 14:30:45"
    newFormat, err := timeutil.FromTo2(timeutil.TimeLayout, timeutil.RFC3339, value)
    if err != nil {
        fmt.Printf("转换失败: %v\n", err)
    } else {
        fmt.Printf("转换结果: %s\n", newFormat) // 输出: 2023-12-25T14:30:45Z07:00
    }
    
    // 尝试使用多个格式解析时间字符串
    formats := []string{timeutil.TimeLayout, timeutil.RFC3339, timeutil.DateLayout}
    date, err := timeutil.ParseFirst(formats, "2023-12-25")
    if err != nil {
        fmt.Printf("解析失败: %v\n", err)
    } else {
        fmt.Printf("解析结果: %v\n", date) // 输出解析后的time.Time对象
    }
}
```

### 2. 时间差计算

```go
import (
    "fmt"
    "github.com/moweilong/mo/timeutil"
    "time"
)

func main() {
    // 计算两个日期字符串之间的天数差
    startDate := "2023-12-01"
    endDate := "2023-12-25"
    days := timeutil.StringDifferenceDays(startDate, endDate)
    fmt.Printf("%s 和 %s 之间相差 %d 天\n", startDate, endDate, days)
    
    // 计算两个time.Time对象之间的天数差
    startTime := time.Date(2023, 12, 1, 0, 0, 0, 0, time.Local)
    endTime := time.Date(2023, 12, 25, 0, 0, 0, 0, time.Local)
    dayDiff := timeutil.TimeDifferenceDays(startTime, endTime)
    fmt.Printf("两个时间之间相差 %d 天\n", dayDiff)
}
```

### 3. 获取时间范围

```go
import (
    "fmt"
    "github.com/moweilong/mo/timeutil"
)

func main() {
    // 获取今天的时间范围
    todayStart, todayEnd := timeutil.GetTodayRangeTime()
    fmt.Printf("今天开始时间: %v, 结束时间: %v\n", todayStart, todayEnd)
    
    // 获取本月的日期字符串范围
    monthStart, monthEnd := timeutil.GetCurrentMonthRangeDateString()
    fmt.Printf("本月开始日期: %s, 结束日期: %s\n", monthStart, monthEnd)
    
    // 获取昨天的时间字符串范围
    yesterdayStart, yesterdayEnd := timeutil.GetYesterdayRangeTimeString()
    fmt.Printf("昨天开始时间: %s, 结束时间: %s\n", yesterdayStart, yesterdayEnd)
}
```

### 4. 时间类型转换

```go
import (
    "fmt"
    "github.com/moweilong/mo/timeutil"
    "time"
)

func main() {
    // 时间戳转换为字符串
    now := time.Now().UnixMilli()
    timeStr := timeutil.UnixMilliToStringPtr(&now)
    if timeStr != nil {
        fmt.Printf("时间戳 %d 转换为字符串: %s\n", now, *timeStr)
    }
    
    // 字符串转换为时间
    str := "2023-12-25 14:30:45"
    timeObj := timeutil.StringTimeToTime(&str)
    if timeObj != nil {
        fmt.Printf("字符串 %s 转换为时间: %v\n", str, *timeObj)
    }
    
    // 时间对象转换为日期字符串
    dateStr := timeutil.TimeToDateString(timeObj)
    if dateStr != nil {
        fmt.Printf("时间对象转换为日期字符串: %s\n", *dateStr)
    }
}
```

### 5. 时间间隔格式化

```go
import (
    "fmt"
    "github.com/moweilong/mo/timeutil"
    "time"
)

func main() {
    // 格式化时间间隔
    duration := 3*time.Hour + 15*time.Minute + 45*time.Second
    formatted := timeutil.FormatTimer(duration)
    fmt.Printf("格式化时间间隔: %s\n", formatted) // 输出: 3:15:45
    
    // 格式化更短的时间间隔
    shortDuration := 15*time.Minute + 30*time.Second
    formattedShort := timeutil.FormatTimer(shortDuration)
    fmt.Printf("格式化短时间间隔: %s\n", formattedShort) // 输出: 15:30
    
    // 使用自定义格式
    customFormat := timeutil.FormatTimerf("%d小时%d分%d秒", duration)
    fmt.Printf("自定义格式化: %s\n", customFormat) // 输出: 3小时15分45秒
}
```

## 特性和优势

1. **丰富的时间格式**：提供了多种标准和常用的时间格式常量，满足不同场景需求
2. **全面的转换功能**：支持各种时间表示形式之间的相互转换
3. **便捷的时间范围**：提供了获取常见时间范围的函数，如今天、昨天、本月等
4. **灵活的时间解析**：支持使用多个格式尝试解析时间字符串
5. **时区支持**：默认使用上海时区，可自定义默认时区
6. **零值处理**：提供了对nil指针和零值时间的安全处理

## 注意事项

1. 所有涉及时区的操作默认使用"Asia/Shanghai"时区，可通过`RefreshDefaultTimeLocation`函数修改
2. 时间差计算函数返回的值可能因计算方法的不同而略有差异
3. 在解析时间字符串时，如果没有找到匹配的格式，部分函数会返回零值时间
4. 字符串和时间对象之间的转换函数支持多种常见格式，但如果格式不匹配，可能会返回nil

## 兼容性

该包兼容Go 1.18及以上版本，依赖于Go标准库中的time包，并使用了`google.golang.org/protobuf`库处理Protocol Buffers相关的时间类型。