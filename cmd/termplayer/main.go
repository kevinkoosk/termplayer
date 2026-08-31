package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"strconv"

	"github.com/kevinkoosk/termplayer/internal/picture"
	"github.com/kevinkoosk/termplayer/internal/player"
	"github.com/kevinkoosk/termplayer/internal/config"
)

func main() {

	if len(os.Args) < 2 {

		fmt.Println(
			"usage: termplayer file",
		)

		return
	}
	
	for i := 1; i < len(os.Args); i++ {

	if os.Args[i] == "--scale" &&
		i+1 < len(os.Args) {

		scale, err :=
			strconv.Atoi(
				os.Args[i+1],
			)

		if err == nil {

			if scale < 25 {
				scale = 25
			}

			if scale > 150 {
				scale = 150
			}

			config.ScalePercent =
				scale
		}
	}
}

	ext := strings.ToLower(
		filepath.Ext(
			os.Args[1],
		),
	)

	switch ext {

	case ".jpg",
		".jpeg",
		".png",
		".gif":

		err := picture.Show(
			os.Args[1],
		)

		if err != nil {
			panic(err)
		}

	default:

		err := player.Play(
			os.Args[1],
		)

		if err != nil {
			panic(err)
		}
	}
}