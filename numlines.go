package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

var lines []string

func scan_stdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
}

func add_lines_of_file() {
	for _, filename := range os.Args[1:] {
		dat, err := os.ReadFile(filename)
		if err == nil {
			lines_in_file :=strings.Split(string(dat), "\n")
			for _, line := range lines_in_file {
				lines = append(lines, line)
			}
		}			
	}
}

func output()  {
	formator := len(fmt.Sprint(len(lines)))
	formatstr := "%" + fmt.Sprint(formator) + "d\t%s\n"
	i := 1
	for _, line := range lines {
		fmt.Printf(formatstr, i, line)
		i++
	}
}

func main() {
	scan_stdin()
	add_lines_of_file()
	output()
}
