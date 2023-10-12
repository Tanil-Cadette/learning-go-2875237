package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().Unix()
	rng := rand.New(rand.NewSource(seed))
	// Generate a random integer representing the day of the week (1 to 7)
	// dow := rng.Intn(7) + 1
	// rand.Seed(time.Now().Unix())
	// dow := rand.Intn(7) + 1
	// fmt.Println("Day", dow)
	var result string
	switch dow := rng.Intn(7) + 1; dow {
	case 1:
		result = "It's Sunday!"
		fmt.Println("Day", dow)
	case 2:
		result = "It's Monday!"
		fmt.Println("Day", dow)
	case 3:
		result = "It's Tuesday!"
		fmt.Println("Day", dow)
	case 4:
		result = "It's Thursday!"
		fmt.Println("Day", dow)
	default:
		result = "It's some other day!"
		fmt.Println("Day", dow)
	}
	fmt.Println(result)
}
