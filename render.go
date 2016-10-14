package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)

	const W=800
	const H=600

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		W, H, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
		return
	}
	defer renderer.Destroy()

	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	for i:=int32(0);i<700;i++ {
		if i>0 {
			rect := sdl.Rect{i-1, 0, 10, 10}
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.FillRect(&rect)
		}

		rect := sdl.Rect{i, 0, 10, 10}
		renderer.SetDrawColor(0, 255, 255, 255)
		renderer.FillRect(&rect)
		renderer.Present()
	}


	sdl.Delay(2000)

}
