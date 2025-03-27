package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/gen2brain/avif"
	"github.com/gen2brain/jpegli"
	"github.com/gen2brain/jpegxl"
	"github.com/gen2brain/webp"
)

func main() {
	inputFile := "test.jpg"
	outputPath := `C:\Users\user\Desktop\test`
	convert := "jpegxl"
	quality := 50

	jpgOptions := jpeg.Options{
		Quality: quality,
	}

	jpegliOptions := jpegli.EncodingOptions{
		Quality:              quality,
		ChromaSubsampling:    image.YCbCrSubsampleRatio420,
		ProgressiveLevel:     0,
		OptimizeCoding:       true,
		AdaptiveQuantization: true,
		StandardQuantTables:  false,
		FancyDownsampling:    true,
		DCTMethod:            jpegli.DCTFloat,
	}

	jpegxlOptions := jpegxl.Options{
		Quality: quality,
		Effort:  10,
	}

	pngEncoder := png.Encoder{
		CompressionLevel: png.BestCompression,
	}

	webpOptions := webp.Options{
		Quality:  quality,
		Method:   6,
		Lossless: false,
		Exact:    false,
	}

	avifOptions := avif.Options{
		Quality:           quality,
		QualityAlpha:      0,
		Speed:             0,
		ChromaSubsampling: image.YCbCrSubsampleRatio420,
	}

	img, err := loadImage(inputFile)
	if err != nil {
		fmt.Printf("画像の読み込みに失敗しました: %v\n", err)
		return
	}

	switch convert {
	case "jpg", "jpeg":
		convertToJPEG(inputFile, outputPath, img, jpgOptions)
	case "jpegli":
		convertToJPEGLI(inputFile, outputPath, img, jpegliOptions)
	case "jpegxl":
		convertToJPEGXL(inputFile, outputPath, img, jpegxlOptions)
	case "png":
		convertToPNG(inputFile, outputPath, img, pngEncoder)
	case "webp":
		convertToWEBP(inputFile, outputPath, img, webpOptions)
	case "avif":
		convertToAVIF(inputFile, outputPath, img, avifOptions)
	default:
		fmt.Println("無効な変換タイプです")
		return
	}

	fmt.Println("変換が正常に完了しました")
}

// 画像読み込み
func loadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	extension := filepath.Ext(filePath)
	var img image.Image

	switch extension {
	case ".jpg", ".jpeg":
		img, err = jpegli.Decode(file)
	case ".jxl":
		img, err = jpegxl.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	case ".webp":
		img, err = webp.Decode(file)
	case ".avif":
		img, err = avif.Decode(file)
	default:
		return nil, fmt.Errorf("サポートされていないファイル形式です: %s", extension)
	}
	if err != nil {
		return nil, fmt.Errorf("画像のデコードに失敗しました: %v", err)
	}

	return img, nil
}

// 出力ファイル名を生成
func generateOutputFilename(inputFile, outputDir, extension string) string {
	fileName := filepath.Base(inputFile[:len(inputFile)-len(filepath.Ext(inputFile))])
	return filepath.Join(outputDir, fileName+extension)
}

// JPEGに変換
func convertToJPEG(inputFile, outputDir string, img image.Image, options jpeg.Options) error {
	outputPath := generateOutputFilename(inputFile, outputDir, ".jpg")

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if err := jpeg.Encode(outFile, img, &options); err != nil {
		return fmt.Errorf("JPEGエンコードに失敗しました: %w", err)
	}
	return nil
}

// JPEG LIに変換
func convertToJPEGLI(inputFile, outputDir string, img image.Image, options jpegli.EncodingOptions) error {
	outputPath := generateOutputFilename(inputFile, outputDir, ".jpg")

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("出力ファイルを作成できません: %w", err)
	}
	defer outFile.Close()

	if err := jpegli.Encode(outFile, img, &options); err != nil {
		return fmt.Errorf("JPEG LIエンコードに失敗しました: %w", err)
	}
	return nil
}

// JPEG XLに変換
func convertToJPEGXL(inputFile, outputDir string, img image.Image, option jpegxl.Options) error {
	outputPath := generateOutputFilename(inputFile, outputDir, ".jxl")

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("出力ファイルを作成できません: %w", err)
	}
	defer outFile.Close()

	if err := jpegxl.Encode(outFile, img, option); err != nil {
		return fmt.Errorf("JPEG XLエンコードに失敗しました: %w", err)
	}
	return nil
}

// PNGに変換
func convertToPNG(inputFile, outputDir string, img image.Image, pngEncoder png.Encoder) error {
	outputPath := generateOutputFilename(inputFile, outputDir, ".png")

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("出力ファイルを作成できません: %w", err)
	}
	defer outFile.Close()

	if err := pngEncoder.Encode(outFile, img); err != nil {
		return fmt.Errorf("PNGエンコードに失敗しました: %w", err)
	}
	return nil
}

// WEBPに変換
func convertToWEBP(inputFile, outputDir string, img image.Image, options webp.Options) error {
	outputPath := generateOutputFilename(inputFile, outputDir, ".webp")

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("出力ファイルを作成できません: %w", err)
	}
	defer outFile.Close()

	if err := webp.Encode(outFile, img, options); err != nil {
		return fmt.Errorf("WEBPエンコードに失敗しました: %w", err)
	}
	return nil
}

// AVIFに変換
func convertToAVIF(inputFile, outputDir string, img image.Image, options avif.Options) error {
	outputPath := generateOutputFilename(inputFile, outputDir, ".avif")

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("出力ファイルを作成できません: %w", err)
	}
	defer outFile.Close()

	if err := avif.Encode(outFile, img, options); err != nil {
		return fmt.Errorf("AVIFエンコードに失敗しました: %w", err)
	}
	return nil
}
