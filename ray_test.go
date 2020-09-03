package main

import (
	"reflect"
	"testing"
)

func TestRayAt(t *testing.T) {
	expect := NewPoint3(1.0, 2.0, 3.0)

	orig := NewPoint3(0, 0, 0)
	dir := NewVec3(1, 2, 3)
	ray := NewRay(orig, dir)

	actual := ray.At(1)

	if !reflect.DeepEqual(&expect, &actual) {
		t.Errorf("actual: %v, expected: %v", actual, expect)
	}
}
