package initial

import (
	"github.com/ZLinFeng/play-server/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

func InitLog(c *config.LogConfig) *zap.Logger {
	err := c.Check()
	if err != nil {
		panic(err)
	}
	cores := make([]zapcore.Core, 0)

	if c.StdLog {
		cores = append(cores, ConsoleCore())
	}

	if c.FileLog {
		cores = append(cores, InfoFileCore(c))
		cores = append(cores, ErrorFileCore(c))
	}

	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()
	return logger
}

func ConsoleCore() zapcore.Core {
	encoderConfig := LogEncoderConfig()
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	consoleWriteSyncer := zapcore.AddSync(os.Stdout)
	core := zapcore.NewCore(consoleEncoder, consoleWriteSyncer, zapcore.DebugLevel)

	return core
}

func InfoFileCore(c *config.LogConfig) zapcore.Core {
	encoderConfig := LogEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	infoWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename: filepath.Join(c.Dir, "info.log"),
		MaxSize:  c.FileSizeMb,
		MaxAge:   c.RetentionDays,
		Compress: true,
	})
	core := zapcore.NewCore(consoleEncoder, infoWriteSyncer, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= getLevel(c.Level) && lvl < zapcore.ErrorLevel
	}))
	return core
}

func ErrorFileCore(c *config.LogConfig) zapcore.Core {
	encoderConfig := LogEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	errorWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename: filepath.Join(c.Dir, "error.log"),
		MaxSize:  c.FileSizeMb,
		MaxAge:   c.RetentionDays,
		Compress: true,
	})
	core := zapcore.NewCore(consoleEncoder, errorWriteSyncer, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	}))
	return core
}

func LogEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:     "time",
		NameKey:     "name",
		LevelKey:    "level",
		CallerKey:   "caller",
		MessageKey:  "message",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
}

func getLevel(level string) zapcore.Level {
	switch level {
	case config.INFO:
		return zap.InfoLevel
	case config.DEBUG:
		return zap.DebugLevel
	default:
		return zap.InfoLevel
	}
}
