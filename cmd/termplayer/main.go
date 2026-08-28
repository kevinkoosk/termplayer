package main

import (
	"fmt"
	"os"

	"github.com/kevinkoosk/termplayer/internal/player"
)

func main() {

	if len(os.Args) < 2 {
	
		fmt.Println(
			"usage: termplayer video.mp4",
		)
		return
	}
	
	err :=
		player.Play(os.Args[1])
	
	if err != nil {
	
		panic(err)
	}
}