package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// currentMonitor := rl.GetCurrentMonitor()

	// open window
	rl.SetConfigFlags(rl.FlagWindowTransparent)
	rl.InitWindow(0, 0, "printit")
	defer rl.CloseWindow()
	rl.SetWindowPosition(
		0, 0,
		// rl.GetMonitorWidth(currentMonitor)/2-int(screenWidth),
		// rl.GetMonitorHeight(currentMonitor)/2-int(screenHeight),
	)
	rl.SetWindowState(rl.FlagWindowUndecorated)

	var screenWidth int32 = int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
	var screenHeight int32 = 1080 - 1
	rl.SetWindowSize(int(screenWidth), int(screenHeight))
	fmt.Printf("screenHeight is %d\n", screenHeight)

	target := rl.LoadRenderTexture(screenWidth, screenHeight)
	defer rl.UnloadRenderTexture(target)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// draw texture
		rl.BeginTextureMode(target)
		rl.ClearBackground(rl.Blank)
		mouse := rl.GetMousePosition()
		rl.DrawRectangle(int32(mouse.X), int32(mouse.Y), 400, 100, rl.Color{
			R: 255,
			G: 0,
			B: 0,
			A: 192,
		})
		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Blank)
		// for some reason we have to flip the Y here
		rl.DrawTexturePro(target.Texture, rl.Rectangle{X: 0.0, Y: 0.0, Width: float32(screenWidth), Height: float32(screenHeight) * -1}, rl.Rectangle{X: 0.0, Y: 0.0, Width: float32(screenWidth), Height: float32(screenHeight)},
			rl.Vector2{X: 0, Y: 0}, 0.0, rl.White)
		rl.EndDrawing()
	}
}
