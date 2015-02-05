package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/exupero/levenshtein"
)

var scores = map[string]int{}

func best(scores map[string]int) string {
	bestPath, bestScore := "", int(^uint(0) >> 1) // largest int
	for path, score := range scores {
		if score < bestScore {
			bestPath, bestScore = path, score
		}
	}
	return bestPath
}

func scoreFilename(filename string) func(string, os.FileInfo, error) error {
	return func(path string, f os.FileInfo, err error) error {
		if !f.Mode().IsDir() {
			scores[path] = levenshtein.LevenshteinDistance(filename, path)
		}
		return nil
	}
}

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Fprintf(os.Stderr, "Too few arguments; expected at least one directory to search\n")
		return
	}

	filename := flag.Arg(0)
	for _, dir := range flag.Args()[1:] {
		err := filepath.Walk(dir, scoreFilename(filename))
		if err != nil {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("%s", err))
		}
	}

	fmt.Println(best(scores))
}
