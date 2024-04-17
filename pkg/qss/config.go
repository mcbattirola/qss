package qss

import (
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ConfigFileName = ".qss.conf"

	// Defaults
	DefaultFontSize = 24
	DefaultShowSize = true
	DefaultShowHelp = true
	DefaultFileDir  = "Pictures"
)

var (
	DefaultFontColor = rl.Lime
)

func DefaultFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, DefaultFileDir), nil
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
	filePath, err := DefaultFilePath()
	if err != nil {
		return Config{}, err
	}

	// defaults
	config := Config{
		FontSize:  DefaultFontSize,
		FontColor: DefaultFontColor,
		ShowSize:  DefaultShowSize,
		ShowHelp:  DefaultShowHelp,
		FilePath:  filePath,
	}

	// overwrite defaults with user configs
	err = parseConfigFile(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
