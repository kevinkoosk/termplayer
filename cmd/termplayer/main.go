package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kevinkoosk/termplayer/internal/config"
	"github.com/kevinkoosk/termplayer/internal/picture"
	"github.com/kevinkoosk/termplayer/internal/player"
)

func main() {
	file, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "usage: termplayer [--loop] [--scale percent] file")
		os.Exit(2)
	}

	ext := strings.ToLower(filepath.Ext(file))

	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		if err := picture.Show(file); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		for {
			err := player.Play(file)
			if errors.Is(err, player.ErrQuit) {
				return
			}
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			if !config.LoopPlayback {
				return
			}
		}
	}
}

func parseArgs(args []string) (string, error) {
	var file string

	for i := 0; i < len(args); i++ {
		arg := args[i]

		switch {
		case arg == "--":
			// everything after -- is the file path, even if it looks like a flag
			if i+1 >= len(args) {
				return "", fmt.Errorf("missing file")
			}
			if file != "" {
				return "", fmt.Errorf("unexpected extra argument: %s", args[i+1])
			}
			return args[i+1], nil

		case arg == "--loop":
			config.LoopPlayback = true

		case arg == "--scale":
			if i+1 >= len(args) {
				return "", fmt.Errorf("--scale needs a number")
			}
			i++
			scale, err := strconv.Atoi(args[i])
			if err != nil {
				return "", fmt.Errorf("invalid --scale: %s", args[i])
			}
			if scale < 25 {
				scale = 25
			}
			if scale > 150 {
				scale = 150
			}
			config.ScalePercent = scale

		case strings.HasPrefix(arg, "-"):
			return "", fmt.Errorf("unknown flag: %s", arg)

		default:
			if file != "" {
				return "", fmt.Errorf("unexpected extra argument: %s", arg)
			}
			file = arg
		}
	}

	if file == "" {
		return "", fmt.Errorf("missing file")
	}
	return file, nil
}