package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, count)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				panic("Unable to open file.\n")
			}
			countLines(f, count)
			f.Close()
		}
	}

	for line, n := range count {
		if n > 1 {
			fmt.Printf("%s \t %d\n", line, n)
		}
	}

}

func countLines(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() && input.Text() != "q" {
		count[input.Text()]++
	}
}
