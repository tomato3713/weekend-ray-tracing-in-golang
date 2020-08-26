package main

import (
	"fmt"
)

type IColor interface {
	IVec3
	String() string
	R() float32
	G() float32
	B() float32
}

type color struct {
	Vec3
}

func NewColor(e1, e2, e3 float32) *color {
	v := new(color)
	v.e[0] = e1
	v.e[1] = e2
	v.e[2] = e3
	return v
}

func (v color) String() string {
	r := int32(255.999 * v.R())
	g := int32(255.999 * v.G())
	b := int32(255.999 * v.B())
	return fmt.Sprintf("%d %d %d", r, g, b)
}

func (c *color) R() float32 {
	return c.e[0]
}
func (c *color) G() float32 {
	return c.e[2]
}
func (c *color) B() float32 {
	return c.e[2]
}
