package qss

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (qss *App) drawHelp() {
	screenWidth := rl.GetScreenWidth()
	currX := int32(screenWidth / 2)

	screenHeight := rl.GetScreenHeight()
	horizontalDisplacement := float32(screenHeight) * 0.25
	currY := int32((screenHeight / 2) - int(horizontalDisplacement))

	qss.drawTextCentered("Click and drag", currX, currY)
	currY += int32(qss.config.FontSize * 2)
	qss.drawTextCentered("Right click to print whole screen", currX, currY)
	currY += int32(qss.config.FontSize * 2)
	qss.drawTextCentered("During selection, right click or esc to cancel", currX, currY)
	currY += int32(qss.config.FontSize * 2)
	qss.drawTextCentered("ESC to quit", currX, currY)
	currY += int32(qss.config.FontSize * 2)
	qss.drawTextCentered("'h' to show/hide this help", currX, currY)
}
