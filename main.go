package main

import (
	"fmt"
	"math"
)

// REF: https://github.com/RayTracing/raytracing.github.io/blob/7b92650a045a887391bf8148af86c5e6c46c7a9e/books/RayTracingInOneWeekend.html#L77-L105

// NOTE:
// If you use PowerShell, You should use `go run ppm.go | Out-File -FilePath image.ppm -Encoding ASCII`
// PowerShell's redirection is output file in UTF16 encoding by default.
// REF: https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_redirection?view=powershell-7#notes

func hit_sphere(center *point3, radius float32, r *Ray) float32 {
	oc := Minus(r.Orig, center)
	a := LengthSquared(r.Dir)
	half_b := Dot(oc, r.Dir)
	c := LengthSquared(oc) - radius*radius
	discriminant := float64(half_b*half_b - a*c)
	if discriminant > 0 {
		// hit!!
		return (-half_b - float32(math.Sqrt(discriminant))) / a
	} else {
		// no hit!!
		return -1.0
	}
}

func ray_color(r *Ray) color {
	t := hit_sphere(NewPoint3(0, 0, -1), 0.5, r)
	if t > 0.0 {
		N := UnitVector(Minus(r.At(t), NewVec3(0, 0, -1)))
		return color(TimesC(NewColor(N.GetElm(0)+1, N.GetElm(1)+1, N.GetElm(2)+1), 0.5))
	}
	// no hit
	unit_direction := point3(UnitVector(r.Dir))
	t = 0.5 * (unit_direction.Y() + 1.0)
	c1 := NewColor(1.0, 1.0, 1.0)
	c2 := NewColor(0.5, 0.7, 1.0)
	v := Add(TimesC(*c1, 1.0-t), TimesC(*c2, t))
	return color(v)
}

func main() {
	aspect_ratio := 16.0 / 9.0
	image_width := 256
	image_height := math.Floor(float64(image_width) / aspect_ratio)

	fmt.Println("P3")
	fmt.Println(image_width, image_height)
	fmt.Println("255")

	viewport_height := 2.0
	viewport_width := aspect_ratio * viewport_height
	focal_length := 1.0

	origin := NewPoint3(0, 0, 0)
	horizontal := NewVec3(float32(viewport_width), 0, 0)
	vertical := NewVec3(0, float32(viewport_height), 0)

	lower_left_corver := Minus(Minus(Minus(origin, DivedC(horizontal, 2)), DivedC(vertical, 2)), *NewVec3(0, 0, float32(focal_length)))

	for j := image_height - 1; j >= 0; j-- {
		fmt.Errorf("\rScanlines remaining: %d", j)
		for i := 0; i < image_width; i++ {
			u := float32(i) / float32(image_width-1)
			v := float32(j) / float32(image_height-1)

			r := NewRay(*origin, Add(Add(lower_left_corver, TimesC(horizontal, u)), Minus(TimesC(vertical, v), origin)))

			pixel_color := ray_color(r)
			fmt.Println(pixel_color)
		}
	}
	fmt.Errorf("\nDone.")
}
