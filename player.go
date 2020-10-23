package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	texture *sdl.Texture
}

func newPlayer(renderer *sdl.Renderer) (player Player, error error) {
	img, err := sdl.LoadBMP("sprites/player.bmp")
	if err != nil {
		return Player{}, fmt.Errorf("Loading player sprite  error: %v", err)
	}
	defer img.Free()

	player.texture, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return Player{}, fmt.Errorf("Creating player texture error: %v", err)
	}

	return player, nil
}

func (player *Player) draw(renderer *sdl.Renderer) {
	renderer.Copy(
		player.texture,
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
		&sdl.Rect{X: 40, Y: 20, W: 210, H: 210})
}
