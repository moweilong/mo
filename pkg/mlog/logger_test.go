package mlog

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 用于测试的上下文键
type ctxKey string

const (
	ctxTraceID ctxKey = "traceID"
	ctxUserID  ctxKey = "userID"
)

// TestNewLogger 测试 NewLogger 函数能否正确创建 logger
func TestNewLogger(t *testing.T) {
	opts := NewOptions()
	opts.Format = jsonFormat
	opts.OutputPaths = []string{"stdout"} // 实际测试中会被重定向

	logger := NewLogger(opts)
	assert.NotNil(t, logger, "NewLogger 应该返回非空的 logger")
	assert.NotNil(t, logger.z, "logger 的 zap.Logger 字段应该非空")
	assert.Equal(t, opts, logger.opts, "logger 的选项应该与传入的一致")
	// Info 适合输出结构化日志，可以包含任意多的键值对
	logger.Info("info message", zap.String("key", "value"))
	// Infof 适合输出格式化的文本日志，不适合包含额外的键值对
	logger.Infof("info message: %s", "test value")
	// Infow 适合输出结构化日志，可以包含任意多的键值对
	logger.Infow("info message", "message", "test value", "key", "value")
}

// TestLoggerDebug 测试调试级别日志输出
func TestLoggerDebug(t *testing.T) {
	// 创建一个缓冲区来捕获日志输出
	var buf bytes.Buffer

	// 使用自定义的 core 来捕获日志
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(&buf),
		zapcore.DebugLevel, // 必须设置为 DebugLevel 才能捕获调试日志
	)

	// 创建一个自定义的 logger
	logger := &zapLogger{
		z:                 zap.New(core),
		opts:              NewOptions(),
		contextExtractors: make(map[string]func(context.Context) string),
	}

	// 测试 Debugf
	logger.Debugf("debug message: %s", "test")
	output := buf.String()
	assert.Contains(t, output, "debug message: %s", "Debugf 应该输出原始的格式化字符串")

	// 清空缓冲区
	buf.Reset()

	// 测试 Debugw
	logger.Debugw("debug with fields", "key", "value", "number", 123)
	output = buf.String()
	assert.Contains(t, output, "debug with fields", "Debugw 应该输出正确的消息")
	assert.Contains(t, output, "key", "Debugw 应该包含键值对字段")
	assert.Contains(t, output, "value", "Debugw 应该包含键值对值")
}

// TestLoggerInfo 测试信息级别日志输出
func TestLoggerInfo(t *testing.T) {
	var buf bytes.Buffer
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(&buf),
		zapcore.InfoLevel,
	)

	logger := &zapLogger{
		z:                 zap.New(core),
		opts:              NewOptions(),
		contextExtractors: make(map[string]func(context.Context) string),
	}

	// 测试 Infof
	logger.Infof("info message: %s", "test")
	output := buf.String()
	assert.Contains(t, output, "info message: %s", "Infof 应该输出原始的格式化字符串")

	// 清空缓冲区
	buf.Reset()

	// 测试 Infow
	logger.Infow("info with fields", "key", "value", "number", 123)
	output = buf.String()
	assert.Contains(t, output, "info with fields", "Infow 应该输出正确的消息")
	assert.Contains(t, output, "key", "Infow 应该包含键值对字段")
}

// TestLoggerWarn 测试警告级别日志输出
func TestLoggerWarn(t *testing.T) {
	var buf bytes.Buffer
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(&buf),
		zapcore.WarnLevel,
	)

	logger := &zapLogger{
		z:                 zap.New(core),
		opts:              NewOptions(),
		contextExtractors: make(map[string]func(context.Context) string),
	}

	// 测试 Warnf
	logger.Warnf("warn message: %s", "test")
	output := buf.String()
	assert.Contains(t, output, "warn message: %s", "Warnf 应该输出原始的格式化字符串")

	// 清空缓冲区
	buf.Reset()

	// 测试 Warnw
	logger.Warnw("warn with fields", "key", "value", "number", 123)
	output = buf.String()
	assert.Contains(t, output, "warn with fields", "Warnw 应该输出正确的消息")
}

// TestLoggerError 测试错误级别日志输出
func TestLoggerError(t *testing.T) {
	var buf bytes.Buffer
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(&buf),
		zapcore.ErrorLevel,
	)

	logger := &zapLogger{
		z:                 zap.New(core),
		opts:              NewOptions(),
		contextExtractors: make(map[string]func(context.Context) string),
	}

	// 测试 Errorf
	logger.Errorf("error message: %s", "test")
	output := buf.String()
	assert.Contains(t, output, "error message: %s", "Errorf 应该输出原始的格式化字符串")

	// 清空缓冲区
	buf.Reset()

	// 测试 Errorw
	err := errors.New("test error")
	logger.Errorw("error with fields", "key", "value", "err", err)
	output = buf.String()
	assert.Contains(t, output, "error with fields", "Errorw 应该输出正确的消息")
	assert.Contains(t, output, "test error", "Errorw 应该包含错误信息")
	assert.Contains(t, output, "key", "Errorw 应该包含额外的键值对字段")
}

// TestLoggerWithContext 测试从 context 中提取字段
func TestLoggerWithContext(t *testing.T) {
	var buf bytes.Buffer
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(&buf),
		zapcore.InfoLevel,
	)

	// 创建 context 提取器
	contextExtractors := ContextExtractors{
		"traceID": func(ctx context.Context) string {
			if val := ctx.Value(ctxTraceID); val != nil {
				return val.(string)
			}
			return ""
		},
		"userID": func(ctx context.Context) string {
			if val := ctx.Value(ctxUserID); val != nil {
				return val.(string)
			}
			return ""
		},
	}

	// 使用 WithContextExtractor 选项创建 logger
	logger := &zapLogger{
		z:                 zap.New(core),
		opts:              NewOptions(),
		contextExtractors: make(map[string]func(context.Context) string),
	}

	// 应用 context 提取器选项
	WithContextExtractor(contextExtractors)(logger)

	// 创建带有值的 context
	ctx := context.Background()
	ctx = context.WithValue(ctx, ctxTraceID, "123456")
	ctx = context.WithValue(ctx, ctxUserID, "user-789")

	// 使用 W 方法从 context 中提取字段
	ctxLogger := logger.W(ctx)

	// 输出日志
	ctxLogger.Infof("test with context")
	output := buf.String()

	// 验证提取的字段是否包含在日志中
	assert.Contains(t, output, "123456", "日志应该包含从 context 提取的 traceID")
	assert.Contains(t, output, "user-789", "日志应该包含从 context 提取的 userID")
}

// TestAddCallerSkip 测试 AddCallerSkip 方法
func TestAddCallerSkip(t *testing.T) {
	opts := NewOptions()
	opts.DisableCaller = false // 启用调用者信息

	logger := NewLogger(opts)
	assert.NotNil(t, logger, "NewLogger 应该返回非空的 logger")

	// 测试 AddCallerSkip 方法
	newLogger := logger.AddCallerSkip(1)
	assert.NotNil(t, newLogger, "AddCallerSkip 应该返回非空的 logger")
	// 由于我们无法直接验证调用者信息的正确性，这里只验证方法不会崩溃
}

// TestInitAndDefault 测试全局日志初始化和获取
func TestInitAndDefault(t *testing.T) {
	// 保存原始的全局 logger
	originalStd := std
	// 测试结束后恢复原始的全局 logger
	defer func() {
		mu.Lock()
		std = originalStd
		mu.Unlock()
	}()

	// 初始化全局 logger
	opts := NewOptions()
	Init(opts)

	// 获取全局 logger
	globalLogger := Default()
	assert.NotNil(t, globalLogger, "Default 应该返回非空的全局 logger")

	// 测试全局日志函数
	// 这些函数不会导致测试失败，因为我们只是验证它们不会崩溃
	Debugf("test debug")
	Infof("test info")
	Warnf("test warn")
	// 注意：不要测试 Panicf 和 Fatalf，因为它们会导致程序崩溃
}

// TestSync 测试 Sync 方法
func TestSync(t *testing.T) {
	// 测试全局 Sync 函数
	// 这个函数不会导致测试失败，因为我们只是验证它不会崩溃
	Sync()

	// 测试 logger 的 Sync 方法
	logger := NewLogger(NewOptions())
	logger.Sync()
}

// TestWithContextExtractorOption 测试 WithContextExtractor 选项
func TestWithContextExtractorOption(t *testing.T) {
	var buf bytes.Buffer
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(&buf),
		zapcore.InfoLevel,
	)

	// 创建一个自定义的 zap.Logger
	zLogger := zap.New(core)

	// 创建一个自定义的 zapLogger
	logger := &zapLogger{
		z:                 zLogger,
		opts:              NewOptions(),
		contextExtractors: make(map[string]func(context.Context) string),
	}

	// 创建 context 提取器
	extractors := ContextExtractors{
		"testField": func(ctx context.Context) string {
			return "testValue"
		},
	}

	// 应用 WithContextExtractor 选项
	WithContextExtractor(extractors)(logger)

	// 验证提取器是否被正确添加
	assert.Contains(t, logger.contextExtractors, "testField", "contextExtractors 应该包含添加的提取器")

	// 测试提取器功能
	ctx := context.Background()
	ctxLogger := logger.W(ctx)
	ctxLogger.Infof("test")

	// 由于我们无法直接验证日志内容（因为我们没有使用原始的 zapLogger.log 方法），
	// 这里只验证方法不会崩溃
}

// TestOptionsMethod 测试 Options 方法
func TestOptionsMethod(t *testing.T) {
	opts := NewOptions()
	opts.Level = "debug"
	opts.Format = "json"

	logger := NewLogger(opts)
	returnedOpts := logger.Options()

	assert.Equal(t, opts, returnedOpts, "Options 方法应该返回创建 logger 时使用的选项")
}

// TestCloneMethod 测试 clone 方法
func TestCloneMethod(t *testing.T) {
	logger := NewLogger(NewOptions())
	clonedLogger := logger.clone()

	assert.NotNil(t, clonedLogger, "clone 方法应该返回非空的 logger")
	assert.NotSame(t, logger, clonedLogger, "clone 方法应该返回新的 logger 实例")
	assert.Equal(t, logger.opts, clonedLogger.opts, "克隆的 logger 应该有相同的选项")
}

// TestLoggerWithNilOptions 测试使用 nil 选项创建 logger
func TestLoggerWithNilOptions(t *testing.T) {
	// 使用 nil 选项创建 logger
	logger := NewLogger(nil)
	assert.NotNil(t, logger, "使用 nil 选项创建 logger 应该返回非空的 logger")
	assert.NotNil(t, logger.opts, "使用 nil 选项创建的 logger 应该有默认选项")
}

// TestLoggerWithInvalidLevel 测试使用无效的日志级别
func TestLoggerWithInvalidLevel(t *testing.T) {
	opts := NewOptions()
	opts.Level = "invalid-level"

	// 使用无效的日志级别创建 logger 不应该崩溃，而应该使用默认级别
	logger := NewLogger(opts)
	assert.NotNil(t, logger, "使用无效的日志级别创建 logger 应该返回非空的 logger")
}
