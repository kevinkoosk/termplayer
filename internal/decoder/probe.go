package decoder

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetDimensions(video string) (int, int, error) {

	cmd := exec.Command(
		"ffprobe",
		"-v", "error",
		"-select_streams", "v:0",
		"-show_entries",
		"stream=width,height",
		"-of",
		"csv=s=x:p=0",
		video,
	)

	out, err := cmd.Output()

	if err != nil {
		return 0, 0, err
	}

	parts :=
		strings.Split(
			strings.TrimSpace(string(out)),
			"x",
		)

	w, err := strconv.Atoi(parts[0])

	if err != nil {
		return 0, 0, err
	}

	h, err := strconv.Atoi(parts[1])

	if err != nil {
		return 0, 0, err
	}

	return w, h, nil
}