package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()

	var total int
	for _, f := range flag.Args() {
		c, err := countFile(f)
		if err != nil {
			continue
		}
		total += c
		fmt.Printf("%d %s", c, f)
		fmt.Println()
	}
	fmt.Printf("%d lines", total)
	fmt.Println()
}

func countDir(dir string) (int, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, err
	}
	var total int
	for _, e := range entries {
		c, err := countFile(filepath.Join(dir, e.Name()))
		if err != nil {
			continue
		}
		total += c
	}
	return total, nil
}

func countFile(file string) (int, error) {
	stat, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	if stat.IsDir() {
		return countDir(file)
	}
	r, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer r.Close()

	var (
		scan = bufio.NewScanner(r)
		line int
	)
	for scan.Scan() {
		line++
	}
	return line, scan.Err()
}
