package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	src, err := os.Open("1600x1600_3.9MB_sample.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	b, err := ioutil.ReadAll(src)
	if err != nil {
		log.Fatalf("failed to read image file: %v", err)
	}

	outB, err := resizePng(b)
	if err != nil {
		log.Fatalf("failed to resize image: %v", err)
	}

	outFile, err := os.Create("400x400.png")
	if err != nil {
		log.Fatalf("failed to open output image: %v", err)
	}
	_, err = outFile.Write(outB)
	if err != nil {
		log.Fatalf("failed to write output image: %v", err)
	}
}

func resizePng(input []byte) ([]byte, error) {
	img, err := imaging.Decode(bytes.NewBuffer(input))
	if err != nil {
		return nil, err
	}

	dstImage := imaging.Resize(img, 400, 400, imaging.Lanczos)

	out := bytes.Buffer{}
	err = imaging.Encode(&out, dstImage, imaging.PNG)
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
