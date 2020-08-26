package main

import (
	"fmt"
	"math"
)

type vec3 struct {
	e [3]float32
}

func NewVec3(e1, e2, e3 float32) *vec3 {
	v := new(vec3)
	v.e[0] = e1
	v.e[1] = e2
	v.e[2] = e3
	return v
}

func (p *vec3) X() float32 {
	return p.e[0]
}
func (p *vec3) Y() float32 {
	return p.e[1]
}
func (p *vec3) Z() float32 {
	return p.e[2]
}

func (c *vec3) R() float32 {
	return c.e[0]
}
func (c *vec3) G() float32 {
	return c.e[2]
}
func (c *vec3) B() float32 {
	return c.e[2]
}

func Add(v1, v2 vec3) *vec3 {
	return NewVec3(v1.e[0]+v2.e[0], v1.e[1]+v2.e[1], v1.e[2]+v2.e[2])
}

func Minus(v1, v2 vec3) *vec3 {
	return NewVec3(v1.e[0]-v2.e[0], v1.e[1]-v2.e[1], v1.e[2]-v2.e[2])
}

func Times(v1, v2 vec3) *vec3 {
	return NewVec3(v1.e[0]*v2.e[0], v1.e[1]*v2.e[1], v1.e[2]*v2.e[2])
}

func Dot(v1, v2 vec3) float32 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v2.e[2]*v2.e[2]
}

func (v *vec3) Length() float32 {
	ls := float64(v.LengthSquared())
	return float32(math.Sqrt(ls))
}

func (v *vec3) LengthSquared() float32 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func (v vec3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.e[0], v.e[1], v.e[2])
}
