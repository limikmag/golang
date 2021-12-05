package main

type fiat struct {
	name   string
	speed  int
	weight int
}

func (g *fiat) setName(name string) {
	g.name = name
}

func (g *fiat) getName() string {
	return g.name
}

func (g *fiat) setSpeed(speed int) {
	g.speed = speed
}

func (g *fiat) getSpeed() int {
	return g.speed
}

func (g *fiat) setWeight(weight int) {
	g.weight = weight
}

func (g *fiat) getWeight() int {
	return g.weight
}
