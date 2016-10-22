package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"fmt"
	"log"
	"reflect"
)


func run() int {
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
		return -1
	}
	defer renderer.Destroy()

	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	renderer.Present()

	L:
	for {
		event := sdl.WaitEvent()
		log.Printf("%v",reflect.ValueOf(event).Type())
		if event==nil {
			break
		}
		switch t:=event.(type) {
		case *sdl.QuitEvent:
			break L
		case *sdl.KeyUpEvent:
			fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
				t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
		}

	}


	sdl.Quit()

	return 0


}

func main() {
	os.Exit(run())
}



