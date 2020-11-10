package main

import (
	"math"
)

type bulletMover struct {
	speed float64
}

type updateParameters struct {
	position vector
	rotation float64
	width    float64
	elapsed  float64
}

func newBulletMover(speed float64) *bulletMover {
	return &bulletMover{
		speed: speed,
	}
}

func (mover *bulletMover) onUpdate(parameters updateParameters) error {

	parameters.position.x += bulletSpeed * math.Cos(parameters.rotation) * parameters.elapsed
	parameters.position.y += bulletSpeed * math.Sin(parameters.rotation) * parameters.elapsed

	if parameters.position.x > screenWidth || parameters.position.x < 0 ||
		parameters.position.y > screenHeight || parameters.position.y < 0 {
		return nil
	}

	return nil
}
