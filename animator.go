package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

type animator struct {
	sequences       map[ElementState]*sequence
	currentSequence ElementState
	lastFrameChange time.Time
	finished        bool
}

type sequence struct {
	textures     []string
	frame        int
	sampleRate   float64
	loop         bool
	finished     bool
	lastSequence bool
}

func (an *animator) onUpdate(parameters updateParameters) error {
	sequence := an.sequences[an.currentSequence]
	frameInterval := float64(time.Second) / sequence.sampleRate

	if time.Since(an.lastFrameChange) >= time.Duration(frameInterval) {
		sequence.nextFrame()
		if sequence.finished && sequence.lastSequence {
			an.finished = true
		}
		an.lastFrameChange = time.Now()
	}
	return nil
}

func (an *animator) setSequence(sequence ElementState) {
	an.currentSequence = sequence
	an.lastFrameChange = time.Now()
}

func newSequence(
	filepath string,
	sampleRate float64,
	loop bool,
	lastSequence bool,
) (*sequence, error) {

	var seq sequence

	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return nil, fmt.Errorf("reading directory %v: %v", filepath, err)
	}

	for _, file := range files {
		seq.textures = append(seq.textures, file.Name())
	}

	seq.sampleRate = sampleRate
	seq.loop = loop
	seq.lastSequence = lastSequence

	return &seq, nil
}

func (seq *sequence) nextFrame() {
	if seq.frame == len(seq.textures)-1 {
		if seq.loop {
			seq.frame = 0
		} else {
			seq.finished = true
		}
	} else {
		seq.frame++
	}
}

func newAnimator(sequences map[ElementState]*sequence, defaultSequence ElementState) *animator {
	var an animator

	an.sequences = sequences
	an.currentSequence = defaultSequence
	an.lastFrameChange = time.Now()

	return &an
}
