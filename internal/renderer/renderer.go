package renderer

import (
	"fmt"
	"image"
	"image/color"
	"strings"
	"github.com/nfnt/resize"
	"github.com/kevinkoosk/termplayer/internal/config"
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

	imgW := img.Bounds().Dx()
	imgH := img.Bounds().Dy()
	
	scale :=
	float64(
		config.ScalePercent,
	) / 100.0
	
	ratio :=
		float64(imgW) /
			float64(imgH)
	renderW := width
	
	renderH :=
		int(
			float64(renderW) /
				ratio,
	)
	
	
	if renderH > height*2 {

		renderH = height * 2

		renderW =
			int(
				float64(renderH) * 
					ratio,
			)
	}

	renderW =
	int(
		float64(renderW) *
			scale,
	)

	renderH =
		int(
			float64(renderH) *
				scale,
		)

		if renderW < 1 {
		renderW = 1
	}

	if renderH < 1 {
		renderH = 1
	}

	resized :=
		resize.Resize(
			uint(renderW),
			uint(renderH),
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