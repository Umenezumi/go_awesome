package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args[0])

	for i, arg := range os.Args {
		fmt.Println("index:", i)
		fmt.Println("arg:", arg)
	}
}
