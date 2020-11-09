package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"time"
)

type keyboardMover struct {
	container      *element
	speed          float64
	spriteRenderer *spriteRenderer
}

func newKeyboardMover(container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container:      container,
		speed:          speed,
		spriteRenderer: container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) onUpdate(elapsed float64) error {
	keys := sdl.GetKeyboardState()

	gameObject := mover.container

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if gameObject.position.x-(mover.spriteRenderer.width/2.0) > 0 {
			gameObject.position.x -= mover.speed * elapsed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if gameObject.position.x+(mover.spriteRenderer.width/2.0) < screenWidth {
			gameObject.position.x += mover.speed * elapsed
		}
	}

	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

type keyboardShooter struct {
	container *element
	cooldown  time.Duration
	lastShot  time.Time
}

func newKeyboardShooter(container *element, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		cooldown:  cooldown,
	}
}

func (mover *keyboardShooter) onUpdate(elapsed float64) error {
	keys := sdl.GetKeyboardState()

	pos := mover.container.position

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(mover.lastShot) >= mover.cooldown {
			mover.shoot(pos.x+25, pos.y-20)
			mover.shoot(pos.x-25, pos.y-20)

			mover.lastShot = time.Now()
		}
	}

	return nil
}

func (mover *keyboardShooter) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardShooter) shoot(x, y float64) {
	if bul, ok := bulletFromPool(); ok {
		bul.active = true
		bul.position.x = x
		bul.position.y = y
		bul.rotation = 270 * (math.Pi / 180)
	}
}
