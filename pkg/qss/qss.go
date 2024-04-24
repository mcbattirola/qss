package qss

import (
	"errors"
	"fmt"
	"image"
	"math"
	"os"
	"path"
	"runtime"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mcbattirola/qss/pkg/logger"
	"github.com/mcbattirola/qss/pkg/utils"
)

type State int8

const (
	Idle State = iota
	Selecting
	Selected
)

const APP_NAME = "qss"

type App struct {
	state                   State
	recInitPos, recFinalPos rl.Vector2
	done                    bool
	showHelp                bool
	config                  Config
}

func New(config Config) *App {
	return &App{
		state:       Idle,
		recInitPos:  rl.Vector2{},
		recFinalPos: rl.Vector2{},
		showHelp:    config.ShowHelp,

		config: config,
	}
}

func (qss *App) Run() error {
	// Before opening the window, take a screenshot of the screen
	// so we can draw it to our window.
	// This is done to have better compatibility between OSs,
	// since in some desktop environments showing a transparent
	// window don't work.
	screenImg, err := utils.CaptureScreen(rl.GetCurrentMonitor())
	if err != nil {
		return err
	}
	rlImage := rl.NewImageFromImage(screenImg)

	// open window
	rl.SetConfigFlags(rl.FlagWindowTransparent)
	rl.InitWindow(0, 0, APP_NAME)
	defer rl.CloseWindow()
	rl.SetWindowPosition(0, 0)
	rl.SetWindowState(rl.FlagWindowUndecorated |
		rl.FlagBorderlessWindowedMode |
		rl.FlagWindowAlwaysRun)

	// load our background texture
	screenTexture := rl.LoadTextureFromImage(rlImage)
	defer rl.UnloadTexture(screenTexture)

	var screenWidth int32 = int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
	var screenHeight int32 = int32(rl.GetMonitorHeight(rl.GetCurrentMonitor())) + 1
	rl.SetWindowSize(int(screenWidth), int(screenHeight))

	target := rl.LoadRenderTexture(screenWidth, screenHeight)
	defer rl.UnloadRenderTexture(target)

	rl.SetTargetFPS(60)
	rl.SetMouseCursor(3) // cross

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
			logger.Error("unknown state %d\n", qss.state)
			return errors.New("app got into unknown state")
		}

		rl.EndTextureMode()
		// ----- end texture drawing

		// ----- start drawing to screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Blank)
		// draw background texture on Linux
		if runtime.GOOS == "linux" {
			rl.DrawTexture(screenTexture, 0, 0, rl.White)
		}

		// for some reason we have to flip the Y here
		rl.DrawTexturePro(target.Texture, rl.Rectangle{X: 0.0, Y: 0.0, Width: float32(screenWidth), Height: float32(screenHeight) * -1}, rl.Rectangle{X: 0.0, Y: 0.0, Width: float32(screenWidth), Height: float32(screenHeight)},
			rl.Vector2{X: 0, Y: 0}, 0.0, rl.White)
		rl.EndDrawing()
		// ----- end drawing to screen
	}
	return nil
}

func (qss *App) handleIdle() {
	// first clicked, set as initial pos
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if !(qss.state == Selecting) {
			qss.state = Selecting
		}
		qss.recInitPos = rl.GetMousePosition()
	}

	// right click, print whole screen
	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		qss.recInitPos = rl.Vector2{X: 0, Y: 0}
		qss.recFinalPos = rl.Vector2{X: float32(rl.GetScreenWidth()), Y: float32(rl.GetScreenHeight())}
		qss.state = Selected
	}

	if rl.IsKeyPressed(rl.KeyH) {
		qss.showHelp = !qss.showHelp
	}

	// draw instructions, UI etc
	if qss.showHelp {
		qss.drawHelp()
	}
}

func (qss *App) handleSelecting() {
	rl.SetExitKey(0)

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

	// cancel selection on right click
	if rl.IsMouseButtonDown(rl.MouseButtonRight) || rl.IsKeyDown(rl.KeyEscape) {
		rl.SetExitKey(rl.KeyEscape) // set esc to quit again
		qss.state = Idle
		return
	}

	// draw
	rectSize := utils.GetRectSize(qss.recInitPos, qss.recFinalPos)

	// Determine the top-left corner of the rectangle
	topLeftX := int32(math.Min(float64(qss.recInitPos.X), float64(qss.recFinalPos.X)))
	topLeftY := int32(math.Min(float64(qss.recInitPos.Y), float64(qss.recFinalPos.Y)))

	rl.DrawRectangleLines(topLeftX, topLeftY, int32(rectSize.X), int32(rectSize.Y), qss.config.FontColor)

	if qss.config.ShowSize {
		rl.DrawText(
			fmt.Sprintf("%d x %d", int(rectSize.X), int(rectSize.Y)),
			int32(qss.recFinalPos.X)+10,
			int32(qss.recFinalPos.Y)+10,
			int32(qss.config.FontSize),
			qss.config.FontColor)
	}
}

func (qss *App) handleSelected() {
	// take screenshot
	topLeftX := float32(math.Min(float64(qss.recInitPos.X), float64(qss.recFinalPos.X)))
	topLeftY := float32(math.Min(float64(qss.recInitPos.Y), float64(qss.recFinalPos.Y)))
	botRightX := float32(math.Max(float64(qss.recInitPos.X), float64(qss.recFinalPos.X)))
	botRightY := float32(math.Max(float64(qss.recInitPos.Y), float64(qss.recFinalPos.Y)))

	os.MkdirAll(qss.config.FilePath, 0700)

	utils.CaptureAndSaveAsPNG(image.Rectangle{
		Min: image.Point{
			X: int(topLeftX),
			Y: int(topLeftY),
		},
		Max: image.Point{
			X: int(botRightX),
			Y: int(botRightY),
		},
	}, path.Join(qss.config.FilePath, fmt.Sprintf("screenshot-%s.png", utils.GetCurrentTimeStr())))

	// after taking screenshot, quit
	qss.done = true
	qss.state = Idle
}

func (qss *App) drawTextCentered(str string, x, y int32) {
	size := rl.MeasureText(str, int32(qss.config.FontSize))
	rl.DrawText(str, x-(size/2), y, int32(qss.config.FontSize), qss.config.FontColor)
}

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
