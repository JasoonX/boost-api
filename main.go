package main

import (
	"os"

	"github.com/BOOST-2021/boost-app-back/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(2)
	}
	os.Exit(0)
}
