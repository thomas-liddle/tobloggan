package main

import (
	"flag"
	"log"
	"os"

	"github.com/mdwhatcott/tobloggan/code/tobloggan"
)

func main() {
	var sourceDirectory string
	flags := flag.NewFlagSet("tobloggan", flag.ExitOnError)
	flags.StringVar(&sourceDirectory, "src", "", "The directory containing blog source files (*.md).")
	_ = flags.Parse(os.Args[1:])
	config := tobloggan.Config{
		SourceDirectory: sourceDirectory,
		Logger:          log.New(os.Stderr, "", log.Ltime),
	}
	ok := tobloggan.GenerateBlog(config)
	if !ok {
		os.Exit(1)
	}
}
