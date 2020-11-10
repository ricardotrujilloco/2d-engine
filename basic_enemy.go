package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const basicEnemySize = 105

func newBasicEnemy(renderer *sdl.Renderer, position vector) *element {
	return &element{
		position:        position,
		rotation:        180,
		active:          true,
		logicComponents: []logicComponent{newBulletMover(bulletSpeed)},
		uiComponents:    []uiComponent{newSpriteRenderer(renderer, "data/sprites/basic_enemy.bmp")},
	}
}
