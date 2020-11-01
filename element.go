package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer)
	onCollision(other *Element)
}

type Element struct {
	position   Vector
	rotation   float64
	active     bool
	tag        string
	collisions []Circle
	components []Component
}

type Vector struct {
	x float64
	y float64
}

type Circle struct {
	center Vector
	radius float64
}
