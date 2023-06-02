package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func main() {

	var opts struct {
		Input  string `short:"i" long:"input" description:"The name of the input Markdown file" value-name:"INPUT"`
		Output string `short:"o" long:"output" description:"The name of the output LaTeX file" value-name:"OUTPUT"`
		// Theme  string `short:"t" long:"theme" description:"The name of the theme file" value-name:"THEME"`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatalf("error parsing command line: %v", err)
	}

	log.Printf("args: INPUT: %q, OUTPUT: %q\n", opts.Input, opts.Output)

	input := os.Stdin
	if opts.Input != "" {
		if input, err = os.Open(opts.Input); err != nil {
			log.Fatalf("error opening input file: %v", err)
		}
	}
	if input != os.Stdin {
		defer input.Close()
	}

	output := os.Stdout
	if opts.Output != "" {
		if output, err = os.Create(opts.Output); err != nil {
			log.Fatalf("error opening output file: %v", err)
		}
	}
	if output != os.Stdout {
		defer output.Close()
	}

	data, err := io.ReadAll(input)
	if err != nil {
		log.Fatalf("error reading data from input: %v", err)
	}

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	var buf bytes.Buffer
	if err := md.Convert(data, &buf); err != nil {
		log.Fatalf("error converting input Markdown: %v", err)
	}

	fmt.Printf("%s\n", buf.String())
}
