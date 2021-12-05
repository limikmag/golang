package main

type fiat125p struct {
	fiat
}

func newFiat125p() iFiat {
	return &fiat125p{
		fiat: fiat{
			name:   "Fiat 125p",
			speed:  140,
			weight: 900,
		},
	}
}
