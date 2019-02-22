package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/fortheloveofgo/ui"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	_, r, cleanup := ui.GetRendererAndWindow(800, 1000)
	defer cleanup()

	// w.SetResizable(true)
	// w.Show()

	err := r.SetDrawColor(23, 23, 200, 255)
	if err != nil {
		panic(err)
	}

	go func(r *sdl.Renderer) {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		for {
			time.Sleep(time.Millisecond * 1000 / 60)
			err := r.DrawLine(34, 234, 23, 100)
			if err != nil {
				panic(err)
			}
			r.Present()
		}

	}(r)

	events := make(chan sdl.Event, 2)
	quit := make(chan struct{}, 2)

	go func() {
		for e := range events {
			switch e := e.(type) {
			case *sdl.MouseMotionEvent:
				if e.State == sdl.BUTTON_LEFT {
					r.DrawLine(e.X, e.Y, e.X+e.XRel, e.Y+e.YRel)
				}
			case *sdl.QuitEvent:
				quit <- struct{}{}
			default:
				log.Printf("%#T", e)
			}
		}
	}()

	for {
		select {
		case events <- sdl.WaitEvent():
			// time.Sleep(time.Millisecond)
		case <-quit:
			close(events)
			fmt.Println("Quitting")
			return
		}
	}
	// w.UpdateSurface()
	// w.Raise()
	// w.Show()
	// w.Raise()
	time.Sleep(time.Minute * 10)
}
