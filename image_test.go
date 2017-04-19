package goimg

import(
	"testing"
)

func TestConvert(t *testing.T) {
	err := Convert("1.jpg", "1.png", PNG, 100)
	if err != nil {
		t.Error(err)
	}

	err = Convert("1.jpg", "1.bmp", BMP, 0)
	if err != nil {
		t.Error(err)
	}

	err = Convert("1.jpg", "1-1.jpg", JPG, 70)
	if err != nil {
		t.Error(err)
	}

	err = Convert("1.jpg", "1.webp", WEBP, 100)
	if err != nil {
		t.Error(err)
	}
}