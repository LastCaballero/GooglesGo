package main

import (
	"bufio"
	"fmt"
	"os"
)

func do_stdin() {
	st_scanner := bufio.NewScanner(os.Stdin)
	for st_scanner.Scan() {
		fmt.Println(st_scanner.Text())
	}
}

func do_files() {
	for _, filename := range os.Args[1:] {
		bytes, err := os.ReadFile(filename)
		if err == nil {
			fmt.Printf(string(bytes))
		}
	}
}

func main() {
	do_stdin()
	do_files()
}
