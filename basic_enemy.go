package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
)

const (
	basicEnemyWidth  = 105
	basicEnemyHeight = 72
	explosionSpeed   = 0.1
)

type enemy struct {
	element
	state ElementState
}

func (elem *enemy) isActive() bool {
	return elem.active
}

func (elem *enemy) getPosition() vector {
	return elem.position
}

func (elem *enemy) getRotation() float64 {
	return elem.rotation
}

func (elem *enemy) getWidth() float64 {
	return elem.width
}

func (elem *enemy) update(updateParameters updateParameters) error {
	var err error = nil
	for _, comp := range elem.logicComponents {
		switch comp.(type) {
		case *animator:
			err = comp.onUpdate(updateParameters)
			animator := comp.(*animator)
			if animator.finished {
				elem.state = Inactive
				elem.active = false
			}
		case *boundingCircleScaler:
			scaler := comp.(*boundingCircleScaler)
			if elem.state == Destroying && !scaler.isMaxRadiusReached {
				err = comp.onUpdate(updateParameters)
			}
		}
	}
	return err
}

func (elem *enemy) onCollision(otherElement gameObject) error {
	switch otherElement.(type) {
	case *bullet:
		elem.onBulletCollision()
	}
	return nil
}

func (elem *enemy) draw() error {
	parameters := multiSpriteDrawParameters{
		position: elem.getPosition(),
		rotation: elem.getRotation(),
	}
	circleParameters := circleDrawParameters{
		drawParameters: &parameters,
		radius:         int32(elem.getBoundingCircle().radius),
	}
	for _, comp := range elem.uiComponents {
		var err error = nil
		if reflect.TypeOf(comp) == reflect.TypeOf(&circleRenderer{}) {
			err = comp.onDraw(&circleParameters)
		} else {
			err = comp.onDraw(&parameters)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (elem *enemy) getBoundingCircle() *boundingCircle {
	return elem.boundingCircle
}

func (elem *enemy) onBulletCollision() {
	isVulnerableToBullets := false
	for _, attr := range elem.attributes {
		switch attr.(type) {
		case *vulnerableToBullets:
			isVulnerableToBullets = true
		}
	}
	if isVulnerableToBullets {
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

func newBasicEnemy(renderer *sdl.Renderer, position vector) enemy {
	destroyingSampleRate := 15.0
	basicEnemyRadiusScaleFactor := 0.25
	basicEnemyInitialRadius := (810 / 4) * basicEnemyRadiusScaleFactor // From sprite dimensions
	basicEnemyFinalRadius := (810 / 2) * basicEnemyRadiusScaleFactor   // From sprite dimensions
	animator := newAnimator(getEnemySequences(destroyingSampleRate), Idle)
	circle := &boundingCircle{center: position, radius: basicEnemyInitialRadius}
	boundingCircles := []*boundingCircle{circle}
	boundingCircleScaler := newBoundingCircleScaler(boundingCircles, basicEnemyFinalRadius)
	return enemy{
		element{
			position: position,
			rotation: 180,
			active:   true,
			logicComponents: []logicComponent{
				animator,
				boundingCircleScaler,
			},
			attributes: []attribute{&vulnerableToBullets{}},
			uiComponents: []uiComponent{
				newMultiSpriteRenderer(
					renderer,
					getEnemyUiSequences(renderer),
					animator,
					basicEnemyRadiusScaleFactor,
				),
				newCircleRenderer(
					renderer,
					boundingCircles,
				),
			},
			boundingCircle: circle,
		},
		Idle,
	}
}

func getEnemySequences(
	destroyingSampleRate float64,
) map[ElementState]*sequence {
	idleSequence, err := newSequence("data/sprites/bomb/idle", 10, true, false)
	if err != nil {
		panic(fmt.Errorf("creating idle sequence: %v", err))
	}
	destroySequence, err := newSequence("data/sprites/bomb/destroy", destroyingSampleRate, false, true)
	if err != nil {
		panic(fmt.Errorf("creating onBulletCollision sequence: %v", err))
	}
	sequences := map[ElementState]*sequence{
		Idle:       idleSequence,
		Destroying: destroySequence,
	}
	return sequences
}

func getEnemyUiSequences(renderer *sdl.Renderer) map[ElementState]*multiSpriteRendererSequence {
	idleSequenceUi, err := newMultiSpriteRendererSequence("data/sprites/bomb/idle", renderer)
	if err != nil {
		panic(fmt.Errorf("creating idle sequence: %v", err))
	}
	destroySequenceUi, err := newMultiSpriteRendererSequence("data/sprites/bomb/destroy", renderer)
	if err != nil {
		panic(fmt.Errorf("creating onBulletCollision sequence: %v", err))
	}
	uiSequences := map[ElementState]*multiSpriteRendererSequence{
		Idle:       idleSequenceUi,
		Destroying: destroySequenceUi,
	}
	return uiSequences
}
