package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/gen2brain/avif"
	"github.com/gen2brain/webp"
)

func main() {
	inputFile := "test.jpg"
	quality := 50
	method := 6       // 0~6 大きいほど圧縮率が良い
	lossless := false // trueで画像が劣化しないようにする。圧縮率はかなり悪化する
	exact := false    // 透明な領域でも正確なRGB値を保持する

	webpOptions := webp.Options{
		Quality:  quality,
		Method:   method,
		Lossless: lossless,
		Exact:    exact,
	}

	avifOptions := avif.Options{
		Quality:           quality,
		Speed:             0, // 0~10 小さいほうが圧縮率が良い
		ChromaSubsampling: image.YCbCrSubsampleRatio420,
	}

	if err := convertToWEBP(inputFile, "test.webp", webpOptions); err != nil {
		fmt.Println("Conversion failed:", err)
		return
	}
	fmt.Println("webp success")

	if err := convertToAVIF(inputFile, "test.avif", avifOptions); err != nil {
		fmt.Println("Conversion failed:", err)
		return
	}
	fmt.Println("avif success")
}

func convertToWEBP(inputFile, outputFile string, options webp.Options) error {
	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	var img image.Image
	switch {
	case strings.HasSuffix(strings.ToLower(inputFile), ".png"):
		img, err = png.Decode(inFile)
	case strings.HasSuffix(strings.ToLower(inputFile), ".jpg"):
		img, err = jpeg.Decode(inFile)
	case strings.HasSuffix(strings.ToLower(inputFile), ".jpeg"):
		img, err = jpeg.Decode(inFile)
	case strings.HasSuffix(strings.ToLower(inputFile), ".gif"):
		img, err = gif.Decode(inFile)
	default:
		return err
	}
	if err != nil {
		return err
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if err := webp.Encode(outFile, img, options); err != nil {
		return err
	}
	return nil
}

func convertToAVIF(inputFile, outputFile string, options avif.Options) error {
	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	var img image.Image
	switch {
	case strings.HasSuffix(strings.ToLower(inputFile), ".png"):
		img, err = png.Decode(inFile)
	case strings.HasSuffix(strings.ToLower(inputFile), ".jpg"):
		img, err = jpeg.Decode(inFile)
	case strings.HasSuffix(strings.ToLower(inputFile), ".jpeg"):
		img, err = jpeg.Decode(inFile)
	case strings.HasSuffix(strings.ToLower(inputFile), ".gif"):
		img, err = gif.Decode(inFile)
	default:
		return err
	}
	if err != nil {
		return err
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if err := avif.Encode(outFile, img, options); err != nil {
		return err
	}
	return nil
}
