package main

type fiat126p struct {
	fiat
}

func newFiat126p() iFiat {
	return &fiat126p{
		fiat: fiat{
			name:   "Fiat 126p",
			speed:  190,
			weight: 1500,
		},
	}
}
