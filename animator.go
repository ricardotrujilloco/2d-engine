package main

import "github.com/veandco/go-sdl2/sdl"

type Animator struct {
	container *Element
	sequences map[string]*Sequence
	current   string
}

func newAnimator(container *Element, sequences map[string]*Sequence, defaultSequence string) *Animator {
	var animator Animator

	animator.container = container
	animator.sequences = sequences
	animator.current = defaultSequence

	return &animator
}

/*func (animator *Animator) onDraw(renderer * sdl.Renderer) error {

}*/

type Sequence struct {
	textures   []*sdl.Texture
	frame      int
	sampleRate float64
}
