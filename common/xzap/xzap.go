package xzap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

type Config struct {
	Format   string
	Level    string
	Color    bool
	FilePath string
	File     bool
	Terminal bool
}

func New(config Config) *zap.Logger {
	var coreConfigs = make([]zapConfig, 0)
	var cors = make([]zapcore.Core, 0)

	var encoder func(cfg zapcore.EncoderConfig) zapcore.Encoder

	switch config.Format {
	case "json":
		encoder = zapcore.NewJSONEncoder
	default:
		encoder = zapcore.NewConsoleEncoder
	}
	var needColor = config.Color
	level := transLevel(config.Level)

	if config.File {
		fileInfoConfig := newZapConfig().
			setEncoder(false, encoder).
			setFileWriteSyncer(config.FilePath + "info.log").
			setLevelEnabler(level).
			getConfig()
		fileInfoCore := fileInfoConfig.getCore()

		fileErrorConfig := newZapConfig().
			setEncoder(needColor, encoder).
			setFileWriteSyncer(config.FilePath + "error.log").
			setLevelEnabler(level).
			getConfig()
		fileErrorCore := fileErrorConfig.getCore()

		coreConfigs = append(coreConfigs, fileInfoConfig, fileErrorConfig)
		cors = append(cors, fileInfoCore, fileErrorCore)
	}

	if config.Terminal {
		consoleInfoConfig := newZapConfig().
			setEncoder(needColor, encoder).
			setStdOutWriteSyncer().
			setLevelEnabler(level).
			getConfig()
		consoleInfoCore := consoleInfoConfig.getCore()

		coreConfigs = append(coreConfigs, consoleInfoConfig)
		cors = append(cors, consoleInfoCore)
	}

	logger := makeZapLogger(cors, zap.AddCallerSkip(1), zap.WithCaller(true))
	defer logger.Sync()
	return logger
}

func makeZapLogger(cors []zapcore.Core, options ...zap.Option) *zap.Logger {
	core := zapcore.NewTee(cors...)
	return zap.New(core, options...)

}

type zapConfig struct {
	core             zapcore.Core
	encoder          zapcore.Encoder
	writeSyncerSlice []zapcore.WriteSyncer
	levelEnabler     zapcore.LevelEnabler
}

func newZapConfig() *zapConfig {
	return &zapConfig{
		writeSyncerSlice: make([]zapcore.WriteSyncer, 0),
	}
}

func (z *zapConfig) getConfig() zapConfig {
	return *z
}

// 定制core
func (z *zapConfig) getCore() zapcore.Core {
	z.core = zapcore.NewCore(z.encoder, zapcore.NewMultiWriteSyncer(z.writeSyncerSlice...), z.levelEnabler)
	return z.core
}

func (z *zapConfig) setEncoder(needColor bool, encoder func(cfg zapcore.EncoderConfig) zapcore.Encoder) *zapConfig {
	encodeLevel := zapcore.CapitalLevelEncoder
	if needColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	z.encoder = encoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime:     newTimeEncoder(),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	})

	return z
}

func (z zapConfig) getEncoder() zapcore.Encoder {
	return z.encoder
}

func (z *zapConfig) setFileWriteSyncer(logFilePath string) *zapConfig {
	lumberWriteSyncer := &lumberjack.Logger{
		MaxSize:    1024,  // megabytes
		MaxBackups: 7,     //最多备份文件数量
		MaxAge:     28,    // days
		Compress:   false, //Compress确定是否应该使用gzip压缩已旋转的日志文件。默认值是不执行压缩。
	}

	z.writeSyncerSlice = append(z.writeSyncerSlice, zapcore.AddSync(lumberWriteSyncer))
	return z
}

func (z zapConfig) getWriteSyncers() []zapcore.WriteSyncer {
	return z.writeSyncerSlice
}

func (z *zapConfig) setStdOutWriteSyncer() *zapConfig {
	z.writeSyncerSlice = append(z.writeSyncerSlice, zapcore.AddSync(os.Stdout))
	return z
}

func (z *zapConfig) setLevelEnabler(enabler zapcore.Level) *zapConfig {
	z.levelEnabler = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= enabler
	})

	return z
}

func newTimeEncoder() zapcore.TimeEncoder {
	return func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006/1/2 15:04:05.000"))
	}
}

func transLevel(level string) zapcore.Level {
	var levels = []string{
		"debug",
		"info",
		"warn",
		"error",
		"panic",
	}
	for index, value := range levels {
		if level == value {
			return zapcore.Level(index - 1)
		}
	}

	return zapcore.DebugLevel
}
