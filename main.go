package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var drawColor = rl.Color{
	R: 255,
	G: 0,
	B: 0,
	A: 192,
}

func main() {
	// currentMonitor := rl.GetCurrentMonitor()

	// open window
	rl.SetConfigFlags(rl.FlagWindowTransparent)
	rl.InitWindow(0, 0, "printit")
	defer rl.CloseWindow()
	rl.SetWindowPosition(0, 0)
	rl.SetWindowState(rl.FlagWindowUndecorated |
		rl.FlagBorderlessWindowedMode |
		// rl.FlagWindowTopmost |
		rl.FlagWindowAlwaysRun)

	var screenWidth int32 = int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
	var screenHeight int32 = int32(rl.GetMonitorHeight(rl.GetCurrentMonitor())) + 1
	rl.SetWindowSize(int(screenWidth), int(screenHeight))
	fmt.Printf("screenHeight is %d\n", screenHeight)

	target := rl.LoadRenderTexture(screenWidth, screenHeight)
	defer rl.UnloadRenderTexture(target)

	rl.SetTargetFPS(60)

	rl.SetMouseCursor(3)

	// state
	recInitPos := rl.Vector2{}
	recFinalPos := rl.Vector2{}
	drawingRec := false

	for !rl.WindowShouldClose() {
		// handle input
		mousePos := rl.GetMousePosition()
		// first clicked, set as initial pos
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			if !drawingRec {
				drawingRec = true
			}
			recInitPos = mousePos
		}
		// being pressed, set final position
		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			recFinalPos = mousePos
		}
		// release, store final pos and set not drawing
		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			if drawingRec {
				drawingRec = false
				recFinalPos = mousePos

				// take screen shot
				// target := rl.LoadRenderTexture(int32(recSize.X), int32(recSize.Y))

			}
		}

		// ----- start texture drawing
		rl.BeginTextureMode(target)
		rl.ClearBackground(rl.Blank)

		// draw rec
		if drawingRec {
			// Calculate the width and height as the absolute difference between the initial and final positions
			width := float32(math.Abs(float64(recFinalPos.X - recInitPos.X)))
			height := float32(math.Abs(float64(recFinalPos.Y - recInitPos.Y)))

			// Determine the top-left corner of the rectangle
			topLeftX := float32(math.Min(float64(recInitPos.X), float64(recFinalPos.X)))
			topLeftY := float32(math.Min(float64(recInitPos.Y), float64(recFinalPos.Y)))

			rl.DrawRectangleV(rl.Vector2{X: topLeftX, Y: topLeftY}, rl.Vector2{X: width, Y: height}, drawColor)
		}
		rl.EndTextureMode()
		// ----- end texture drawing

		// ----- start drawing to screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Blank)
		// for some reason we have to flip the Y here
		rl.DrawTexturePro(target.Texture, rl.Rectangle{X: 0.0, Y: 0.0, Width: float32(screenWidth), Height: float32(screenHeight) * -1}, rl.Rectangle{X: 0.0, Y: 0.0, Width: float32(screenWidth), Height: float32(screenHeight)},
			rl.Vector2{X: 0, Y: 0}, 0.0, rl.White)
		rl.EndDrawing()
		// ----- end drawing to screen
	}
}
