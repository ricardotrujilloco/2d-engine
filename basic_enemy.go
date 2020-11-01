package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const basicEnemySize = 105

type BasicEnemy struct {
	texture *sdl.Texture
	x, y    float64
}

func newBasicEnemy(renderer *sdl.Renderer, x, y float64) (basicEnemy BasicEnemy, error error) {
	img, err := sdl.LoadBMP("data/sprites/basic_enemy.bmp")
	if err != nil {
		return BasicEnemy{}, fmt.Errorf("loading basic enemy sprite error: %v", err)
	}
	defer img.Free()

	basicEnemy.texture, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return BasicEnemy{}, fmt.Errorf("creating basic enemy texture error: %v", err)
	}

	basicEnemy.x = x
	basicEnemy.y = y

	return basicEnemy, nil
}

func (basicEnemy *BasicEnemy) draw(renderer *sdl.Renderer) {
	x, y := basicEnemy.getPlayerCoordinatesWithSizeOffset()

	renderer.CopyEx(
		basicEnemy.texture,
		&sdl.Rect{X: 0, Y: 0, W: basicEnemySize, H: basicEnemySize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: basicEnemySize, H: basicEnemySize},
		180,
		&sdl.Point{X: basicEnemySize / 2, Y: basicEnemySize / 2},
		sdl.FLIP_NONE)
}

func (basicEnemy *BasicEnemy) getPlayerCoordinatesWithSizeOffset() (float64, float64) {
	x := basicEnemy.x - basicEnemySize/2.0
	y := basicEnemy.y - basicEnemySize/2.0
	return x, y
}
