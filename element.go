package main

import (
	"fmt"
	"reflect"
)

type vector struct {
	x, y float64
}

type uiComponent interface {
	onDraw(parameters drawParameters) error
}

type logicComponent interface {
	onUpdate(parameters updateParameters) error
}

type updatableElement interface {
	onComponentsUpdated()
}

type element struct {
	position        vector
	width           float64
	rotation        float64
	active          bool
	maxPosition     vector
	logicComponents []logicComponent
	uiComponents    []uiComponent
}

func (elem *element) update(updateParameters updateParameters) error {
	for _, comp := range elem.logicComponents {
		err := comp.onUpdate(updateParameters)
		if err != nil {
			return err
		}
	}
	elem.onComponentsUpdated()
	return nil
}

func (elem *element) draw(parameters drawParameters) error {
	for _, comp := range elem.uiComponents {
		err := comp.onDraw(parameters)
		if err != nil {
			return err
		}
	}

	return nil
}

func (elem *element) addUiComponent(component uiComponent) {
	for _, existing := range elem.uiComponents {
		if reflect.TypeOf(component) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attempt to add new component with existing type %v",
				reflect.TypeOf(component)))
		}
	}
	elem.uiComponents = append(elem.uiComponents, component)
}

func (elem *element) getUiComponent(withType uiComponent) uiComponent {
	typ := reflect.TypeOf(withType)
	for _, comp := range elem.uiComponents {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}

	panic(fmt.Sprintf("no ui component with type %v", reflect.TypeOf(withType)))
}

func (elem *element) addLogicComponent(component logicComponent) {
	for _, existing := range elem.logicComponents {
		if reflect.TypeOf(component) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attempt to add new component with existing type %v",
				reflect.TypeOf(component)))
		}
	}
	elem.logicComponents = append(elem.logicComponents, component)
}

func (elem *element) getLogicComponent(withType logicComponent) logicComponent {
	typ := reflect.TypeOf(withType)
	for _, comp := range elem.logicComponents {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}

	panic(fmt.Sprintf("no logic component with type %v", reflect.TypeOf(withType)))
}

var elements []*element
