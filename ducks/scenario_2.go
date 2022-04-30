package main

import (
	"fmt"
	"math/rand"
)

// -----

type Thing interface {
	ThingKind() string
}

type ThingCounter interface {
	AcceptThings(things []Thing)
	GetTotalThings() int
}

type Colour int

const (
	Brown Colour = iota
	GreyGreen
	ColourSentinel
)

type Duck struct {
	Colour Colour
}

func (x Duck) ThingKind() string {
	return "Duck"
}

type NonDuck struct {
}

func (x NonDuck) ThingKind() string {
	return "not a duck"
}

// -----

type SimpleThingCounter int

func NewSimpleThingCounter(initialThings []Thing) ThingCounter {
	c := new(SimpleThingCounter)
	*c = SimpleThingCounter(0)
	c.AcceptThings(initialThings)
	return c
}

func (c *SimpleThingCounter) AcceptThings(things []Thing) {
	*c = (SimpleThingCounter)(len(things) + int(*c))
}

func (c *SimpleThingCounter) GetTotalThings() int {
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

func NewRandomThing() Thing {
	if rand.Intn(2) == 0 {
		return NonDuck{}
	} else {
		return Duck{Colour(rand.Intn(int(ColourSentinel)))}
	}
}

func main() {
	c := NewSimpleThingCounter([]Thing{NewRandomThing(), NewRandomThing(), NewRandomThing()})
	c.AcceptThings([]Thing{NewRandomThing(), NewRandomThing()})

	actual := c.GetTotalThings()
	expected := 5

	checkEqual(expected, actual)
}

