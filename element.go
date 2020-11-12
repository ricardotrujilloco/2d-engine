package main

type vector struct {
	x, y float64
}

type gameObject interface {
	isActive() *bool
	getPosition() *vector
	getRotation() *float64
	getWidth() *float64
	update(updateParameters updateParameters) error
	draw(parameters drawParameters) error
}

type uiComponent interface {
	onDraw(parameters drawParameters) error
}

type logicComponent interface {
	onUpdate(parameters updateParameters) error
}

var elements []gameObject
