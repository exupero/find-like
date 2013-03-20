package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"github.com/exupero/levenshtein"
)

var filename = ""
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

func scoreFilename(path string, f os.FileInfo, err error) error {
	if !f.Mode().IsDir() {
		scores[path] = levenshtein.LevenshteinDistance(filename, path[5:]) // trim 'spec/'
	}
	return nil
}

func main() {
	flag.Parse()
	filename = flag.Arg(0)
	err := filepath.Walk("spec", scoreFilename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(best(scores))
}
