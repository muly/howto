package main

import (
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	src, err := imaging.Open("1600x1600_3.9MB_sample.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	dstImage := imaging.Resize(src, 400, 400, imaging.Lanczos)

	err = imaging.Save(dstImage, "out400x400.png")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
