package decoder

import (
	"image"
	"image/color"
	"io"
	"os/exec"
)

const (
	Width = 160
	Height = 90
)

type Stream struct {
	cmd *exec.Cmd
	reader io.Reader
}

func New(video string) (*Stream, error) {

	cmd := exec.Command(
		"ffmpeg",
		"-loglevel", "quiet",
		"-i", video,
		"-vf", "fps=10,scale=160:90",
		"-pix_fmt", "rgb24",
		"-f", "rawvideo",
		"pipe:1",
	)

	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return &Stream{
		cmd: cmd,
		reader: out,
	}, nil
}

func (s *Stream) NextFrame() (image.Image, error) {
	frameSize := Width * Height * 3
	buf := make([]byte, frameSize)
	_, err := io.ReadFull(s.reader, buf)
	if err != nil {
		return nil, err
	}

	img := image.NewRGBA(
		image.Rect(
			0,
			0,
			Width,
			Height,
		),
	)

	p := 0

	for y := 0; y < Height; y++ {

		for x := 0; x < Width; x++ {

			r := buf[p]
			g := buf[p+1]
			b := buf[p+2]

			p += 3

			img.Set(
				x,
				y,
				color.RGBA{
					R: r,
					G: g,
					B: b,
					A: 255,
				},
			)
		}
	}

	return img, nil
}

func (s *Stream) Close() {
	if s.cmd != nil &&
		s.cmd.Process != nil {

		_ = s.cmd.Process.Kill()
	}
}