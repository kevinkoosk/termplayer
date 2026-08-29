package picture

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/kevinkoosk/termplayer/internal/renderer"
	"github.com/kevinkoosk/termplayer/internal/terminal"
)

func Load(path string) (image.Image, error) {

	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	img, _, err := image.Decode(f)

	return img, err
}

func Show(path string) error {

	img, err := Load(path)

	if err != nil {
		return err
	}

	w, h := terminal.Size()

	renderer.Draw(
		img,
		w,
		h-1,
	)

	return nil
}