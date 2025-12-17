package main

import (
	"fmt"
)

func dummy() {
	var x int = 10
	fmt.Println("Hello", x)
	fmt.Printf("Hello %v\n", x)
	fmt.Printf("Type of %v is %T\n", x, x)
	// print memory address of x
	fmt.Printf("Memory address of %v is %v\n", x, &x)
	// read input from user
	var y int
	fmt.Println("Enter a number")
	fmt.Scanf("%d", &y)
	fmt.Printf("You entered %v\n", y)

	// decalared array
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[0])
	// arr[5] = 6

	// empty array or just declare and dont assign values
	var arr1 [5]int
	arr1[0] = 1
	fmt.Println(arr1[0])

	// slice
	var slice []int
	slice = append(slice, 1)
	fmt.Println(slice[0])
	fmt.Printf("Length of slice is %v\n", len(slice))

	// for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("Infinite loop")
		if true {
			break
		}

	}

	for _, v := range arr {
		fmt.Println(v)
	}

	boolean := y > 10
	if boolean {
		fmt.Println("Y is greater than 10")
	} else {
		fmt.Println("Y is less than 10")
	}

}

func some(x int, y int) (sum int) {
	sum = x + y
	return sum
}

func main() {
	if x := some(1, 20); x > 3 {
		fmt.Println("X is greater than 3")
	}
}
