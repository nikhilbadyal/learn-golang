package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// args[0] hold the path of the file
	fmt.Println("Path is " + os.Args[0])
	fmt.Println("Arguments are " + strings.Join(os.Args[1:], "  "))
}
