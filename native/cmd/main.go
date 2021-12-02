package main

import (
	"fmt"
)

type T = int;

func main() {
	var j T = 0;
	var value interface {} = j
	fmt.Println(value.(int))
	fmt.Println(value.(T))
}

