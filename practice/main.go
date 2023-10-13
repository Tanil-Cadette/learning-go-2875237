package main

import (
	"fmt"
)

func main() {
	doSomething()
}

func doSomething() {
	fmt.Println("Doing Something")
	sum := addValues(10, 34)
	fmt.Println("The sum is:", sum)

	multiSum, multiCount := addAllValues(1, 5, 43, 4, -9)
	fmt.Println("The sum of multiple values:", multiSum)
	fmt.Println("The count of multiple count:", multiCount)

}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
