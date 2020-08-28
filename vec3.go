package main

import (
	"math"
)

type IVec3 interface {
	GetElm(int) float32
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

func (v Vec3) GetElm(k int) float32 {
	return v.e[k]
}

func Add(v1, v2 IVec3) Vec3 {
	v := NewVec3(v1.GetElm(0)+v2.GetElm(0), v1.GetElm(1)+v2.GetElm(1), v1.GetElm(2)+v2.GetElm(2))
	return *v
}

func Minus(v1, v2 IVec3) Vec3 {
	v := NewVec3(v1.GetElm(0)-v2.GetElm(0), v1.GetElm(1)-v2.GetElm(1), v1.GetElm(2)-v2.GetElm(2))
	return *v
}

func Times(v1, v2 IVec3) Vec3 {
	v := NewVec3(v1.GetElm(0)*v2.GetElm(0), v1.GetElm(1)*v2.GetElm(1), v1.GetElm(2)*v2.GetElm(2))
	return *v
}

func Dived(v1, v2 IVec3) Vec3 {
	v := NewVec3(v1.GetElm(0)/v2.GetElm(0), v1.GetElm(1)/v2.GetElm(1), v1.GetElm(2)/v2.GetElm(2))
	return *v
}

func Dot(v1, v2 IVec3) float32 {
	return v1.GetElm(0)*v2.GetElm(0) + v1.GetElm(1)*v2.GetElm(1) + v1.GetElm(2)*v2.GetElm(2)
}

func Length(v IVec3) float32 {
	ls := float64(LengthSquared(v))
	return float32(math.Sqrt(ls))
}

func LengthSquared(v IVec3) float32 {
	return v.GetElm(0)*v.GetElm(0) + v.GetElm(1)*v.GetElm(1) + v.GetElm(2)*v.GetElm(2)
}

func UnitVector(v IVec3) Vec3 {
	l := Length(v)
	e := NewVec3(l, l, l)
	return Dived(v, *e)
}

func TimesC(v IVec3, k float32) Vec3 {
	r := NewVec3(v.GetElm(0)*k, v.GetElm(1)*k, v.GetElm(2)*k)
	return *r
}
func DivedC(v IVec3, k float32) Vec3 {
	r := NewVec3(v.GetElm(0)/k, v.GetElm(1)/k, v.GetElm(2)/k)
	return *r
}

func AddC(v IVec3, k float32) Vec3 {
	r := NewVec3(v.GetElm(0)+k, v.GetElm(1)+k, v.GetElm(2)+k)
	return *r
}
