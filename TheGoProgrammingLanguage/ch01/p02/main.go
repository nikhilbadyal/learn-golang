package main

import (
	"fmt"
	"os"
)

func main() {
	for k, v := range os.Args[1:] {
		fmt.Printf("Index is %v and value is %v\n", k, v)
	}
}
