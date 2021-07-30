package logger

import (
	"go.uber.org/zap"
)

// Logger struct
// type Logger struct{}

// ILogger interface
// type ILogger interface {
// 	Info(args ...interface{})
// 	Infof(template string, args ...interface{})
// 	Error(args ...interface{})
// 	Errorf(template string, args ...interface{})
// 	Panic(args ...interface{})
// 	Panicf(template string, args ...interface{})
// }

// TODO:
// Load Config and read environment

// Info func
func Info(args ...interface{}) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()
	logger.Info(args)
}

// Infof func
func Infof(template string, args ...interface{}) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()
	logger.Infof(template, args)
}

// Infow func
func Infow(msg string, keysAndValue ...interface{}) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()
	logger.Infow(msg, keysAndValue...)
}

// Error func
func Error(args ...interface{}) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()
	logger.Error(args)
}

// Errorf func
func Errorf(template string, args ...interface{}) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()
	logger.Errorf(template, args)
}

// Errorw func
func Errorw(msg string, keysAndValue ...interface{}) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()
	logger.Errorw(msg, keysAndValue...)
}

// Fatal func
func Fatal(args ...interface{}) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()
	logger.Fatal(args)
}

// Fatalf func
func Fatalf(template string, args ...interface{}) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()
	logger.Fatalf(template, args)
}

// Fatalw func
func Fatalw(msg string, keysAndValue ...interface{}) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()
	logger.Fatalw(msg, keysAndValue...)
}

// Panic func
func Panic(args ...interface{}) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()
	logger.Panic(args)
}

// Panicf func
func Panicf(template string, args ...interface{}) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()
	logger.Panicf(template, args)
}

// Panicw func
func Panicw(msg string, keysAndValue ...interface{}) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()
	logger.Panicw(msg, keysAndValue...)
}
