package main

import (
	"math"
)

type IVec3 interface {
	Add() *vec3
	Minus() *vec3
	Times() *vec3

	Dot() float32
	Length() float32 
	LengthSquared() float32
}

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

