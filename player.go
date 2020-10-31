package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 0.5
	playerSize  = 105
)

type Player struct {
	texture *sdl.Texture
	x, y    float64
}

func newPlayer(renderer *sdl.Renderer) (player Player, error error) {
	img, err := sdl.LoadBMP("sprites/player.bmp")
	if err != nil {
		return Player{}, fmt.Errorf("loading player sprite error: %v", err)
	}
	defer img.Free()

	player.texture, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return Player{}, fmt.Errorf("creating player texture error: %v", err)
	}

	player.x = screenWidth / 2.0
	player.y = screenHeight - playerSize/2.0

	return player, nil
}

func (player *Player) draw(renderer *sdl.Renderer) {
	x, y := player.getPlayerCoordinatesWithSizeOffset()

	renderer.Copy(
		player.texture,
		&sdl.Rect{X: 0, Y: 0, W: playerSize, H: playerSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerSize, H: playerSize})
}

func (player *Player) getPlayerCoordinatesWithSizeOffset() (float64, float64) {
	x := player.x - playerSize/2.0
	y := player.y - playerSize/2.0
	return x, y
}

func (player *Player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		player.x = player.x - playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		player.x = player.x + playerSpeed
	}
}
