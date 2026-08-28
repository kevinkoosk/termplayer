package decoder

import (
	"bufio"
	"image"
	"image/png"
	"os/exec"
)

type Stream struct {
	cmd *exec.Cmd
	reader *bufio.Reader
}

func New(video string) (*Stream, error) {

	cmd := exec.Command(
		"ffmpeg",
		"-loglevel", "quiet",
		"-i", video,
		"-vf", "fps=12",
		"-f", "image2pipe",
		"-vcodec", "png",
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
		reader: bufio.NewReader(out),
	}, nil
}

func (s *Stream) NextFrame() (image.Image, error) {
	return png.Decode(s.reader)
}

func (s *Stream) Close() {
	_ = s.cmd.Process.Kill()
}