package player
 
import (
	"fmt"
	"time"

	"github.com/kevinkoosk/termplayer/internal/decoder"
	"github.com/kevinkoosk/termplayer/internal/renderer"
	"github.com/kevinkoosk/termplayer/internal/terminal"
)
 
func Play(video string) error {
 
	stream, err := decoder.New(video)

	if err != nil {
		return err
	}

	defer stream.Close()

	fmt.Print("\x1b[2J")

	frameDuration :=
		time.Second / 12
 
	for {

		start := time.Now()

		frame, err :=
			stream.NextFrame()
 
		if err != nil {
			fmt.Println(err)
			break
		}
 
		w, h :=
			terminal.Size()

		renderer.Draw(
			frame,
			w,
			h-1,
		)

		elapsed :=
			time.Since(start)
 
		if elapsed < frameDuration {
 
			time.Sleep(
				frameDuration - elapsed,
			)
		}
	}

	return nil
}
