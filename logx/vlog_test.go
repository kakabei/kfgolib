package logx

import (
	"context"
	"testing"

	"github.com/kakabei/kfgolib/logx/tracing"
)

func TestVLog(t *testing.T) {
	Info("test logx.Debug")
	tracing.Info(context.Background(), "test tracing.Debug")
}
