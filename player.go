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

func newPlayer(renderer *sdl.Renderer) *element {
	spriteRenderer := newSpriteRenderer(renderer, "data/sprites/player.bmp")
	return &element{
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
	}
}

func (elem *element) onComponentsUpdated() {
	for _, comp := range elem.logicComponents {
		switch comp.(type) {
		case *keyboardMover:
			elem.position.x = comp.(*keyboardMover).position.x
		}
	}
}
