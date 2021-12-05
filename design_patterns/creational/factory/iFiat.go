package main

type iFiat interface {
	setName(name string)
	setSpeed(speed int)
	setWeight(weight int)
	getName() string
	getSpeed() int
	getWeight() int
}
