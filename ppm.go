package main

import (
	"fmt"
)

// REF: https://github.com/RayTracing/raytracing.github.io/blob/7b92650a045a887391bf8148af86c5e6c46c7a9e/books/RayTracingInOneWeekend.html#L77-L105

// NOTE:
// If you use PowerShell, You should use `go run ppm.go | Out-File -FilePath image.ppm -Encoding ASCII`
// PowerShell's redirection is output file in UTF16 encoding by default.
// REF: https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_redirection?view=powershell-7#notes

func main() {
	image_width := 256
	image_height := 256

	fmt.Println("P3")
	fmt.Println(image_width, image_height)
	fmt.Println("255")

	for j := image_height - 1; j >= 0; j-- {
		fmt.Errorf("\rScanlines remaining: %d", j)
		for i := 0; i < image_width; i++ {
			r := float32(i) / float32(image_width - 1)
			g := float32(j) / float32(image_height - 1)
			b := 0.25

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			fmt.Println(ir, ig, ib)
		}
	}
	fmt.Errorf("\nDone.")
}
