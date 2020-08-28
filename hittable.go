package main

type Hittable interface {
	hit(Ray, float32, float32, HitRecord) bool
}

type HitRecord struct {
	p      point3
	normal Vec3
	t      float32
}

type Sphere struct {
	center point3
	radius float32
}
