package player

import (
	"fmt"
	"time"

	"github.com/kevinkoosk/termplayer/internal/decoder"
	"github.com/kevinkoosk/termplayer/internal/renderer"
	"github.com/kevinkoosk/termplayer/internal/terminal"
)

func Play(video string) error {

	quit := ListenForQuit()

	stream, err := decoder.New(video)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer stream.Close()

	framesRendered := 0
	fps := 0.0
	fpsTimer := time.Now()

	frameDuration := time.Second / 10

	fmt.Print("\x1b[2J")

	for {

		select {

		case <-quit:
			return nil

		default:
		}

		start := time.Now()

		frame, err := stream.NextFrame()

		if err != nil {
			break
		}

		w, h := terminal.Size()

		// Leave first row for status text
		renderer.Draw(
			frame,
			w,
			h-2,
		)

		framesRendered++

		elapsedFPS := time.Since(fpsTimer)

		if elapsedFPS >= time.Second {

			fps =
				float64(framesRendered) /
					elapsedFPS.Seconds()

			framesRendered = 0
			fpsTimer = time.Now()
		}

		// Draw FPS on top row
		fmt.Printf(
			"\x1b[-2;1HFPS: %.1f   (Q + Enter to quit)",
			fps,
		)

		elapsed := time.Since(start)

		if elapsed < frameDuration {

			time.Sleep(
				frameDuration - elapsed,
			)
		}
	}

	return nil
}