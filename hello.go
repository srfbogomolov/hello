package main

import (
	"fmt"

	"github.com/srfbogomolov/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
