package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/exupero/levenshtein"
)

func best(line string, lines []string) string {
	bestLine, bestScore := "", int(^uint(0) >> 1) // largest int
	for _, testLine := range lines {
		score := levenshtein.LevenshteinDistance(line, testLine)
		if score < bestScore {
			bestLine, bestScore = testLine, score
		}
	}
	return bestLine
}

func main() {
	flag.Parse()

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(best(flag.Arg(0), lines))
}
