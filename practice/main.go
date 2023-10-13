package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10, "woof!"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
	poodle.Speak()
	poodle.Sound = "Arf!!"
	poodle.Speak()
	poodle.speakThreeTimes()
	poodle.speakThreeTimes()
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Speak is how the dog speaks
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog) speakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v\n", d.Sound, d.Sound, d.Sound)
	fmt.Print(d.Sound)
}
