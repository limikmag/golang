package main

import (
	"errors"
	"fmt"
)

// factory design pattern is a creational pattern and it is also one of the most commonly used pattern.
// this pattern provides a way to hide creation logic of the instances being created
// client or programmer only interacts with factory struct and tells the kind of instances that needs to be created

// Below example:
// We have Fiat manufacture factory which can product different kinds of cars

/* A factory is simply a wrapper function around a constructor (possibly one in a different class).
   The key difference is that a factory method pattern requires the entire object to be built in a single method call,
   with all the parameters passed in on a single line. The final object will be returned.
   A builder pattern, on the other hand, is in essence a wrapper object around all the possible parameters you might want to pass into a constructor invocation.
   This allows you to use setter methods to slowly build up your parameter list.
   One additional method on a builder class is a build() method, which simply passes the builder object into the desired constructor, and returns the result.

   Basic Factory Example:

   // Factory
	static class FruitFactory {
		static Fruit create(name, color, firmness) {
			// Additional logic
			return new Fruit(name, color, firmness);
		}
	}

   // Usage
	Fruit fruit = FruitFactory.create("apple", "red", "crunchy");

	Basic Builder Example:

	// Builder
	class FruitBuilder {
		String name, color, firmness;
		FruitBuilder setName(name)         { this.name     = name;     return this; }
		FruitBuilder setColor(color)       { this.color    = color;    return this; }
		FruitBuilder setFirmness(firmness) { this.firmness = firmness; return this; }
		Fruit build() {
			return new Fruit(this); // Pass in the builder
		}
	}

    // Usage
	Fruit fruit = new FruitBuilder()
			.setName("apple")
			.setColor("red")
			.setFirmness("crunchy")
			.build();



	Analogy:

    Factory: Consider a restaurant. The creation of "today's meal" is a factory pattern, because you tell the kitchen "get me today's meal"
	and the kitchen (factory) decides what object to generate, based on hidden criteria.
    Builder: The builder appears if you order a custom pizza. In this case, the waiter tells the chef (builder)
	"I need a pizza; add cheese, onions and bacon to it!"
	Thus, the builder exposes the attributes the generated object should have, but hides how to set them.

*/

func GetFiatCar(fiatCarType string) (iFiat, error) {
	if fiatCarType == "Fiat 126p" {
		return newFiat126p(), nil
	}

	if fiatCarType == "Fiat 125p" {
		return newFiat125p(), nil
	}

	if fiatCarType == "Fiat Panda" {
		return newFiatPanda(), nil
	}

	if fiatCarType == "Fiat Punto" {
		return newFiatPunto(), nil
	}
	return nil, errors.New(fmt.Sprintf("We does not product car with name %v", fiatCarType))
}

func main() {
	fiat125p, _ := GetFiatCar("Fiat 125p")
	fiat126p, _ := GetFiatCar("Fiat 126p")
	fiatPanda, _ := GetFiatCar("Fiat Panda")
	fiatPunto, _ := GetFiatCar("Fiat Punto")

	printDetails(fiat125p)
	printDetails(fiat126p)
	printDetails(fiatPanda)
	printDetails(fiatPunto)

}

func printDetails(g iFiat) {
	fmt.Printf("Fiat name: %s", g.getName())
	fmt.Println()
	fmt.Printf("Speed: %d", g.getSpeed())
	fmt.Println()
	fmt.Printf("Weight: %d", g.getWeight())
	fmt.Println()
}
