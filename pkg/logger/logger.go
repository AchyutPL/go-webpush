package logger

import (
	"fmt"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

// Define a custom encoder
type ColorEncoder struct {
	zapcore.Encoder
}

func (ce *ColorEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buffer, err := ce.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}

	color := ce.getColor(entry.Level)
	coloredLine := fmt.Sprintf("%s%s%s", color, buffer.String(), "\033[0m")

	buffer.Reset()
	buffer.AppendString(coloredLine)
	return buffer, nil
}

func (ce *ColorEncoder) getColor(level zapcore.Level) string {
	switch level {
	case zapcore.DebugLevel:
		return "\033[36m" // Cyan
	case zapcore.InfoLevel:
		return "\033[34m" // Blue
	case zapcore.WarnLevel:
		return "\033[33m" // Yellow
	case zapcore.ErrorLevel:
		return "\033[31m" // Red
	case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		return "\033[35m" // Magenta
	default:
		return "\033[0m" // Reset
	}
}

var Log *zap.Logger

func InitializeLogger() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Create a custom encoder with color
	encoder := &ColorEncoder{zapcore.NewConsoleEncoder(encoderConfig)}

	// New core with colorable stdout
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	)

	Log = zap.New(core)
	defer Log.Sync() // Flushes buffer, if any
}
