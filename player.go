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
	maxX    float64
	maxY    float64
}

func newPlayer(renderer *sdl.Renderer, maxX float64, maxY float64) (player Player, error error) {
	img, err := sdl.LoadBMP("data/sprites/player.bmp")
	if err != nil {
		return Player{}, fmt.Errorf("loading player sprite error: %v", err)
	}
	defer img.Free()

	player.texture, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return Player{}, fmt.Errorf("creating player texture error: %v", err)
	}

	player.x = maxX / 2.0
	player.y = maxY - playerSize/2.0

	player.maxX = maxX
	player.maxY = maxY

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

func (player *Player) update(elapsed float64) {
	keys := sdl.GetKeyboardState()

	player.moveInX(elapsed, keys)
	player.moveInY(elapsed, keys)
}

func (player *Player) moveInX(elapsed float64, keys []uint8) {
	if keys[sdl.SCANCODE_LEFT] == 1 {
		newX := player.x - playerSpeed*elapsed
		if newX >= playerSize/2.0 {
			player.x = newX
		} else {
			player.x = playerSize / 2.0
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		newX := player.x + playerSpeed*elapsed
		if newX <= player.maxX-playerSize/2.0 {
			player.x = newX
		} else {
			player.x = player.maxX - playerSize/2.0
		}
	}
}

func (player *Player) moveInY(elapsed float64, keys []uint8) {
	if keys[sdl.SCANCODE_UP] == 1 {
		newY := player.y - playerSpeed*elapsed
		if newY >= playerSize/2.0 {
			player.y = newY
		} else {
			player.y = playerSize / 2.0
		}
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		newY := player.y + playerSpeed*elapsed
		if newY <= player.maxY-playerSize/2.0 {
			player.y = newY
		} else {
			player.y = player.maxY - playerSize/2.0
		}
	}
}
