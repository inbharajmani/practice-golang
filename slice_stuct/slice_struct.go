package main

import "fmt"

func main() {
	type Person struct {
		Name string
		Age  int
	}

	p := []Person {
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	for _, person := range p {
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	}

	i := Person{"David", 40}
	fmt.Printf("Name: %s, Age: %d\n", i.Name, i.Age)
}