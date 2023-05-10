package log

import (
	"ecommerce/pkg/libs/errors"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

// Initialize function
// func Initialize(devMode bool, logFile string) {
func Initialize(devMode bool) {

	if logger != nil {
		panic("Log already initialized")
	}

	var err error

	var config zap.Config
	if devMode {
		config = zap.Config{
			Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
			Development:       true,
			Encoding:          "console",
			EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
			OutputPaths:       []string{"stderr"},
			ErrorOutputPaths:  []string{"stderr"},
			DisableStacktrace: true,
			DisableCaller:     true,
		}
	} else {
		config = zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
			Development: false,
			Sampling: &zap.SamplingConfig{
				Initial:    100,
				Thereafter: 100,
			},
			Encoding:          "json",
			EncoderConfig:     zap.NewProductionEncoderConfig(),
			OutputPaths:       []string{"stderr"},
			ErrorOutputPaths:  []string{"stderr"},
			DisableStacktrace: true,
			DisableCaller:     true,
		}
	}

	logger, err = config.Build()
	if err != nil {
		panic(err)
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// Debug function
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// Info function
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Warn function
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// Error function
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

// Fatal function
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

// Panic function
func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

// String string field
func String(key string, value string) zap.Field {
	return zap.String(key, value)
}

// Int int field
func Int(key string, value int) zap.Field {
	return zap.Int(key, value)
}

// Duration duration field
func Duration(key string, value time.Duration) zap.Field {
	return zap.Duration(key, value)
}

// Bool bool field
func Bool(key string, value bool) zap.Field {
	return zap.Bool(key, value)
}

// Float float field
func Float(key string, value float64) zap.Field {
	return zap.Float64(key, value)
}

// Time time field
func Time(key string, value time.Time) zap.Field {
	return zap.Time(key, value)
}

// Namespace namespace field
func Namespace(namespace string) zap.Field {
	return zap.Namespace(namespace)
}

// Stack stack field
func Stack(key string) zap.Field {
	return zap.Stack(key)
}

// Any any field
func Any(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

// Close close function
func Close() {
	logger.Info("Logger finalizado")
	logger = nil
}

func Error2(msg errors.Operation, err error, fields ...zap.Field) {

	e, ok := err.(*errors.Error)
	if !ok {
		var f []zap.Field
		if err.Error() != "" {
			f = []zap.Field{zap.Error(err)}
		}
		f = append(f, fields...)
		logger.Error(string(msg), f...)
		return
	}

	var f []zap.Field
	if e.Error() != "" {
		f = []zap.Field{zap.Error(e)}
	}
	f = append(f, zap.String("kind", errors.Kind(e).String()))
	f = append(f, zap.Any("stacktrace", errors.Operations(e)))
	f = append(f, fields...)
	f = append(f, errors.Fields(e)...)

	switch errors.Level(e) {
	case errors.LevelDebug:
		logger.Debug(string(msg), f...)
	case errors.LevelInfo:
		logger.Info(string(msg), f...)
	case errors.LevelWarn:
		logger.Warn(string(msg), f...)
	case errors.LevelError:
		logger.Error(string(msg), f...)
	default:
		logger.Error(string(msg), f...)
	}
}
