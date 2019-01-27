package main

import (
	"os"
)

const (
	Name    = "gossh"
	Version = "1.0"
)

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args[1:]))
}
