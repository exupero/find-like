package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	lev "github.com/exupero/levenshtein"
)

const MAX_INDEX = int(^uint(0) >> 1) // Largest possible int
const MAX_DISTANCE = int(^uint(0) >> 1) // Largest possible int

type Edge struct {
	From, To, Weight int
}

func sort(items lev.Sequence) []Edge {
	results := []Edge{}

	for i := 0; i < items.Length(); i++ {
		best := MAX_INDEX
		bestDist := MAX_DISTANCE

		for j := 0; j < items.Length(); j++ {
			if i == j { continue }

			dist := lev.Distance(items.Pair(i, j))
			if dist < bestDist {
				best, bestDist = j, dist
			}
		}

		results = append(results, Edge{i, best, bestDist})
	}

	return results
}

func report(lines []string, edges []Edge, showLines bool) {
	if showLines {
		for _, e := range edges {
			fmt.Println(lines[e.From])
			fmt.Println(lines[e.To])
			fmt.Println(e.Weight)
			fmt.Println()
		}
	} else {
		for _, e := range edges {
			fmt.Printf("%d %d %d\n", e.From, e.To, e.Weight)
		}
	}
}

func main() {
	byWord := flag.Bool("w", false, "Sort by word chunks")
	showLines := flag.Bool("l", false, "Show lines instead of line numbers")

	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	var edges []Edge

	if *byWord {
		var rawLines []string
		var lines lev.ByWord
		for scanner.Scan() {
			line := scanner.Text()
			rawLines = append(rawLines, line)
			lines = append(lines, strings.Split(line, " "))
		}
		edges = sort(lines)
		report(rawLines, edges, *showLines)
	} else {
		var lines lev.ByLetter
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		edges = sort(lines)
		report(lines, edges, *showLines)
	}
}
