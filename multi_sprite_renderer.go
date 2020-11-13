package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"io/ioutil"
	"path"
)

type multiSpriteRenderer struct {
	renderer  *sdl.Renderer
	sequences map[ElementState]*multiSpriteRendererSequence
	animator  *animator
}

type multiSpriteRendererSequence struct {
	textures []*sdl.Texture
}

func newMultiSpriteRenderer(renderer *sdl.Renderer, sequences map[ElementState]*multiSpriteRendererSequence, animator *animator) *multiSpriteRenderer {
	return &multiSpriteRenderer{
		renderer:  renderer,
		sequences: sequences,
		animator:  animator,
	}
}

func (sr *multiSpriteRenderer) onDraw(parameters drawParameters) error {
	frame := sr.animator.sequences[sr.animator.currentSequence].frame
	tex := sr.sequences[sr.animator.currentSequence].textures[frame]

	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}

	// Converting coordinates to top left of sprite
	x := parameters.position.x - float64(width)/2.0
	y := parameters.position.y - float64(height)/2.0

	sr.renderer.CopyEx(
		tex,
		&sdl.Rect{X: 0, Y: 0, W: width, H: height},
		&sdl.Rect{X: int32(x), Y: int32(y), W: width, H: height},
		parameters.rotation,
		&sdl.Point{X: width / 2, Y: height / 2},
		sdl.FLIP_NONE)

	return nil
}

func newMultiSpriteRendererSequence(
	filepath string,
	renderer *sdl.Renderer) (*multiSpriteRendererSequence, error) {

	var seq multiSpriteRendererSequence

	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return nil, fmt.Errorf("reading directory %v: %v", filepath, err)
	}

	for _, file := range files {
		filename := path.Join(filepath, file.Name())
		tex := textureFromBMP(renderer, filename)

		seq.textures = append(seq.textures, tex)
	}

	return &seq, nil
}
