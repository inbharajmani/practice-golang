package main

import "fmt"

type axis struct {
	X int
	Y int
}

type player struct {
	Name string
	axis
	Keys []string
}

func main() {
	i := axis{
		X: 10,
		Y: 1,
	}
	fmt.Println(i)
	i.Move(1, 2)
	fmt.Println(i)

	p1 := player{
		Name: "Inbharaj",
		Keys: []string{"Copper"},
	}

	fmt.Printf("%+v\n", p1)
	p1.Move(1, 1)
	fmt.Printf("%+v\n", p1)

	p1.Found()

}

func (i *axis) Move(x, y int) {
	i.X += x
	i.Y += y
}

func (k player) Found(i string) string {
	k.
}