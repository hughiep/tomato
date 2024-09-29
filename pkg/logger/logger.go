package logger

import "go.uber.org/zap"

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

// Logger is a global logger
var Logger = zap.L()
