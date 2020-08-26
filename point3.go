package main

import (
	"fmt"
)

type IPoint interface {
	String() string
	X() float32
	Y() float32
	Z() float32
}

type point struct {
	vec3
}

func NewPoint(e1, e2, e3 float32) *point {
	v := new(point)
	v.e[0] = e1
	v.e[1] = e2
	v.e[2] = e3
	return v
}

func (v vec3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.e[0], v.e[1], v.e[2])
}


func (p *point) X() float32 {
	return p.e[0]
}
func (p *point) Y() float32 {
	return p.e[1]
}
func (p *point) Z() float32 {
	return p.e[2]
}