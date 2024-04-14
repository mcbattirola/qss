package qss

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Key struct {
	desc   string
	key    int32
	keyStr string
	fn     func(qss *App)
}

var keys = []Key{
	{
		desc:   "Show/hide help",
		key:    rl.KeyH,
		keyStr: "h",
		fn: func(qss *App) {
			qss.showHelp = !qss.showHelp
		},
	},
	{
		desc:   "close",
		key:    rl.KeyEscape,
		keyStr: "esc",
		fn: func(qss *App) {
			// handled by raylib
		},
	},
}

func (qss *App) drawHelp() {
	currX := int32(rl.GetScreenWidth()/2) - 300
	// if negative values provided, subtract amount from end of screen
	// if currX < 0 {
	// 	currX = int32(rl.GetScreenWidth()) + qss.config.HelpX
	// }

	currY := int32(rl.GetScreenHeight() / 2)
	// if currY < 0 {
	// 	currY = int32(rl.GetScreenHeight()) + qss.config.HelpY
	// }

	qss.drawText("Click and drag", currX, currY)
	currY += int32(qss.config.FontSize)
	qss.drawText("Right click to print whole screen", currX, currY)
	currY += int32(qss.config.FontSize)
	qss.drawText("During selection, right click or esc to cancel", currX, currY)
}
