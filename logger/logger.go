package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level      string `json:"level"`
	Encoding   string `json:"encoding"`
	Caller     string `json:"caller"`
	Stacktrace bool   `json:"stacktrace"`
}

func MustNewLogger(cfg *Config) *zap.Logger {
	logger, err := NewLogger(cfg)
	if err != nil {
		panic(err)
	}
	return logger
}
func NewLogger(cfg *Config) (*zap.Logger, error) {
	level := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	development := false

	switch cfg.Level {
	case "debug", "":
		level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		development = true
	case "info":
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "fatal":
		level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
	case "panic":
		level = zap.NewAtomicLevelAt(zapcore.PanicLevel)
	}

	encodeCaller := zapcore.FullCallerEncoder
	disableCaller := false
	switch cfg.Caller {
	case "short":
		encodeCaller = zapcore.ShortCallerEncoder
	case "disable":
		disableCaller = true
	}

	zapEncoderConfig := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "ts",
		NameKey:       "logger",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		//EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		//EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeCaller: encodeCaller,
	}
	if cfg.Encoding == "" {
		cfg.Encoding = "console"
	}
	zapConfig := zap.Config{
		Level:             level,
		Development:       development,
		DisableCaller:     disableCaller,
		DisableStacktrace: !cfg.Stacktrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100000,
			Thereafter: 100000,
		},
		Encoding:         cfg.Encoding,
		EncoderConfig:    zapEncoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}
	return zapLogger, nil
}
