package main

import (
	"math"
)

type Hittable interface {
	Hit(Ray, float32, float32, HitRecord) bool
}

type IHitRecord interface {
	setFaceNormal(Ray, Vec3)
}

type HitRecord struct {
	p          point3
	normal     Vec3
	t          float32
	front_face bool
}

type Sphere struct {
	Center point3
	Radius float32
}

func (obj Sphere) Hit(r Ray, t_min, t_max float32, rec *HitRecord) bool {
	oc := Minus(r.Orig, obj.Center)
	a := LengthSquared(r.Dir)
	half_b := Dot(oc, r.Dir)
	c := LengthSquared(oc) - obj.Radius*obj.Radius
	discriminant := half_b*half_b - a*c

	if discriminant > 0 {
		root := float32(math.Sqrt(float64(discriminant)))
		temp := (-half_b - root) / a
		if temp < t_max && temp > t_min {
			rec.t = temp
			rec.p = r.At(rec.t)
			outward_normal := DivedC(Minus(rec.p, obj.Center), obj.Radius)
			rec.setFaceNormal(r, outward_normal)
			return true
		}
		temp = (-half_b + root) / a
		if temp < t_max && temp > t_min {
			rec.t = temp
			rec.p = r.At(rec.t)
			outward_normal := DivedC(Minus(rec.p, obj.Center), obj.Radius)
			rec.setFaceNormal(r, outward_normal)
			return true
		}
	}
	return false
}

func (hr *HitRecord) setFaceNormal(r Ray, outward_normal Vec3) {
	hr.front_face = Dot(r.Dir, outward_normal) < 0
	if hr.front_face {
		hr.normal = outward_normal
	} else {
		hr.normal = TimesC(outward_normal, -1)
	}
}
