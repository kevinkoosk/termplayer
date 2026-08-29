package decoder

import (
	"fmt"
	"image"
	"io"
	"os/exec"
)

type Stream struct {
	cmd    *exec.Cmd
	reader io.Reader
	Width  int
	Height int
	img    *image.RGBA
	buf    []byte
}

func New(video string, maxWidth, maxHeight int) (*Stream, error) {
	// 1. Get source video dimensions first
	videoW, videoH, err := GetDimensions(video)
	if err != nil {
		return nil, err
	}

	// 2. Compute target dimensions using the passed bounds
	targetW, targetH := BestFit(
		videoW,
		videoH,
		maxWidth,
		maxHeight,
	)

	// 3. Allocate buffers now that target dimensions exist
	img := image.NewRGBA(
		image.Rect(0, 0, targetW, targetH),
	)

	buf := make([]byte, targetW*targetH*3)

	// 4. Build FFmpeg command
	vf := fmt.Sprintf("fps=10,scale=%d:%d", targetW, targetH)

	cmd := exec.Command(
		"ffmpeg",
		"-loglevel", "quiet",
		"-i", video,
		"-vf", vf,
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
		cmd:    cmd,
		reader: out,
		Width:  targetW,
		Height: targetH,
		img:    img,
		buf:    buf,
	}, nil
}

func (s *Stream) NextFrame() (image.Image, error) {
	_, err := io.ReadFull(s.reader, s.buf)
	if err != nil {
		return nil, err
	}

	src := 0
	dst := 0

	for src < len(s.buf) {
		s.img.Pix[dst] = s.buf[src]
		s.img.Pix[dst+1] = s.buf[src+1]
		s.img.Pix[dst+2] = s.buf[src+2]
		s.img.Pix[dst+3] = 255

		src += 3
		dst += 4
	}

	return s.img, nil
}

func (s *Stream) Close() {
	if s.cmd != nil && s.cmd.Process != nil {
		_ = s.cmd.Process.Kill()
	}
}