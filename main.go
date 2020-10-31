package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"strconv"
	"time"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println("Initializing SDL error: ", err)
		return
	}
	window, err := sdl.CreateWindow(
		"gaming in Go Episode 2",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth,
		screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Creating window error: ", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(
		window,
		-1,
		sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing renderer error: ", err)
		return
	}
	defer renderer.Destroy()

	player, err := newPlayer(renderer)
	if err != nil {
		fmt.Println("Creating player error: ", err)
		return
	}

	err = ttf.Init()
	if err != nil {
		fmt.Println("TTF init error: ", err)
		return
	}
	font, err := ttf.OpenFont("fonts/cabal.ttf", 48)
	if err != nil {
		fmt.Println("TTF open font error: ", err)
		return
	}
	defer font.Close()

	now := float64(0)
	elapsed := float64(0)

	for {
		now = float64(time.Now().UnixNano()) / 1000000.0

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		player.draw(renderer)
		player.update()

		elapsedMillisString := strconv.FormatFloat(elapsed, 'f', 2, 64)
		drawFps(elapsedMillisString, font, renderer)

		renderer.Present()

		elapsed = float64(time.Now().UnixNano())/1000000.0 - now

		fmt.Printf("it took %v\n", elapsedMillisString)
	}
}

func drawFps(fpsValue string, font *ttf.Font, renderer *sdl.Renderer) bool {
	surface, err := font.RenderUTF8Blended(fpsValue, sdl.Color{R: 255, G: 151, B: 157})
	if err != nil {
		fmt.Println("RenderUTF8Solid error: ", err)
		return true
	}
	textTexture, err := renderer.CreateTextureFromSurface(surface)
	_, _, w, h, err := textTexture.Query()
	if err != nil {
		fmt.Println("RenderUTF8Solid error: ", err)
		return true
	}
	renderer.Copy(
		textTexture,
		&sdl.Rect{X: 0, Y: 0, W: w, H: h},
		&sdl.Rect{X: 20, Y: 20, W: w, H: h})
	return false
}
