package player

import (
	"fmt"
	"time"
	"errors"
	"github.com/kevinkoosk/termplayer/internal/config"
	"github.com/kevinkoosk/termplayer/internal/decoder"
	"github.com/kevinkoosk/termplayer/internal/renderer"
	"github.com/kevinkoosk/termplayer/internal/terminal"
)

var ErrQuit =
	errors.New("quit requested")

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
				return ErrQuit

			case CmdPause:
				paused = !paused

			case CmdZoomIn:
			    config.ScalePercent += 10

			    if config.ScalePercent > 150 {
			        config.ScalePercent = 150
			    }

			case CmdZoomOut:
			    config.ScalePercent -= 10

			    if config.ScalePercent < 25 {
			        config.ScalePercent = 25
			    }

			case CmdZoomReset:

			    config.ScalePercent = 100

			case CmdHud:

			    config.HudMode++

			    if config.HudMode > 2 {
			        config.HudMode = 0
			    }
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

		switch config.HudMode {

		case 0:

		    fmt.Printf(
		        "\x1b[-2;1HQ Quit | Space Pause | +/- Zoom | 0 Reset | Tab HUD  ",
		    )

		case 1:

		    fmt.Printf(
		        "\x1b[-2;1HFPS: %.1f | Drop: %d | Scale: %d%%                  ",
		        fps,
		        framesDropped,
		        config.ScalePercent,
		    )

		case 2:

		    fmt.Printf(
		        "                                                               ",
		    )
		    // no HUD
		}

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