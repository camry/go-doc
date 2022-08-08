package logger

import (
    "fmt"
    
    "github.com/camry/g/glog"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

// ZapLogger 实现 log.Logger
type ZapLogger struct {
    log  *zap.Logger
    Sync func() error
}

// NewZapLogger 返回 zap 日志记录器。
func NewZapLogger(opts ...ZapOption) *ZapLogger {
    opt := &zapOption{}
    for _, o := range opts {
        o(opt)
    }
    core := zapcore.NewCore(
        zapcore.NewConsoleEncoder(opt.encoder),
        zapcore.NewMultiWriteSyncer(opt.ws...),
        opt.level,
    )
    zapLogger := zap.New(core, opt.opts...)
    return &ZapLogger{log: zapLogger, Sync: zapLogger.Sync}
}

// Log 实现 log.Logger 接口。
func (l *ZapLogger) Log(level glog.Level, keyvals ...interface{}) error {
    if len(keyvals) == 0 || len(keyvals)%2 != 0 {
        l.log.Warn(fmt.Sprint("keyvals must appear in pairs: ", keyvals))
        return nil
    }
    // Zap.Field 出现 keyvals 对时使用。
    var data []zap.Field
    for i := 0; i < len(keyvals); i += 2 {
        data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
    }
    switch level {
    case glog.LevelDebug:
        l.log.Debug("", data...)
    case glog.LevelInfo:
        l.log.Info("", data...)
    case glog.LevelWarn:
        l.log.Warn("", data...)
    case glog.LevelError:
        l.log.Error("", data...)
    }
    return nil
}
