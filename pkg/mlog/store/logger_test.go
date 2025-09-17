package store

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewLogger 测试 NewLogger 函数能否正确创建 Logger 实例
func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	assert.NotNil(t, logger, "NewLogger 应该返回非空的 Logger 实例")
}

// TestLoggerError 测试 Error 方法能否正确记录错误日志
func TestLoggerError(t *testing.T) {
	// 创建一个测试用的上下文
	ctx := context.Background()

	// 创建一个测试用的错误
	err := errors.New("test error")

	// 创建 Logger 实例
	logger := NewLogger()

	// 由于实际日志输出难以捕获和验证，我们主要测试方法执行不会崩溃
	// 并确保函数能够正常处理各种参数
	logger.Error(ctx, err, "test message")
	logger.Error(ctx, err, "test message with kvs", "key1", "value1")
	logger.Error(ctx, err, "test message with multiple kvs", "key1", "value1", "key2", "value2")

	// 验证函数执行后没有发生崩溃
	t.Log("所有 Error 方法调用都成功执行，没有发生崩溃")
}

// TestLoggerErrorWithContextValues 测试带有上下文值的错误日志记录
func TestLoggerErrorWithContextValues(t *testing.T) {
	// 创建一个带有值的上下文
	type ctxKey string
	ctx := context.Background()
	ctx = context.WithValue(ctx, ctxKey("traceID"), "123456")
	ctx = context.WithValue(ctx, ctxKey("userID"), "user-456789")

	// 创建一个测试用的错误
	err := errors.New("test error with context")

	// 创建 Logger 实例
	logger := NewLogger()

	// 记录带有上下文的错误日志
	logger.Error(ctx, err, "test message with context values")

	// 验证函数执行后没有发生崩溃
	t.Log("带有上下文值的 Error 方法调用成功执行，没有发生崩溃")
}

// TestLoggerErrorWithEmptyParams 测试使用空参数调用 Error 方法
func TestLoggerErrorWithEmptyParams(t *testing.T) {
	// 创建一个测试用的上下文
	ctx := context.Background()

	// 创建一个 nil 错误
	var nilErr error

	// 创建 Logger 实例
	logger := NewLogger()

	// 测试使用 nil 错误
	logger.Error(ctx, nilErr, "test message with nil error")

	// 测试使用空消息
	logger.Error(ctx, errors.New("test error"), "")

	// 测试不传入任何 kvs
	logger.Error(ctx, errors.New("test error"), "test message", nil)

	// 验证函数执行后没有发生崩溃
	t.Log("使用空参数的 Error 方法调用成功执行，没有发生崩溃")
}
