package mlog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Options_Validate 测试 Options 的 Validate 方法。
func Test_Options_Validate(t *testing.T) {
	opts := Options{
		Level:            "test",
		Format:           "test",
		EnableColor:      true,
		DisableCaller:    false,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	errs := opts.Validate()
	expected := `[unrecognized level: "test" not a valid log format: "test"]`
	assert.Equal(t, expected, fmt.Sprintf("%s", errs))
}
