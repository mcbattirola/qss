package main

import (
	"os"

	"github.com/mcbattirola/qss/pkg/logger"
	"github.com/mcbattirola/qss/pkg/qss"
)

func main() {
	config, err := qss.ReadConfig()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := qss.New(config)
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
