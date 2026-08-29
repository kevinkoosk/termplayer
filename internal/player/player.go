package player

import (
	"fmt"
	"time"

	"github.com/kevinkoosk/termplayer/internal/decoder"
	"github.com/kevinkoosk/termplayer/internal/renderer"
	"github.com/kevinkoosk/termplayer/internal/terminal"
)

func Play(video string) error {

	fmt.Print("\x1b[?25l")
	defer fmt.Print("\x1b[?25h")

	cmdChan, restore, err :=
		StartKeyboard()

	if err != nil {
		return err
	}

	defer restore()

	stream, err := decoder.New(
		video,
		160,
		90,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer stream.Close()

	framesRendered := 0
	framesDropped := 0
	fps := 0.0
	fpsTimer := time.Now()
	paused := false

	frameDuration := time.Second / 10

	fmt.Print("\x1b[2J")

	for {

		select {
	
		case cmd := <-cmdChan:

			switch cmd {

			case CmdQuit:
				return nil

			case CmdPause:
				paused = !paused
			}
		default:
		}

		start := time.Now()

		if paused {

			time.Sleep(
				50 * time.Millisecond,
			)

			continue
		}

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
			"\x1b[-2;1HFPS: %.1f | Drop: %d | Q Quit | Space Pause   ",
			fps,
			framesDropped,
		)

		elapsed := time.Since(start)

		if elapsed > frameDuration {

			framesDropped++

			continue
		}

		if elapsed < frameDuration {

			time.Sleep(
				frameDuration - elapsed,
			)
		}
	}

	return nil
}