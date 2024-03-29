package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	monitor := glfw.GetPrimaryMonitor().GetVideoMode()
	window, err := glfw.CreateWindow(monitor.Width, monitor.Height, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.TransparentFramebuffer, 1)
	glfw.WindowHint(glfw.TransparentFramebuffer, 1)
	window.SetAttrib(glfw.Decorated, 0)
	window.SetPos(0, 0)
	window.SetOpacity(0)
	window.MakeContextCurrent()

	for !window.ShouldClose() {
		fmt.Println("ok")
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
