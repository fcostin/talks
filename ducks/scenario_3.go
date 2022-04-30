package main

import (
	"fmt"
)

// -----

type Counter interface {
	Clear()     //	mutation, idempotent
	Inc(n int)  //	mutation, not idempotent
	Value() int //	read only
}

// -----

type SimpleCounter int

func NewSimpleCounter() Counter {
	c := new(SimpleCounter)
	c.Clear()
	return c
}

func (c *SimpleCounter) Clear() {
	*c = SimpleCounter(0)
}

func (c *SimpleCounter) Inc(n int) {
	*c = (SimpleCounter)(n + int(*c))
}

func (c *SimpleCounter) Value() int {
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
	c := NewSimpleCounter()
	c.Inc(3)
	c.Inc(2)

	actual := c.Value()
	expected := 5

	checkEqual(expected, actual)
}

