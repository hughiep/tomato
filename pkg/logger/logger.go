package logger

import (
	"go.uber.org/zap"
)

func Init() error {
	logger := zap.Must(zap.NewProduction())

	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	return nil
}
