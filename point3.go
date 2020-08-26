package main

import (
	"fmt"
)

type IPoint interface {
	IVec3
	String() string
	X() float32
	Y() float32
	Z() float32
}

type point3 struct {
	vec3
}

func NewPoint3(e1, e2, e3 float32) *point3 {
	v := new(point3)
	v.e[0] = e1
	v.e[1] = e2
	v.e[2] = e3
	return v
}

func (v vec3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.e[0], v.e[1], v.e[2])
}

func (p *point3) X() float32 {
	return p.e[0]
}
func (p *point3) Y() float32 {
	return p.e[1]
}
func (p *point3) Z() float32 {
	return p.e[2]
}
