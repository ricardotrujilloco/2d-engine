package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

const (
	gravity       = 0.4
	yAcceleration = 0.9
	xAcceleration = 0.5
)

type keyboardMover struct {
	position vector
	velocity vector
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
	} else if keys[sdl.SCANCODE_X] == 1 {
		if mover.velocity.y == 0 {
			mover.velocity.y = yAcceleration
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
			mover.shoot(parameters)
			mover.lastShot = time.Now()
		}
	}
	return nil
}

func (mover *keyboardShooter) shoot(parameters updateParameters) {
	if bul, ok := bulletFromPool(); ok {
		bul.update(parameters)
	}
}

type jumpMover struct {
	position vector
	velocity vector
	state    ElementState
}

func newJumpMover() *jumpMover {
	return &jumpMover{
		position: vector{
			x: screenWidth / 2.0,
			y: screenHeight - playerSize/2.0,
		},
		velocity: vector{
			x: 0,
			y: 0,
		},
	}
}

func (jumper *jumpMover) setState(state ElementState) {
	jumper.state = state
}

func (jumper *jumpMover) onUpdate(parameters updateParameters) error {
	if jumper.state == Jumping {
		jumper.position.y -= jumper.velocity.y * parameters.elapsed
		jumper.velocity.y -= gravity / parameters.elapsed
	}
	return nil
}
