package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	help := flag.Bool("h", false, "Help")
	flag.Parse()
	if *help {
		usage(0)
	}

	args := flag.Args()
	if len(args) == 0 {
		usage(0)
	}

	for _, file := range args {
		width, height := GetSize(file)
		fmt.Printf("%s: %d / %d", file, width, height)
		fmt.Println()
	}
}

func usage(ec int) {
	doc := `Usage: imgsize v1.0
remtilde [OPTIONS] [image-file1 image-file2 ... image-fileX]
[OPTIONS]
 -h              Help (this page)
`

	s := getStream(ec)
	fmt.Fprint(s, doc)
	os.Exit(ec)
}

func getStream(exitCode int) io.Writer {
	if exitCode != 0 {
		return os.Stderr
	}

	return os.Stdout
}
