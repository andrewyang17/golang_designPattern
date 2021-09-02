package main

import "fmt"

type Switch struct {
	State State
}

func NewSwitch() *Switch {
	return &Switch{NewOffState()}
}

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

type State interface {
	On(*Switch)
	Off(*Switch)
}

type BaseState struct {}

func (*BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

func (*BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}

type OnState struct {
	BaseState
}

func NewOnState() *OnState {
	return &OnState{BaseState: BaseState{}}
}

func (o *OnState) Off(sw *Switch) {
	fmt.Println("Turning light off...")
	sw.State = NewOffState()
}

type OffState struct {
	BaseState
}

func NewOffState() *OffState {
	return &OffState{BaseState: BaseState{}}
}

func (o *OffState) On(sw *Switch) {
	fmt.Println("Turning light on...")
	sw.State = NewOnState()
}

func main() {
	sw := NewSwitch()
	sw.On()
	sw.On()
	sw.Off()
	sw.Off()
}


