package logger

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

type ZapOption func(o *zapOption)

type zapOption struct {
    encoder zapcore.EncoderConfig
    level   zap.AtomicLevel
    opts    []zap.Option
    ws      []zapcore.WriteSyncer
}

func Encoder(encoder zapcore.EncoderConfig) ZapOption {
    return func(o *zapOption) { o.encoder = encoder }
}

func Level(level zap.AtomicLevel) ZapOption {
    return func(o *zapOption) { o.level = level }
}

func Option(opts ...zap.Option) ZapOption {
    return func(o *zapOption) { o.opts = opts }
}

func WriteSyncer(ws ...zapcore.WriteSyncer) ZapOption {
    return func(o *zapOption) { o.ws = ws }
}
