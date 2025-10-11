package main

import "fmt"

type vkorb int

var x vkorb

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
