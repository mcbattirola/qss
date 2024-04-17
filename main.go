package main

import (
	"fmt"
	"os"

	"github.com/mcbattirola/qss/pkg/qss"
)

func main() {
	config, err := qss.ReadConfig()
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	app := qss.New(config)
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
