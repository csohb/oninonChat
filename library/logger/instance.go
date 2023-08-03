package logger

import "go.uber.org/zap"

var Instance *zap.Logger

func init() {
	Instance, _ = zap.NewProduction()
}
