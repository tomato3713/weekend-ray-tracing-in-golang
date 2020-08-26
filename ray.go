package main

type IRay interface {
	At() point3
}

type Ray struct {
	Orig point3
	Dir  Vec3
}

func NewRay(orig point3, dir Vec3) *Ray {
	ray := new(Ray)
	ray.Orig = orig
	ray.Dir = dir
	return ray
}

func (r Ray) At(t float32) *point3 {
	e := NewVec3(t, t, t)
	v := Times(r.Dir, *e)

	p := NewPoint3(r.Orig.X()+v.e[0], r.Orig.Y()+v.e[1], r.Orig.Z()+v.e[2])
	return p
}
