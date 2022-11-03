package main

import "go.uber.org/zap"

func main() {
	logger, err := zap.NewProductionConfig().Build()
	if err != nil {
		return
	}
	logger.Info("ddd")
}
