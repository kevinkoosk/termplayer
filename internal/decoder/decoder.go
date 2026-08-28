package decoder

import (
    "image"
    "io"
    "os/exec"
)

const (
    Width  = 160
    Height = 90
)

type Stream struct {
    cmd    *exec.Cmd
    reader io.Reader
    img    *image.RGBA
    buf    []byte
}

func New(video string) (*Stream, error) {

    img := image.NewRGBA(
        image.Rect(
            0,
            0,
            Width,
            Height,
        ),
    )

    buf := make(
        []byte,
        Width*Height*3,
    )

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
        cmd:    cmd,
        reader: out,
        img:    img,
        buf:    buf,
    }, nil
}

func (s *Stream) NextFrame() (image.Image, error) {

    _, err := io.ReadFull(
        s.reader,
        s.buf,
    )

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

    if s.cmd != nil &&
        s.cmd.Process != nil {

        _ = s.cmd.Process.Kill()
    }
}