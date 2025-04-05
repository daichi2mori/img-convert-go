package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/gen2brain/avif"
	"github.com/gen2brain/jpegli"
	"github.com/gen2brain/jpegxl"
	"github.com/gen2brain/webp"
	"golang.org/x/image/draw"
)

func main() {
	inputFile := "test.jpg"
	outputPath := `output`
	convert := "jpegxl" // "jpg", "jpegli", "jpegxl", "png", "webp", "avif"
	quality := 50

	// リサイズ関連のパラメータ
	resize := false              // リサイズするかどうか
	resizeWidth := 0             // リサイズ後の幅（0の場合はアスペクト比を維持）
	resizeHeight := 1024         // リサイズ後の高さ（0の場合はアスペクト比を維持）
	resizeMethod := "CatmullRom" // リサイズメソッド: "NearestNeighbor", "ApproxBiLinear", "Bilinear", "CatmullRom"

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

	// リサイズが有効な場合は画像をリサイズ
	if resize {
		img, err = resizeImage(img, resizeWidth, resizeHeight, resizeMethod)
		if err != nil {
			fmt.Printf("画像のリサイズに失敗しました: %v\n", err)
			return
		}
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

	ext := strings.ToLower(filepath.Ext(filePath))
	var img image.Image

	switch ext {
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
		return nil, fmt.Errorf("サポートされていないファイル形式です: %s", ext)
	}
	if err != nil {
		return nil, fmt.Errorf("画像のデコードに失敗しました: %v", err)
	}

	return img, nil
}

// 画像をリサイズする関数
func resizeImage(img image.Image, width, height int, method string) (image.Image, error) {
	bounds := img.Bounds()
	srcWidth := bounds.Dx()
	srcHeight := bounds.Dy()

	// 幅または高さが0の場合、アスペクト比を維持
	if width == 0 && height == 0 {
		return nil, fmt.Errorf("幅と高さの両方が0です")
	} else if width == 0 {
		width = srcWidth * height / srcHeight
	} else if height == 0 {
		height = srcHeight * width / srcWidth
	}

	// 元の画像と同じサイズの場合は何もしない
	if width == srcWidth && height == srcHeight {
		return img, nil
	}

	// 新しい画像を作成
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// リサイズメソッドを選択
	var scaler draw.Scaler
	switch strings.ToLower(method) {
	case "NearestNeighbor":
		scaler = draw.NearestNeighbor
	case "ApproxBiLinear":
		scaler = draw.ApproxBiLinear
	case "Bilinear":
		scaler = draw.BiLinear
	case "CatmullRom":
		scaler = draw.CatmullRom
	default:
		scaler = draw.ApproxBiLinear
	}

	// リサイズを実行
	scaler.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	return dst, nil
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
