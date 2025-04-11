package tracing_test

import (
	"context"
	"testing"

	"github.com/kakabei/kfgolib/logx/tracing"
)

func NewTestConfig(filename string) tracing.Config {
	return tracing.Config{
		EnableFile:       true,
		Filename:         filename,
		MaxSize:          1,
		MaxBackups:       1,
		LocalTime:        true,
		Compress:         false,
		EnableConsole:    false,
		EnableCaller:     true,
		EnableSourceIP:   true,
		EnablePID:        true,
		FileLevel:        "debug",
		ConsoleLevel:     "warn",
		FileEncodeing:    "json",
		ConsoleEncodeing: "console",
		AppName:          "logx_test",
		SourceEth:        "eth0",
		DisableTraceID:   false,
	}
}

func BenchmarkLogger_Debug(b *testing.B) {
	ctx := context.Background()
	logger := tracing.NewLogger(NewTestConfig("./logs/debug.log"))
	tracing.ReplaceLogger(logger)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tracing.Debug(ctx, "[1] test tracing.Debug", 1234, 1234.567, "test")
	}
	b.StopTimer()
}

func BenchmarkLogger_Debugf(b *testing.B) {
	ctx := context.Background()
	logger := tracing.NewLogger(NewTestConfig("./logs/debugf.log"))
	tracing.ReplaceLogger(logger)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tracing.Debugf(ctx, "[2] test tracing.Debugf %v %v %v", 1234, 1234.567, "test")
	}
	b.StopTimer()
}
