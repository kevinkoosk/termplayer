package decoder

func BestFit(
	videoW,
	videoH,
	maxW,
	maxH int,
) (int, int) {

	ratio :=
		float64(videoW) /
			float64(videoH)

	w := maxW

	h :=
		int(
			float64(w) /
				ratio,
		)

	if h > maxH {

		h = maxH

		w =
			int(
				float64(h) *
					ratio,
			)
	}

	return w, h
}