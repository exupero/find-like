package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	lev "github.com/exupero/levenshtein"
)

func main() {
	byWord := flag.Bool("w", false, "Sort by word chunks")

	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	if *byWord {
		var lines lev.ByWord
		for scanner.Scan() {
			lines = append(lines, strings.Split(scanner.Text(), " "))
		}
		lev.Sort(lines)
		for _, line := range lines {
			fmt.Println(strings.Join(line, " "))
		}
	} else {
		var lines lev.ByLetter
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		lev.Sort(lines)
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}
