package utils

import (
	"image"
	"image/png"
	"math"
	"os"
	"strings"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kbinani/screenshot"
	"github.com/mcbattirola/qss/pkg/logger"
)

// captureAndSaveAsPNG captures the area inside bounds and
// save it.
// fileName must include '.png' sufix
func CaptureAndSaveAsPNG(bounds image.Rectangle, fileName string) error {
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		logger.Error("error taking screenshot:%s\n", err.Error())
		return err
	}
	// TODO: add a sufix is file already exists
	file, err := os.Create(fileName)
	if err != nil {
		logger.Error("error creating image:%s\n", err.Error())
		return err
	}
	defer file.Close()
	png.Encode(file, img)

	return nil
}

func GetCurrentTimeStr() string {
	return strings.ReplaceAll(time.Now().UTC().Format(time.RFC3339), ":", "_")
}

// getRectSize returns the size of the rect
func GetRectSize(recInitPos, recFinalPos rl.Vector2) rl.Vector2 {
	width := float32(math.Abs(float64(recFinalPos.X - recInitPos.X)))
	height := float32(math.Abs(float64(recFinalPos.Y - recInitPos.Y)))
	return rl.Vector2{X: width, Y: height}
}

// CaptureScreen captures the whole screen
func CaptureScreen(display int) (*image.RGBA, error) {
	bounds := screenshot.GetDisplayBounds(display)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return nil, err
	}
	return img, nil
}
