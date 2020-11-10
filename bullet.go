package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 0.3
)

func newBullet(renderer *sdl.Renderer) *element {
	return &element{
		active:          false,
		logicComponents: []logicComponent{newBulletMover(bulletSpeed)},
		uiComponents:    []uiComponent{newSpriteRenderer(renderer, "data/sprites/player_bullet.bmp")},
	}
}

var bulletPool []*element

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul := newBullet(renderer)
		elements = append(elements, bul)
		bulletPool = append(bulletPool, bul)
	}
}

func bulletFromPool() (*element, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}

	return nil, false
}
