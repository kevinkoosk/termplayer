package terminal

import "golang.org/x/term"

func Size() (int, int) {

	w, h, err :=
		term.GetSize(0)

	if err != nil {

		return 80, 24
	}

	return w, h
}