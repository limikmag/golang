package main

type fiatPanda struct {
	fiat
}

func newFiatPanda() iFiat {
	return &fiatPanda{
		fiat: fiat{
			name:   "Fiat Panda",
			speed:  190,
			weight: 1100,
		},
	}
}
