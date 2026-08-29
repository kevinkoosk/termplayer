package renderer

import (
	"fmt"
	"image"
	"image/color"
	"strings"
	"github.com/nfnt/resize"
)

func ansiFgBg(fg, bg color.Color) string {
	fr, fgc, fb, _ := fg.RGBA()
	br, bgc, bb, _ := bg.RGBA()

return fmt.Sprintf(
		"\x1b[38;2;%d;%d;%dm\x1b[48;2;%d;%d;%dm",
		fr>>8, fgc>>8, fb>>8,
		br>>8, bgc>>8, bb>>8,
	)
}

func Draw(img image.Image, width, height int) {

	resized :=
		resize.Resize(
			uint(width),
			uint(height*2),
			img,
			resize.Lanczos3,
		)

	var sb strings.Builder
	sb.WriteString("\x1b[H")

	for y := 0; y < resized.Bounds().Dy(); y += 2 {

		for x := 0; x < resized.Bounds().Dx(); x++ {

			top := resized.At(x, y)

			var bottom color.Color = color.Black

			if y+1 < resized.Bounds().Dy() {
			bottom = resized.At(x, y+1)
			}

			sb.WriteString(
				ansiFgBg(top, bottom),
			)
			sb.WriteRune('▀')
		}

		sb.WriteString(("\x1b[0m\n"))
	}
	fmt.Print(sb.String())
}