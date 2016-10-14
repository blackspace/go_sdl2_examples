package test

import (
	"testing"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

func BenchmarkSurfaceUpdate(b *testing.B) {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		1680, 1050, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	for i:=0;i<b.N;i++ {
		rect := sdl.Rect{0, 0, 200, 200}
		surface.FillRect(&rect, 0xffff0000)
		window.UpdateSurface()
	}

	sdl.Quit()
}

func BenchmarkRenderFullScreen(b *testing.B) {
	sdl.Init(sdl.INIT_EVERYTHING)
	defer  sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		1680, 1050, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
		return
	}
	defer renderer.Destroy()

	rect := sdl.Rect{0, 0, 1680, 1050}

	for i:=0;i<b.N;i++ {
		renderer.SetDrawColor(255, 0, 255, 255)
		renderer.FillRect(&rect)
		renderer.Present()
	}
}

func Benchmark16801050Animation(b *testing.B) {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		1680, 1050, sdl.WINDOW_SHOWN)
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


	rect := sdl.Rect{0, 0, 1680, 1050}
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.FillRect(&rect)

	for i:=0;i<b.N;i++ {
		if i>0 {
			rect = sdl.Rect{int32(i-1), 100, 100, 100}
			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.FillRect(&rect)
		}

		rect = sdl.Rect{int32(i), 100, 100, 100}
		renderer.SetDrawColor(0, 255, 255, 255)
		renderer.FillRect(&rect)
		renderer.Present()
	}
}

func Benchmark800600Animation(b *testing.B) {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
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

	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()

	for i:=0;i<b.N;i++ {
		if i>0 {
			rect := sdl.Rect{int32(i-1), 100, 100, 100}
			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.FillRect(&rect)
		}

		rect := sdl.Rect{int32(i), 100, 100, 100}
		renderer.SetDrawColor(0, 255, 255, 255)
		renderer.FillRect(&rect)
		renderer.Present()
	}
}

func Benchmark16801050ImageAnimation(b *testing.B) {
	sdl.Init(sdl.INIT_EVERYTHING)

	const W=1680
	const H=1050

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		W, H, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
		return
	}
	defer renderer.Destroy()

	image,_:=img.Load("sdl2.jpeg")

	defer image.Free()


	texture,_:=renderer.CreateTextureFromSurface(image)
	defer texture.Destroy()

	src := image.ClipRect

	for i:=int32(0);i<int32(b.N);i++ {
		renderer.Clear()
		dst := sdl.Rect{100, i, image.ClipRect.W, image.ClipRect.H}
		renderer.Copy(texture, &src, &dst)
		renderer.Present()
	}

}

func Benchmark800600ImageAnimation(b *testing.B) {
	sdl.Init(sdl.INIT_EVERYTHING)

	const W=800
	const H=600

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		W, H, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
		return
	}
	defer renderer.Destroy()

	image,_:=img.Load("sdl2.jpeg")

	defer image.Free()


	texture,_:=renderer.CreateTextureFromSurface(image)
	defer texture.Destroy()

	src := image.ClipRect

	for i:=int32(0);i<int32(b.N);i++ {
		renderer.Clear()
		dst := sdl.Rect{100, i, image.ClipRect.W, image.ClipRect.H}
		renderer.Copy(texture, &src, &dst)
		renderer.Present()
	}

}


