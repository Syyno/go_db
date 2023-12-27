package main

import (
	"bufio"
	"fmt"
	"mydb/internal/app"
	"os"
	"strings"

	"go.uber.org/zap"
)

func main() {
	//todo: add a config that may be used in app
	logger := Logger()
	db := app.New(logger)

	reader := bufio.NewReader(os.Stdin)

	for {
		logger.Info("Enter command: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			logger.Info("An error occured while reading input. Please try again")
			continue
		}
		input = strings.TrimSuffix(input, "\n")
		res, err := db.Database.Handle(input)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Result: %v\n", res)
		}
	}
}

// todo: make better logging setup
func Logger() *zap.Logger {
	var logger, _ = zap.NewProduction()
	return logger
}
