package main

import (
	"fmt"

	"github.com/akarshippili/golang/designpatterns/factory/factories"
)

type Zoo struct {
	factory factories.AnimalFactory
}

func (zoo *Zoo) visitAnimals() {
	for i := 0; i < 10; i++ {

		animal := zoo.factory.GetAnimal()
		animal.Eat()
		animal.Run()
		animal.Talk()

		fmt.Println()
		fmt.Println()
	}
}

func main() {
	zoo := Zoo{
		factory: &factories.RandomFactory{},
	}
	zoo.visitAnimals()
	fmt.Println("---------------------------------------")

	zoo.factory = &factories.IndianFactory{}
	zoo.visitAnimals()
	fmt.Println("---------------------------------------")

	zoo.factory = &factories.AfricanFactory{}
	zoo.visitAnimals()
}
