package main

import (
	zap2 "LearnGo/pkg/zap"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Start!")
	logger, err := zap2.ConsoleConfig("debug")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	logger2, err := zap2.JSONConfig("debug")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	logger.Info("info", zap.String("package", "main"))
	logger.Debug("dubug")
	logger.Warn("warn")
	logger.Error("error")
	//logger.Fatal("fatal")

	fmt.Println("\n")

	logger2.Info("info", zap.String("package", "main"))
	logger2.Debug("dubug")
	logger2.Warn("warn")
	logger2.Error("error")
	logger2.Fatal("fatal")
	fmt.Println("Finish!")
}
