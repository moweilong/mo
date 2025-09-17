package mlog

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"go.uber.org/zap/zapcore"
)

const (
	flagLevel             = "log.level"
	flagDisableCaller     = "log.disable-caller"
	flagDisableStacktrace = "log.disable-stacktrace"
	flagFormat            = "log.format"
	flagEnableColor       = "log.enable-color"
	flagOutputPaths       = "log.output-paths"
	flagErrorOutputPaths  = "log.error-output-paths"
	flagDevelopment       = "log.development"
	flagName              = "log.name"

	consoleFormat = "console"
	jsonFormat    = "json"
)

// Options 包含与日志相关的配置项。
type Options struct {
	OutputPaths       []string `json:"output-paths"       mapstructure:"output-paths"`
	Level             string   `json:"level"              mapstructure:"level"`
	Format            string   `json:"format"             mapstructure:"format"`
	DisableCaller     bool     `json:"disable-caller"     mapstructure:"disable-caller"`
	DisableStacktrace bool     `json:"disable-stacktrace" mapstructure:"disable-stacktrace"`
	EnableColor       bool     `json:"enable-color"       mapstructure:"enable-color"`
	Development       bool     `json:"development"        mapstructure:"development"`
	Name              string   `json:"name"               mapstructure:"name"`
}

// NewOptions 创建一个带有默认参数的 Options 对象。
func NewOptions() *Options {
	return &Options{
		Level:             zapcore.InfoLevel.String(),
		DisableCaller:     false,
		DisableStacktrace: false,
		Format:            jsonFormat,
		EnableColor:       false,
		Development:       false,
		OutputPaths:       []string{"stdout"},
	}
}

// Validate 验证选项字段。
func (o *Options) Validate() []error {
	var errs []error

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(o.Level)); err != nil {
		errs = append(errs, err)
	}

	format := strings.ToLower(o.Format)
	if format != consoleFormat && format != jsonFormat {
		errs = append(errs, fmt.Errorf("not a valid log format: %q", o.Format))
	}

	return errs
}

// AddFlags 向指定的 FlagSet 对象添加日志相关的标志。
func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Level, flagLevel, o.Level, "最低日志输出 `LEVEL`。")
	fs.BoolVar(&o.DisableCaller, flagDisableCaller, o.DisableCaller, "禁用日志中的调用者信息输出。")
	fs.BoolVar(&o.DisableStacktrace, flagDisableStacktrace,
		o.DisableStacktrace, "禁用日志记录所有 panic 级别或以上消息的堆栈跟踪。")
	fs.StringVar(&o.Format, flagFormat, o.Format, "日志输出 `FORMAT`，支持文本或 json 格式。")
	fs.BoolVar(&o.EnableColor, flagEnableColor, o.EnableColor, "在文本格式日志中启用 ANSI 颜色输出。")
	fs.StringSliceVar(&o.OutputPaths, flagOutputPaths, o.OutputPaths, "日志输出路径。")
	fs.BoolVar(
		&o.Development,
		flagDevelopment,
		o.Development,
		"Development 将日志器置于开发模式，这会改变 DPanicLevel 的行为并更自由地记录堆栈跟踪。",
	)
	fs.StringVar(&o.Name, flagName, o.Name, "日志器的名称。")
}
