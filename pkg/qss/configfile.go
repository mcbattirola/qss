package qss

import (
	"bufio"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/mcbattirola/qss/pkg/logger"
)

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, configFileName), nil
}

func loadConfig(config *Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	logger.Info("config file path is %s\n", path)

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Error("config file not found at %s\n", path)
			return nil
		}
		return err
	}
	defer file.Close()
	logger.Info("reading config file")

	parseConfigFile(file, config)
	return nil
}

// parseConfigFile parses the content of the pointed file into config.
// In case of an unknown key, it logs the error and don't fail.
// Caller is responsible for closing the file descriptor.
func parseConfigFile(file *os.File, config *Config) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "=")
		if len(split) >= 2 {
			k := split[0]
			v := split[1]

			// ignore lines starting with #
			if len(k) > 0 && k[0] == '#' {
				continue
			}

			k = strings.TrimSpace(k)

			switch k {
			case "font-size":
				size, err := strconv.Atoi(v)
				if err != nil {
					logger.Error("could not parse 'font-size': %s\n", err)
				}
				config.FontSize = size
			case "show-help":
				if v == "true" {
					config.ShowHelp = true
				} else {
					config.ShowHelp = false
				}
			case "save-path":
				config.FilePath = v
			default:
				logger.Warn("unknown config key %s\n", k)
			}
		}
	}
}
