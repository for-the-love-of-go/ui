package ui

import (
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func GetRendererAndWindow(h, w int32) (*sdl.Window, *sdl.Renderer, func()) {
	runtime.LockOSThread()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	if err := ttf.Init(); err != nil {
		panic(err)
	}

	window, r, err := sdl.CreateWindowAndRenderer(w, h, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	window.SetResizable(true)
	window.SetBordered(true)

	return window, r,
		func() {
			window.Destroy()
			sdl.Quit()
			ttf.Quit()
			runtime.UnlockOSThread()
		}

}
