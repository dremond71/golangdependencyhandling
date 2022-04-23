package main

import (
	"fmt"

	"github.com/dremond71/golangdependencyexampleone"
	"github.com/dremond71/golangdependencyexampletwo"
)

func main() {

	someString := golangdependencyexampletwo.GetSomeString()
	fmt.Println(someString, "should look like", "OBIWAN KENOBI")

	someString2 := "Anakin Skywalker"
	someString2Modified := golangdependencyexampleone.DoSomething(someString2)
	fmt.Println(someString2Modified, "should look like", "ANAKIN SKYWALKER")

}
