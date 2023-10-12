package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	poodle.Weight = 9
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
}

type Dog struct {
	Breed  string
	Weight int
}
