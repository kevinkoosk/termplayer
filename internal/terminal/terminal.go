package terminal

import "golang.org/x/term"
import "os"

func Size() (int, int) {

	w, h, err :=
		term.GetSize(int(os.Stdout.Fd()))

	if err != nil {

		return 80, 24
	}

	return w, h
}