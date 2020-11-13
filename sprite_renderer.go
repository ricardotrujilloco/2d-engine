package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	renderer      *sdl.Renderer
	tex           *sdl.Texture
	width, height float64
}

type drawParameters struct {
	position vector
	rotation float64
}

func newSpriteRenderer(renderer *sdl.Renderer, filename string) *spriteRenderer {
	tex := textureFromBMP(renderer, filename)

	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}

	return &spriteRenderer{
		renderer: renderer,
		tex:      textureFromBMP(renderer, filename),
		width:    float64(width),
		height:   float64(height),
	}
}

func (sr *spriteRenderer) onDraw(parameters drawParameters) error {
	// Converting coordinates to top left of sprite
	x := parameters.position.x - sr.width/2.0
	y := parameters.position.y - sr.height/2.0

	sr.renderer.CopyEx(
		sr.tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(sr.width), H: int32(sr.height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(sr.width), H: int32(sr.height)},
		parameters.rotation,
		&sdl.Point{X: int32(sr.width) / 2, Y: int32(sr.height) / 2},
		sdl.FLIP_NONE)

	return nil
}

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v: %v", filename, err))
	}
	return tex
}
