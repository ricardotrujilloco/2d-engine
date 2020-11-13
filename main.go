package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/ttf"
	"reflect"
	"strconv"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL: ", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Gaming in Go Episode 2",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window: ", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		fmt.Println("initializing renderer: ", err)
		return
	}
	defer renderer.Destroy()

	elements = append(elements, newPlayer(renderer))

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidth + (basicEnemyWidth / 2.0)
			y := float64(j)*basicEnemyWidth + (basicEnemyWidth / 2.0)

			enemy := newBasicEnemy(renderer, vector{x, y})
			elements = append(elements, enemy)
		}
	}

	initBulletPool(renderer)

	font, errorGettingFont := getFont(err)
	if errorGettingFont {
		return
	}
	defer font.Close()

	sdl.GetNumVideoDisplays()
	mode, err := sdl.GetDisplayMode(0, 0)
	if err != nil {
		panic(fmt.Errorf("display mode error: %v", err))
	}
	desiredFps := float64(mode.RefreshRate)
	mainLoop(renderer, font, desiredFps, err)
}

func mainLoop(renderer *sdl.Renderer, font *ttf.Font, desiredFps float64, err error) {
	now := float64(0)
	fps := 0.0
	desiredFrameTime := 1000 / desiredFps
	timeElapsedSinceLastLoop := float64(16)
	timeElapsedSinceLastFpsDraw := float64(500)

	for {
		if timeElapsedSinceLastLoop >= desiredFrameTime {
			if checkForQuitEvent() {
				return
			}

			now = float64(time.Now().UnixNano()) / 1000000.0

			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.Clear()

			go updateElements(timeElapsedSinceLastLoop)
			go checkCollisions()
			if drawElements(err) {
				return
			}

			drawFps(strconv.FormatFloat(fps, 'f', 0, 64), font, renderer)

			renderer.Present()

			if timeElapsedSinceLastFpsDraw >= 500 {
				fps = 1000 / timeElapsedSinceLastLoop
				timeElapsedSinceLastFpsDraw = 0
			}
		}

		time.Sleep(time.Duration((desiredFrameTime-timeElapsedSinceLastLoop)/2) * time.Millisecond)

		timeElapsedSinceLastLoop = float64(time.Now().UnixNano())/1000000.0 - now
		timeElapsedSinceLastFpsDraw = timeElapsedSinceLastFpsDraw + timeElapsedSinceLastLoop
	}
}

func updateElements(timeElapsedSinceLastLoop float64) {
	var playerInstance *player
	for _, elem := range elements {
		if reflect.TypeOf(elem) == reflect.TypeOf(&player{}) {
			playerInstance = elem.(*player)
		}
	}
	for _, elem := range elements {
		if *elem.isActive() && playerInstance != nil {
			updateParameters := updateParameters{
				position: *playerInstance.getPosition(),
				elapsed:  timeElapsedSinceLastLoop,
				width:    *elem.getWidth(),
			}
			err := elem.update(updateParameters)
			if err != nil {
				fmt.Println("updating element: ", elem)
				return
			}
		}
	}
}

func drawElements(err error) bool {
	for _, elem := range elements {
		if *elem.isActive() {
			err = elem.draw()
			if err != nil {
				fmt.Println("drawing element: ", elem)
				return true
			}
		}
	}
	return false
}

func checkForQuitEvent() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return true
		}
	}
	return false
}

func getFont(err error) (*ttf.Font, bool) {
	err = ttf.Init()
	if err != nil {
		fmt.Println("TTF init error: ", err)
		return nil, true
	}
	font, err := ttf.OpenFont("data/fonts/cabal.ttf", 48)
	if err != nil {
		fmt.Println("TTF open font error: ", err)
		return nil, true
	}
	return font, false
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
