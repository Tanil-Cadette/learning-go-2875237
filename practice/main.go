package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5)
	numbers[0] = 134
	numbers[1] = 2
	numbers[2] = 35
	numbers[3] = 459
	numbers[4] = 45
	fmt.Println(numbers)
	numbers = append(numbers, 356)
	fmt.Println(numbers)
	sort.Ints(numbers)
	fmt.Println(numbers)

	var numbers2 = []int{1, 2, 3, 4, 5}
	fmt.Println(numbers2)
	numbers2 = append(numbers2, 6)
	fmt.Println(numbers2)

}
