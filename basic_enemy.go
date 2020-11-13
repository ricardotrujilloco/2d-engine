package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const basicEnemySize = 105

type enemy struct {
	element
	state ElementState
}

func (elem *enemy) isActive() *bool {
	return &elem.active
}

func (elem *enemy) getPosition() *vector {
	return &elem.position
}

func (elem *enemy) getRotation() *float64 {
	return &elem.rotation
}

func (elem *enemy) getWidth() *float64 {
	return &elem.width
}

func (elem *enemy) update(updateParameters updateParameters) error {
	for _, comp := range elem.logicComponents {
		err := comp.onUpdate(updateParameters)
		if err != nil {
			return err
		}
		switch comp.(type) {
		case *animator:
			animator := comp.(*animator)
			if animator.finished {
				elem.state = Inactive
				elem.active = false
			}
		}
	}
	return nil
}

func (elem *enemy) onCollision(otherElement gameObject) error {
	for _, attr := range elem.attributes {
		switch attr.(type) {
		case *vulnerableToBullets:
			elem.state = Destroying
			for _, comp := range elem.logicComponents {
				switch comp.(type) {
				case *animator:
					animator := comp.(*animator)
					animator.setSequence(Destroying)
				}
			}
		}
	}
	return nil
}

func (elem *enemy) draw() error {
	parameters := drawParameters{
		position: *elem.getPosition(),
		rotation: *elem.getRotation(),
	}
	for _, comp := range elem.uiComponents {
		err := comp.onDraw(parameters)
		if err != nil {
			return err
		}
	}
	return nil
}

func (elem *enemy) getBoundingCircle() boundingCircle {
	return elem.boundingCircle
}

func newBasicEnemy(renderer *sdl.Renderer, position vector) *enemy {
	animator := newAnimator(getEnemySequences(), Idle)
	return &enemy{
		element{
			position: position,
			rotation: 180,
			active:   true,
			logicComponents: []logicComponent{
				newBulletMover(bulletSpeed),
				animator,
			},
			attributes: []attribute{&vulnerableToBullets{}},
			uiComponents: []uiComponent{
				newMultiSpriteRenderer(
					renderer,
					getEnemyUiSequences(renderer),
					animator,
				),
			},
			boundingCircle: boundingCircle{
				center: position,
				radius: 52,
			},
		},
		Idle,
	}
}

func getEnemySequences() map[ElementState]*sequence {
	idleSequence, err := newSequence("data/sprites/basic_enemy/idle", 5, true)
	if err != nil {
		panic(fmt.Errorf("creating idle sequence: %v", err))
	}
	destroySequence, err := newSequence("data/sprites/basic_enemy/destroy", 15, false)
	if err != nil {
		panic(fmt.Errorf("creating destroy sequence: %v", err))
	}
	sequences := map[ElementState]*sequence{
		Idle:       idleSequence,
		Destroying: destroySequence,
	}
	return sequences
}

func getEnemyUiSequences(renderer *sdl.Renderer) map[ElementState]*multiSpriteRendererSequence {
	idleSequenceUi, err := newMultiSpriteRendererSequence("data/sprites/basic_enemy/idle", renderer)
	if err != nil {
		panic(fmt.Errorf("creating idle sequence: %v", err))
	}
	destroySequenceUi, err := newMultiSpriteRendererSequence("data/sprites/basic_enemy/destroy", renderer)
	if err != nil {
		panic(fmt.Errorf("creating destroy sequence: %v", err))
	}
	uiSequences := map[ElementState]*multiSpriteRendererSequence{
		Idle:       idleSequenceUi,
		Destroying: destroySequenceUi,
	}
	return uiSequences
}
