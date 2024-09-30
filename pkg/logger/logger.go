package logger

import (
	"tomato/internal/config"

	"go.uber.org/zap"
)

func Init(cfg *config.Config) error {
	var zcfg zap.Config

	logger, err := zcfg.Build()
	if err != nil {
		return err
	}
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	return nil
}
