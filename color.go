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

type Color Vec3

func NewColor(e1, e2, e3 float32) *Color {
	v := new(Color)
	v.e[0] = e1
	v.e[1] = e2
	v.e[2] = e3
	return v
}

func (v Color) String() string {
	r := int32(255.999 * v.R())
	g := int32(255.999 * v.G())
	b := int32(255.999 * v.B())
	return fmt.Sprintf("%d %d %d", r, g, b)
}

func (c *Color) R() float32 {
	return c.e[0]
}
func (c *Color) G() float32 {
	return c.e[1]
}
func (c *Color) B() float32 {
	return c.e[2]
}

func (v Color) GetElm(k int) float32 {
	return v.e[k]
}
