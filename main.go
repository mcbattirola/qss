package main

import (
	"fmt"
	"os"

	"github.com/mcbattirola/qss/pkg/qss"
)

func main() {
	defaultConfig, err := qss.DefaultConfig()
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}

	app := qss.New(defaultConfig)
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
