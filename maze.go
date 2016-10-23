package main

import (
	"go_sdl2_examples/maze"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)


func main() {
	mm:=maze.BuildMaze(5)
	path:=maze.NewPointStack()


	ps:=maze.NewPointSet()
	mm.GetOpenPointSet(0,mm.Len()-1,ps)

	if ps.HasPoint(mm.Len()-1,0) {
		log.Println("The maze has a path to out")
		mm.FindPath(0,mm.Len()-1,mm.Len()-1,0,path)
	} else {
		log.Println("The maze has NOT a path to out")
	}

	w:=10
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
	renderer.SetViewport(&sdl.Rect{10,10,int32(mm.Len()*w+1),int32(mm.Len()*w+1)})

	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	mm.Draw(renderer,w)


	renderer.SetDrawColor(100, 100, 100, 255)

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