package main

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type circleRenderer struct {
	renderer *sdl.Renderer
	circles  []*boundingCircle
}

type circleDrawParameters struct {
	drawParameters drawParameters
	radius         int32
}

func (parameters *circleDrawParameters) getPosition() vector {
	return parameters.drawParameters.getPosition()
}

func (parameters *circleDrawParameters) getRotation() float64 {
	return parameters.drawParameters.getRotation()
}

func newCircleRenderer(
	renderer *sdl.Renderer,
	circles []*boundingCircle,
) *circleRenderer {
	return &circleRenderer{
		renderer: renderer,
		circles:  circles,
	}
}

func (sr *circleRenderer) onDraw(parameters drawParameters) error {
	radius := parameters.(*circleDrawParameters).radius
	gfx.CircleColor(sr.renderer, int32(parameters.getPosition().x), int32(parameters.getPosition().y), radius, sdl.Color{0, 0, 255, 255})
	return nil
}
