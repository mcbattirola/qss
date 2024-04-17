package qss

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, ConfigFileName), nil
}

func parseConfigFile(config *Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	fmt.Printf("config file path is %s\n", path)

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("config file not found at %s\n", path)
			return nil
		}
		return err
	}
	fmt.Println("reading config file")

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "=")
		if len(split) >= 2 {
			k := split[0]
			v := split[1]

			// ignore lintes starting with #
			if len(k) > 0 && k[0] == '#' {
				continue
			}

			switch k {
			case "font-size":
				size, err := strconv.Atoi(v)
				if err != nil {
					fmt.Printf("could not parse 'font-size': %s\n", err)
				}
				config.FontSize = size
			case "show-help":
				if v == "true" {
					config.ShowHelp = true
				}
			case "save-path":
				config.FilePath = v
			default:
				fmt.Printf("unknown config key %s\n", k)
			}
		}
	}

	return nil
}
