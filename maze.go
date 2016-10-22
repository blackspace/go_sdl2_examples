package main

import (
	"go_sdl2_examples/maze"
	"github.com/veandco/go-sdl2/sdl"
)


func main() {
	w:=5
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
	}
	defer renderer.Destroy()


	mm:=maze.BuildMaze(30)
	renderer.SetViewport(&sdl.Rect{10,10,int32(mm.Len()*w+1),int32(mm.Len()*w+1)})

	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	mm.Draw(renderer,w)
	renderer.Present()

	L:
	for {
		event := sdl.WaitEvent()
		switch event.(type) {
		case *sdl.QuitEvent:
			break L
		}

	}

	sdl.Quit()
}