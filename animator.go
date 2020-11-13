package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

type animator struct {
	sequences       map[string]*sequence
	currentSequence string
	lastFrameChange time.Time
	finished        bool
}

type sequence struct {
	textures   []string
	frame      int
	sampleRate float64
	loop       bool
}

func newAnimator(sequences map[string]*sequence, defaultSequence string) *animator {
	var an animator

	an.sequences = sequences
	an.currentSequence = defaultSequence
	an.lastFrameChange = time.Now()

	return &an
}

func (an *animator) onUpdate(parameters updateParameters) error {
	sequence := an.sequences[an.currentSequence]
	frameInterval := float64(time.Second) / sequence.sampleRate

	if time.Since(an.lastFrameChange) >= time.Duration(frameInterval) {
		an.finished = sequence.nextFrame()
		an.lastFrameChange = time.Now()
	}
	return nil
}

func (an *animator) setSequence(name string) {
	an.currentSequence = name
	an.lastFrameChange = time.Now()
}

func newSequence(
	filepath string,
	sampleRate float64,
	loop bool) (*sequence, error) {

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

	return &seq, nil
}

func (seq *sequence) nextFrame() bool {
	if seq.frame == len(seq.textures)-1 {
		if seq.loop {
			seq.frame = 0
		} else {
			return true
		}
	} else {
		seq.frame++
	}

	return false
}
