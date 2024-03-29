package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"strings"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kbinani/screenshot"
)

type State int8

const (
	Idle State = iota
	Selecting
)

var drawColor = rl.Color{
	R: 255,
	G: 0,
	B: 0,
	A: 192,
}

func main() {
	// open window
	rl.SetConfigFlags(rl.FlagWindowTransparent)
	rl.InitWindow(0, 0, "printit")
	defer rl.CloseWindow()
	rl.SetWindowPosition(0, 0)
	rl.SetWindowState(rl.FlagWindowUndecorated |
		rl.FlagBorderlessWindowedMode |
		rl.FlagWindowAlwaysRun)

	var screenWidth int32 = int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
	var screenHeight int32 = int32(rl.GetMonitorHeight(rl.GetCurrentMonitor())) + 1
	rl.SetWindowSize(int(screenWidth), int(screenHeight))

	target := rl.LoadRenderTexture(screenWidth, screenHeight)
	defer rl.UnloadRenderTexture(target)

	rl.SetTargetFPS(60)
	rl.SetMouseCursor(3) // cross

	// state
	recInitPos := rl.Vector2{}
	recFinalPos := rl.Vector2{}
	state := Idle

	for !rl.WindowShouldClose() {
		// ----- start texture drawing
		rl.BeginTextureMode(target)
		rl.ClearBackground(rl.Blank)

		// handle input
		mousePos := rl.GetMousePosition()
		// first clicked, set as initial pos
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			if !(state == Selecting) {
				state = Selecting
			}
			recInitPos = mousePos
		}
		// being pressed, set final position
		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			recFinalPos = mousePos
		}
		// release, store final pos and set not drawing
		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			if state == Selecting {
				state = Idle
				recFinalPos = mousePos

				// take screen shot
				topLeftX := float32(math.Min(float64(recInitPos.X), float64(recFinalPos.X)))
				topLeftY := float32(math.Min(float64(recInitPos.Y), float64(recFinalPos.Y)))
				botRightX := float32(math.Max(float64(recInitPos.X), float64(recFinalPos.X)))
				botRightY := float32(math.Max(float64(recInitPos.Y), float64(recFinalPos.Y)))

				captureAndSaveAsPNG(image.Rectangle{
					Min: image.Point{
						X: int(topLeftX),
						Y: int(topLeftY),
					},
					Max: image.Point{
						X: int(botRightX),
						Y: int(botRightY),
					},
				}, fmt.Sprintf("screenshot-%s.png", getCurrentTimeStr()))
				// after taking screenshot, quit
				rl.CloseWindow()
			}
		}

		// draw instructions etc
		if state != Selecting {
			rl.DrawText(getCurrentTimeStr(), 10, 10, 20, rl.White)
		}

		// draw rec
		if state == Selecting {
			// Calculate the width and height as the absolute difference between the initial and final positions
			rectSize := getRectSize(recInitPos, recFinalPos)

			// Determine the top-left corner of the rectangle
			topLeftX := int32(math.Min(float64(recInitPos.X), float64(recFinalPos.X)))
			topLeftY := int32(math.Min(float64(recInitPos.Y), float64(recFinalPos.Y)))

			rl.DrawRectangleLines(topLeftX-1, topLeftY-1, int32(rectSize.X)+2, int32(rectSize.Y)+2, drawColor)
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

// getRectSize returns the size of the rect
func getRectSize(recInitPos, recFinalPos rl.Vector2) rl.Vector2 {
	width := float32(math.Abs(float64(recFinalPos.X - recInitPos.X)))
	height := float32(math.Abs(float64(recFinalPos.Y - recInitPos.Y)))
	return rl.Vector2{X: width, Y: height}
}

// captureAndSaveAsPNG captures the area inside boungs and
// save it.
// fileName must include '.png' sufix
func captureAndSaveAsPNG(bounds image.Rectangle, fileName string) {
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		// TODO: warn user
		fmt.Printf("error taking screenshot:%s\n", err.Error())
		return
	}
	// TODO: add a sufix is file already exists
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("error creating image:%s\n", err.Error())
		return
	}
	defer file.Close()
	png.Encode(file, img)
	fmt.Printf("%s saved\n", fileName)
}

func getCurrentTimeStr() string {
	return strings.ReplaceAll(time.Now().UTC().Format(time.RFC3339), ":", "_")
}
