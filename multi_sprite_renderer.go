package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"io/ioutil"
	"path"
)

type multiSpriteRenderer struct {
	renderer    *sdl.Renderer
	sequences   map[ElementState]*multiSpriteRendererSequence
	animator    *animator
	scaleFactor float64
}

type multiSpriteDrawParameters struct {
	position vector
	rotation float64
}

func (parameters *multiSpriteDrawParameters) getPosition() vector {
	return parameters.position
}

func (parameters *multiSpriteDrawParameters) getRotation() float64 {
	return parameters.rotation
}

type multiSpriteRendererSequence struct {
	textures []*sdl.Texture
}

func newMultiSpriteRenderer(
	renderer *sdl.Renderer,
	sequences map[ElementState]*multiSpriteRendererSequence,
	animator *animator, // Must reference the same animator instance uses as a logi cmoponent
	scaleFactor float64,
) *multiSpriteRenderer {
	return &multiSpriteRenderer{
		renderer:    renderer,
		sequences:   sequences,
		animator:    animator,
		scaleFactor: scaleFactor,
	}
}

func (sr *multiSpriteRenderer) onDraw(parameters drawParameters) error {
	frame := sr.animator.sequences[sr.animator.currentSequence].frame
	tex := sr.sequences[sr.animator.currentSequence].textures[frame]

	_, _, width, height, err := tex.Query()
	scaledWidth := int32(float64(width) * sr.scaleFactor)
	scaledHeight := int32(float64(height) * sr.scaleFactor)
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}

	// Converting coordinates to top left of sprite
	scaledX := parameters.getPosition().x - float64(scaledWidth)/2.0
	scaledY := parameters.getPosition().y - float64(scaledHeight)/2.0

	sr.renderer.CopyEx(
		tex,
		&sdl.Rect{X: 0, Y: 0, W: width, H: height},
		&sdl.Rect{X: int32(scaledX), Y: int32(scaledY), W: scaledWidth, H: scaledHeight},
		parameters.getRotation(),
		&sdl.Point{X: scaledWidth / 2, Y: scaledHeight / 2},
		sdl.FLIP_NONE)

	return nil
}

func newMultiSpriteRendererSequence(
	filepath string,
	renderer *sdl.Renderer,
) (*multiSpriteRendererSequence, error) {

	var seq multiSpriteRendererSequence

	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return nil, fmt.Errorf("reading directory %v: %v", filepath, err)
	}

	for _, file := range files {
		filename := path.Join(filepath, file.Name())
		tex := textureFromPNG(renderer, filename)

		seq.textures = append(seq.textures, tex)
	}

	return &seq, nil
}
