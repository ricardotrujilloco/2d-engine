package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type BasicEnemy struct {
	texture *sdl.Texture
	x, y    float64
}

func newBasicEnemy(renderer *sdl.Renderer, x, y float64) (basicEnemy BasicEnemy, error error) {
	img, err := sdl.LoadBMP("sprites/basic_enemy.bmp")
	if err != nil {
		return BasicEnemy{}, fmt.Errorf("loading basic enemy sprite error: %v", err)
	}
	defer img.Free()

	basicEnemy.texture, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return BasicEnemy{}, fmt.Errorf("creating basic enemy texture error: %v", err)
	}

	basicEnemy.x = screenWidth / 2.0
	basicEnemy.y = screenHeight - playerSize/2.0

	return basicEnemy, nil
}

func (basicEnemy *BasicEnemy) draw(renderer *sdl.Renderer) {
	x, y := basicEnemy.getPlayerCoordinatesWithSizeOffset()

	renderer.Copy(
		basicEnemy.texture,
		&sdl.Rect{X: 0, Y: 0, W: playerSize, H: playerSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerSize, H: playerSize})
}

func (basicEnemy *BasicEnemy) getPlayerCoordinatesWithSizeOffset() (float64, float64) {
	x := basicEnemy.x - playerSize/2.0
	y := basicEnemy.y - playerSize/2.0
	return x, y
}
