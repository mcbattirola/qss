package qss

import (
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	configFileName = ".qss.conf"

	// Defaults
	defaultFontSize = 24
	defaultShowSize = true
	defaultShowHelp = true
	defaultFileDir  = "Pictures"
)

var (
	DefaultFontColor = rl.Lime
)

func defaultFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, defaultFileDir), nil
}

type Config struct {
	FontSize  int
	FontColor rl.Color
	ShowSize  bool
	ShowHelp  bool
	HelpX     int32
	HelpY     int32
	FilePath  string
}

// ReadConfig reads the config file if it exists
// and returns a config object
func ReadConfig() (Config, error) {
	filePath, err := defaultFilePath()
	if err != nil {
		return Config{}, err
	}

	// defaults
	config := Config{
		FontSize:  defaultFontSize,
		FontColor: DefaultFontColor,
		ShowSize:  defaultShowSize,
		ShowHelp:  defaultShowHelp,
		FilePath:  filePath,
	}

	// overwrite defaults with user configs
	err = parseConfigFile(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
