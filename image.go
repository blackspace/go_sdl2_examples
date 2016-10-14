package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)

	const W=1680
	const H=1050

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		W, H, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
		return
	}
	defer renderer.Destroy()

	image,err:=img.Load("sdl2.jpeg")
	if err!=nil {
		panic(err)
	}
	defer image.Free()


	texture,_:=renderer.CreateTextureFromSurface(image)
	defer texture.Destroy()

	src := image.ClipRect


	for i:=int32(0);i<100;i++ {
		renderer.Clear()
		dst := sdl.Rect{100, i, image.ClipRect.W, image.ClipRect.H}
		renderer.Copy(texture, &src, &dst)
		renderer.Present()
	}

	sdl.Delay(2000)

}
