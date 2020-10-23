package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
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

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		player.draw(renderer)

		renderer.Present()
	}
}
