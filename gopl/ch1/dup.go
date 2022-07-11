package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	count := make(map[string]map[string]int64)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, count)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, count)
			f.Close()
		}
	}

	for fileName, lineMap := range count {
		for line, n := range lineMap {
			if n > 1 {
				fmt.Printf("file: %s\t line: %s\t count: %d\n", fileName, line, n)
			}
		}

	}
}

func countLines(f *os.File, count map[string]map[string]int64) {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if count[f.Name()] == nil{
			count[f.Name()] = make(map[string]int64)
		}
		count[f.Name()][scanner.Text()]++
	}
}
