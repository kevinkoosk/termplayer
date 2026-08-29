package player

import (
	"os"

	"golang.org/x/term"
)

type Command int

const (
	CmdNone Command = iota
	CmdQuit
	CmdPause
)

func StartKeyboard() (chan Command, func(), error) {

	oldState, err :=
		term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		return nil, nil, err
	}

	restore := func() {
		_ = term.Restore(
			int(os.Stdin.Fd()),
			oldState,
		)
	}

	cmdChan := make(chan Command)

	go func() {

		var buf [1]byte

		for {

			_, err :=
				os.Stdin.Read(buf[:])

			if err != nil {
				return
			}

			switch buf[0] {

			case 'q', 'Q':
				cmdChan <- CmdQuit

			case ' ':
				cmdChan <- CmdPause
			}
		}
	}()

	return cmdChan, restore, nil
}