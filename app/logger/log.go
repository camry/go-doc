package logger

import (
    "fmt"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
)

func NewAppLogger() *ZapLogger {
    encoder := zap.NewProductionEncoderConfig()
    encoder.EncodeTime = zapcore.ISO8601TimeEncoder
    encoder.EncodeLevel = zapcore.CapitalLevelEncoder
    l := NewZapLogger(
        Level(zap.NewAtomicLevel()),
        Encoder(encoder),
        WriteSyncer(zapcore.AddSync(&lumberjack.Logger{
            Filename:   fmt.Sprintf("%s/app.log", "./logs"),
            MaxSize:    256,
            MaxBackups: 3,
            MaxAge:     28,
            Compress:   true,
        })),
    )
    return l
}
