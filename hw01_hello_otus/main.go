package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	hello := reverse.String("Hello, OTUS!")
	fmt.Println(hello)
}
