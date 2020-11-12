package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

const (
	playerSpeed        = 0.5
	playerSize         = 105
	playerShotCoolDown = time.Millisecond * 250
)

type element struct {
	position        vector
	width           float64
	rotation        float64
	active          bool
	logicComponents []logicComponent
	uiComponents    []uiComponent
}

type player struct {
	element
}

func (elem *player) isActive() *bool {
	return &elem.active
}

func (elem *player) getPosition() *vector {
	return &elem.position
}

func (elem *player) getRotation() *float64 {
	return &elem.rotation
}

func (elem *player) getWidth() *float64 {
	return &elem.width
}

func (elem *player) update(updateParameters updateParameters) error {
	for _, comp := range elem.logicComponents {
		err := comp.onUpdate(updateParameters)
		if err != nil {
			return err
		}
	}
	for _, comp := range elem.logicComponents {
		switch comp.(type) {
		case *keyboardMover:
			elem.position.x = comp.(*keyboardMover).position.x
		}
	}
	return nil
}

func (elem *player) draw(parameters drawParameters) error {
	for _, comp := range elem.uiComponents {
		err := comp.onDraw(parameters)
		if err != nil {
			return err
		}
	}

	return nil
}

func newPlayer(renderer *sdl.Renderer) *player {
	spriteRenderer := newSpriteRenderer(renderer, "data/sprites/player.bmp")
	return &player{
		element{
			position: vector{
				x: screenWidth / 2.0,
				y: screenHeight - playerSize/2.0,
			},
			width:  spriteRenderer.width,
			active: true,
			logicComponents: []logicComponent{
				newKeyboardMover(playerSpeed),
				newKeyboardShooter(playerShotCoolDown),
			},
			uiComponents: []uiComponent{spriteRenderer},
		},
	}
}
