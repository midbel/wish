package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	for _, f := range flag.Args() {
		err := catFile(f)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func catFile(file string) error {
	r, err := os.Open(file)
	if err != nil {
		return err
	}
	defer r.Close()

	_, err = io.Copy(os.Stdout, r)
	return err
}
