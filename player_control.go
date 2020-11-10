package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"time"
)

type keyboardMover struct {
	position vector
	speed    float64
}

func newKeyboardMover(speed float64) *keyboardMover {
	return &keyboardMover{
		position: vector{
			x: screenWidth / 2.0,
			y: screenHeight - playerSize/2.0,
		},
		speed: speed,
	}
}

func (mover *keyboardMover) onUpdate(parameters updateParameters) error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if parameters.position.x-(parameters.width/2.0) >= 0 {
			mover.position.x -= mover.speed * parameters.elapsed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if parameters.position.x+(parameters.width/2.0) <= screenWidth {
			mover.position.x += mover.speed * parameters.elapsed
		}
	}
	return nil
}

type keyboardShooter struct {
	coolDown time.Duration
	lastShot time.Time
}

func newKeyboardShooter(coolDown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		coolDown: coolDown,
	}
}

func (mover *keyboardShooter) onUpdate(parameters updateParameters) error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(mover.lastShot) >= mover.coolDown {
			mover.shoot(parameters.position.x+25, parameters.position.y-20)
			mover.shoot(parameters.position.x-25, parameters.position.y-20)

			mover.lastShot = time.Now()
		}
	}
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
