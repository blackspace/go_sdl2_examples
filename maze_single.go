package main

import (
	"go_sdl2_examples/maze"
	"github.com/veandco/go-sdl2/sdl"
)


func main() {
	sdl.Init(sdl.INIT_EVERYTHING)

	const W=1000
	const H=800

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

	w:=10
	n:=30
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	mm:=maze.BuildMaze(n,0)
	path:=maze.NewPointStack()
	mm.FindPath(0,n-1,n-1,0,path)

	renderer.SetDrawColor(200,200,200,255)
	mm.Draw(renderer,w)
	renderer.SetDrawColor(150, 10, 10, 255)

	for i:=0;i<path.Count()-1;i++ {
		x0,y0:=path.Index(i)
		x1,y1:=path.Index(i+1)
		renderer.DrawLine(x0*w+w/2,y0*w+w/2,x1*w+w/2,y1*w+w/2)
	}

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