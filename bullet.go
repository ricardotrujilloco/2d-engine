package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 0.6
)

type bullet struct {
	element
}

func (elem *bullet) isActive() bool {
	return elem.active
}

func (elem *bullet) getPosition() vector {
	return elem.position
}

func (elem *bullet) getRotation() float64 {
	return elem.rotation
}

func (elem *bullet) getWidth() float64 {
	return elem.width
}

func (elem *bullet) update(updateParameters updateParameters) error {
	for _, comp := range elem.logicComponents {
		err := comp.onUpdate(updateParameters)
		if err != nil {
			return err
		}
	}
	for _, comp := range elem.logicComponents {
		switch comp.(type) {
		case *bulletMover:
			position := comp.(*bulletMover).position
			elem.position.x = position.x
			elem.position.y = position.y
			elem.boundingCircle.center = position
			elem.active = comp.(*bulletMover).active
		}
	}
	return nil
}

func (elem *bullet) onCollision(otherElement gameObject) error {
	canCollide := false
	switch otherElement.(type) {
	case *enemy:
		canCollide = true
	}
	for _, comp := range elem.attributes {
		switch comp.(type) {
		case *vulnerableToBullets:
			if canCollide {
				elem.reset()
			}
		}
	}
	return nil
}

func (elem *bullet) draw() error {
	parameters := &spriteDrawParameters{
		position: elem.getPosition(),
		rotation: elem.getRotation(),
	}
	for _, comp := range elem.uiComponents {
		err := comp.onDraw(parameters)
		if err != nil {
			return err
		}
	}
	return nil
}

func (elem *bullet) getBoundingCircle() *boundingCircle {
	return elem.boundingCircle
}

func (elem *bullet) reset() {
	elem.active = false
	for _, comp := range elem.logicComponents {
		switch comp.(type) {
		case *bulletMover:
			comp.(*bulletMover).position.x = 0
			comp.(*bulletMover).position.y = 0
		}
	}
}

func newBullet(renderer *sdl.Renderer) bullet {
	return bullet{
		element{
			active:          false,
			logicComponents: []logicComponent{newBulletMover(bulletSpeed)},
			attributes:      []attribute{&vulnerableToBullets{}},
			uiComponents:    []uiComponent{newSpriteRenderer(renderer, "data/sprites/player_bullet.bmp")},
			boundingCircle: &boundingCircle{
				center: vector{
					x: 0,
					y: 0,
				},
				radius: 16,
			},
		},
	}
}

var bulletPool []gameObject

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul := newBullet(renderer)
		gameObjects = append(gameObjects, &bul)
		bulletPool = append(bulletPool, &bul)
	}
}

func bulletFromPool() (gameObject, bool) {
	for _, bul := range bulletPool {
		if !bul.isActive() {
			return bul, true
		}
	}
	return nil, false
}
