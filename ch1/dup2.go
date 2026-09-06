package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileCounts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		processReader(os.Stdin, "No files", fileCounts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			processReader(f, arg, fileCounts)
			f.Close()
		}
	}
	for file, value := range fileCounts {
		for line, n := range value {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, file)
			}
		}

	}
}

func processReader(f *os.File, fileName string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	if counts[fileName] == nil {
		counts[fileName] = make(map[string]int) // Создаем внутреннюю карту
	}
	for input.Scan() {
		counts[fileName][input.Text()]++
	}
}
