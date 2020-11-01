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
	desiredFps   = 60.0
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println("Initializing SDL error: ", err)
		return
	}

	window, errorGettingWindow := getWindow(err)
	if errorGettingWindow {
		return
	}
	defer window.Destroy()

	renderer, errorGettingRenderer := getRenderer(err, window)
	if errorGettingRenderer {
		return
	}
	defer renderer.Destroy()

	player, errorGettingPlayer := getPlayer(err, renderer)
	if errorGettingPlayer {
		return
	}

	enemies, errorGettingEnemies := getEnemies(renderer)
	if errorGettingEnemies {
		return
	}

	font, errorGettingFont := getFont(err)
	if errorGettingFont {
		return
	}
	defer font.Close()

	mainLoop(renderer, player, enemies, font)
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

func getRenderer(err error, window *sdl.Window) (*sdl.Renderer, bool) {
	renderer, err := sdl.CreateRenderer(
		window,
		-1,
		sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		fmt.Println("Initializing renderer error: ", err)
		return nil, true
	}
	return renderer, false
}

func getWindow(err error) (*sdl.Window, bool) {
	window, err := sdl.CreateWindow(
		"gaming in Go Episode 2",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth,
		screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Creating window error: ", err)
		return nil, true
	}
	return window, false
}

func getPlayer(err error, renderer *sdl.Renderer) (Player, bool) {
	player, err := newPlayer(renderer, screenWidth, screenHeight)
	if err != nil {
		fmt.Println("Creating player error: ", err)
		return Player{}, true
	}
	return player, false
}

func getEnemies(renderer *sdl.Renderer) ([]BasicEnemy, bool) {
	var enemies []BasicEnemy
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidth + (basicEnemySize / 2.0)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2.0)

			enemy, err := newBasicEnemy(renderer, x, y)
			if err != nil {
				fmt.Println("Creating basic enemy error: ", err)
				return []BasicEnemy{}, true
			}

			enemies = append(enemies, enemy)
		}
	}
	return enemies, false
}

func mainLoop(renderer *sdl.Renderer, player Player, enemies []BasicEnemy, font *ttf.Font) {
	now := float64(0)
	timeElapsedSinceLastLoop := float64(16)
	timeElapsedSinceLastFpsDraw := float64(500)
	fps := 0.0

	desiredFrameTime := 1000 / desiredFps

	for {
		if timeElapsedSinceLastLoop >= desiredFrameTime {
			if checkForQuitEvent() {
				return
			}

			now = float64(time.Now().UnixNano()) / 1000000.0
			player.update(timeElapsedSinceLastLoop)

			if timeElapsedSinceLastFpsDraw >= 500 {
				fps = 1000 / timeElapsedSinceLastLoop
				timeElapsedSinceLastFpsDraw = 0
			}

			drawFrame(renderer, player, enemies, fps, font)
		}

		time.Sleep(time.Duration(desiredFrameTime-timeElapsedSinceLastLoop) * time.Millisecond)

		timeElapsedSinceLastLoop = float64(time.Now().UnixNano())/1000000.0 - now
		timeElapsedSinceLastFpsDraw = timeElapsedSinceLastFpsDraw + timeElapsedSinceLastLoop
	}
}

func drawFrame(renderer *sdl.Renderer, player Player, enemies []BasicEnemy, fps float64, font *ttf.Font) {
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()

	player.draw(renderer)

	for _, enemy := range enemies {
		enemy.draw(renderer)
	}

	drawFps(strconv.FormatFloat(fps, 'f', 0, 64), font, renderer)

	renderer.Present()
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
