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
	Selected
)

type App struct {
	state                   State
	recInitPos, recFinalPos rl.Vector2
	done                    bool
}

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
	qss := App{
		state:       Idle,
		recInitPos:  rl.Vector2{},
		recFinalPos: rl.Vector2{},
	}

	for !rl.WindowShouldClose() && !qss.done {
		// ----- start texture drawing
		rl.BeginTextureMode(target)
		rl.ClearBackground(rl.Blank)

		switch qss.state {
		case Idle:
			qss.handleIdle()
		case Selecting:
			qss.handleSelecting()
		case Selected:
			qss.handleSelected()
		default:
			fmt.Printf("unknown state %d\n", qss.state)
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
	fmt.Printf("finished")
	os.Exit(0)
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

func (qss *App) handleIdle() {
	// handle input
	// first clicked, set as initial pos
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if !(qss.state == Selecting) {
			qss.state = Selecting
		}
		qss.recInitPos = rl.GetMousePosition()
	}

	// draw instructions, UI etc
	if qss.state != Selecting {
		rl.DrawText(getCurrentTimeStr(), 10, 10, 20, rl.White)
	}
}

func (qss *App) handleSelecting() {
	// release, store final pos and set not drawing
	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		qss.state = Selected
		qss.recFinalPos = rl.GetMousePosition()

		// return here so we don't  draw the rectangle this frame
		return
	}

	// being pressed, set final position to current mouse location
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		qss.recFinalPos = rl.GetMousePosition()
	}

	// draw
	rectSize := getRectSize(qss.recInitPos, qss.recFinalPos)

	// Determine the top-left corner of the rectangle
	topLeftX := int32(math.Min(float64(qss.recInitPos.X), float64(qss.recFinalPos.X)))
	topLeftY := int32(math.Min(float64(qss.recInitPos.Y), float64(qss.recFinalPos.Y)))

	rl.DrawRectangleLines(topLeftX, topLeftY, int32(rectSize.X), int32(rectSize.Y), drawColor)
}

func (qss *App) handleSelected() {
	// take screen shot
	topLeftX := float32(math.Min(float64(qss.recInitPos.X), float64(qss.recFinalPos.X)))
	topLeftY := float32(math.Min(float64(qss.recInitPos.Y), float64(qss.recFinalPos.Y)))
	botRightX := float32(math.Max(float64(qss.recInitPos.X), float64(qss.recFinalPos.X)))
	botRightY := float32(math.Max(float64(qss.recInitPos.Y), float64(qss.recFinalPos.Y)))

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
	qss.done = true
	qss.state = Idle
}
