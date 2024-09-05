package logger

import "go.uber.org/zap"

func New() *zap.SugaredLogger {
	// Logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Infow("Logger initialized")

	return sugar
}
