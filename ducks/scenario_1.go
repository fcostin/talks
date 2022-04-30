package main

import "fmt"

// -----

type Colour int

const (
	Brown Colour = iota
	GreyGreen
)

type Duck struct {
	Colour Colour
}

type DuckCounter interface {
	AcceptDucks(ducks []Duck)
	GetTotalDucks() int
}

// -----

type SimpleDuckCounter int

func NewSimpleDuckCounter(initialDucks []Duck) DuckCounter {
	c := new(SimpleDuckCounter)
	*c = SimpleDuckCounter(0)
	c.AcceptDucks(initialDucks)
	return c
}

func (c *SimpleDuckCounter) AcceptDucks(ducks []Duck) {
	*c = (SimpleDuckCounter)(len(ducks) + int(*c))
}

func (c *SimpleDuckCounter) GetTotalDucks() int {
	return int(*c)
}

// -----

func checkEqual(expected, actual int) {
	if expected != actual {
		fmt.Println("FAIL")
	} else {
		fmt.Println("PASS")
	}
}

func main() {
	c := NewSimpleDuckCounter([]Duck{Duck{Colour: Brown}, Duck{Colour: Brown}, Duck{Colour: Brown}})
	c.AcceptDucks([]Duck{Duck{Colour: GreyGreen}, Duck{Colour: GreyGreen}})

	actual := c.GetTotalDucks()
	expected := 5

	checkEqual(expected, actual)
}

