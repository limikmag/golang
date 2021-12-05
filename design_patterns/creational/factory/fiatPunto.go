package main

type fiatPunto struct {
	fiat
}

func newFiatPunto() iFiat {
	return &fiatPunto{
		fiat: fiat{
			name:   "Fiat Punto",
			speed:  220,
			weight: 1300,
		},
	}
}
