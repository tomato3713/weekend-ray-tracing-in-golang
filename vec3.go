package main

import (
	"math"
)

type IVec3 interface {
	Add() Vec3
	Minus() Vec3
	Times() Vec3

	Dot() float32
	Length() float32
	LengthSquared() float32
}

type Vec3 struct {
	e [3]float32
}

func NewVec3(e1, e2, e3 float32) *Vec3 {
	v := new(Vec3)
	v.e[0] = e1
	v.e[1] = e2
	v.e[2] = e3
	return v
}

func (v1 Vec3) Add(v2 Vec3) Vec3 {
	v := NewVec3(v1.e[0]+v2.e[0], v1.e[1]+v2.e[1], v1.e[2]+v2.e[2])
	return *v
}

func (v1 Vec3) Minus(v2 Vec3) Vec3 {
	v := NewVec3(v1.e[0]-v2.e[0], v1.e[1]-v2.e[1], v1.e[2]-v2.e[2])
	return *v
}

func (v1 Vec3) Times(v2 Vec3) Vec3 {
	v := NewVec3(v1.e[0]*v2.e[0], v1.e[1]*v2.e[1], v1.e[2]*v2.e[2])
	return *v
}

func (v1 *Vec3) Dot(v2 Vec3) float32 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v2.e[2]*v2.e[2]
}

func (v *Vec3) Length() float32 {
	ls := float64(v.LengthSquared())
	return float32(math.Sqrt(ls))
}

func (v *Vec3) LengthSquared() float32 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}
