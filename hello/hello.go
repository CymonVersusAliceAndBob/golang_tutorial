package main

import (
	"fmt"

	"github.com/CymonVersusAliceAndBob/golang_tutorial/greetings"
)

func main() {
	message := greetings.Hello("Malcolm")
	fmt.Println(message)
}
