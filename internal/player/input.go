package player

import (
	"bufio"
	"os"
	"strings"
)

func ListenForQuit() chan struct{} {

	quit := make(chan struct{})

	go func() {

		reader := bufio.NewReader(os.Stdin)

		for {

			text, err := reader.ReadString('\n')

			if err != nil {
				return
			}

			text = strings.TrimSpace(text)

			if text == "q" || text == "Q" {
				close(quit)
				return
			}
		}
	}()

	return quit
}