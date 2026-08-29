package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kevinkoosk/termplayer/internal/picture"
	"github.com/kevinkoosk/termplayer/internal/player"
)

func main() {

	if len(os.Args) < 2 {

		fmt.Println(
			"usage: termplayer file",
		)

		return
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