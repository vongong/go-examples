package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	var logger = logrus.New()
	defer logger.Writer().Close()
	logger.Out = os.Stdout
	// You could set this to any `io.Writer` such as a file
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	//  logger.Out = file
	// } else {
	//  log.Info("Failed to log to file, using default stderr")
	// }
	logger.Info("Hello World")
	logger.Warn("This is a warning")

}
