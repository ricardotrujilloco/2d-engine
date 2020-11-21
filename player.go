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

type player struct {
	element
	state ElementState
}

func (elem *player) isActive() bool {
	return elem.state == Active ||
		elem.state == Jumping ||
		elem.state == Destroying
}

func (elem *player) getPosition() vector {
	return elem.position
}

func (elem *player) getRotation() float64 {
	return elem.rotation
}

func (elem *player) getWidth() float64 {
	return elem.width
}

func (elem *player) update(updateParameters updateParameters) error {
	for _, comp := range elem.logicComponents {
		err := comp.onUpdate(updateParameters)
		if err != nil {
			return err
		}
	}
	elem.onKeyboardMoverUpdated()
	return nil
}

func (elem *player) onCollision(otherElement gameObject) error {
	return nil
}

func (elem *player) draw() error {
	parameters := &spriteDrawParameters{
		position: elem.getPosition(),
		rotation: elem.getRotation(),
	}
	for _, comp := range elem.uiComponents {
		err := comp.onDraw(parameters)
		if err != nil {
			return err
		}
	}
	return nil
}

func (elem *player) getBoundingCircle() *boundingCircle {
	return elem.boundingCircle
}

func (elem *player) onKeyboardMoverUpdated() {
	if component, ok := elem.logicComponents[KeyboardMover]; ok {
		keyboardMover := component.(*keyboardMover)
		elem.position.x = keyboardMover.position.x
		if keyboardMover.velocity.y > 0 {
			elem.onJumpMoverUpdated(keyboardMover.velocity.y)
		}
	}
}

func (elem *player) onJumpMoverUpdated(yVelocity float64) {
	if component, ok := elem.logicComponents[JumpMover]; ok {
		jumpMover := component.(*jumpMover)
		if jumpMover.velocity.y == 0 {
			jumpMover.velocity.y = yVelocity
			elem.state = Jumping
			jumpMover.state = Jumping
		}
		elem.position.y = jumpMover.position.y
	}
}

func newPlayer(renderer *sdl.Renderer) player {
	spriteRenderer := newSpriteRenderer(renderer, "data/sprites/player.bmp")
	position := vector{
		x: screenWidth / 2.0,
		y: screenHeight - playerSize/2.0,
	}
	return player{
		state: Active,
		element: element{
			position: position,
			width:    spriteRenderer.width,
			logicComponents: map[LogicComponentType]logicComponent{
				KeyboardMover:   newKeyboardMover(playerSpeed),
				KeyboardShooter: newKeyboardShooter(playerShotCoolDown),
				JumpMover:       newJumpMover(),
			},
			uiComponents: []uiComponent{spriteRenderer},
			boundingCircle: &boundingCircle{
				center: position,
				radius: 8,
			},
		},
	}
}
