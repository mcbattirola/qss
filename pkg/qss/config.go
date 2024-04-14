package qss

import (
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Defaults
const (
	DefaultFontSize = 18
	DefaultShowSize = true
	DefaultShowHelp = true
	DefaultHelpX    = -400
	DefaultHelpY    = -400
	DefaultFileDir  = "Pictures"
)

var (
	DefaultFontColor = rl.Red
)

type Config struct {
	FontSize  int
	FontColor rl.Color
	ShowSize  bool
	ShowHelp  bool
	HelpX     int32
	HelpY     int32
	FilePath  string
}

func DefaultConfig() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	return Config{
		FontSize:  DefaultFontSize,
		FontColor: DefaultFontColor,
		ShowSize:  DefaultShowSize,
		ShowHelp:  DefaultShowHelp,
		HelpX:     DefaultHelpX,
		HelpY:     DefaultHelpY,
		FilePath:  path.Join(homeDir, DefaultFileDir),
	}, nil
}
