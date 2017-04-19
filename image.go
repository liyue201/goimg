package goimg

import (
	"errors"
	"fmt"
	"github.com/chai2010/webp"
	"golang.org/x/image/bmp"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

const (
	BMP  = "bmp"
	PNG  = "png"
	JPG  = "jpg"
	WEBP = "webp"
)

func Convert(srcPath, destPath, format string, quality float32) error {
	r, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer w.Close()

	return ConvertFormat(r, w, format, quality)
}

func ConvertFormat(r io.Reader, w io.ReadWriter, format string, quality float32) error {
	srcImage, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	switch format {
	case BMP:
		return Encode2Bmp(srcImage, w)
	case PNG:
		return Encode2Png(srcImage, w)
	case JPG:
		return Encode2Jpeg(srcImage, w, quality)
	case WEBP:
		return Encode2WebP(srcImage, w, quality)
	default:
		return errors.New(fmt.Sprintf("Unsupported format:", format))
	}
	return nil
}

func Encode2Jpeg(img image.Image, w io.Writer, quality float32) error {
	o := jpeg.Options{}
	o.Quality = int(quality)
	return jpeg.Encode(w, img, &o)
}

func Encode2WebP(img image.Image, w io.Writer, quality float32) error {
	o := webp.Options{}
	o.Lossless = false
	o.Quality = quality
	return webp.Encode(w, img, &o)
}

func Encode2Bmp(img image.Image, w io.Writer) error {
	return bmp.Encode(w, img)
}

func Encode2Png(img image.Image, w io.Writer) error {
	return png.Encode(w, img)
}
