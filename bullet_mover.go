package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type bulletMover struct {
	container *element
	speed     float64
}

func newBulletMover(container *element, speed float64) *bulletMover {
	return &bulletMover{
		container: container,
		speed:     speed,
	}
}

func (mover *bulletMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *bulletMover) onUpdate(elapsed float64) error {
	cont := mover.container

	cont.position.x += bulletSpeed * math.Cos(cont.rotation) * elapsed
	cont.position.y += bulletSpeed * math.Sin(cont.rotation) * elapsed

	if cont.position.x > screenWidth || cont.position.x < 0 ||
		cont.position.y > screenHeight || cont.position.y < 0 {
		cont.active = false
	}

	return nil
}
