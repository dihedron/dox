package main

import (
	"bytes"
	"io"
	"log"
	"os"

	latex "github.com/dihedron/goldmark-latex"
	"github.com/jessevdk/go-flags"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

func main() {

	var opts struct {
		Input    string  `short:"i" long:"input" description:"The name of the input Markdown file" value-name:"INPUT" default:"_test/test0.md"`
		Output   string  `short:"o" long:"output" description:"The name of the output LaTeX file" value-name:"OUTPUT" default:"_test/test0.tex"`
		Preamble *string `short:"p" long:"preamble" description:"The path to the preamble file" value-name:"PREAMBLE" default:"_test/preamble.tex"`
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

	var preamble []byte
	if opts.Preamble != nil {
		var p *os.File
		if p, err = os.Open(*opts.Preamble); err != nil {
			log.Fatalf("error opening preamble file: %v", err)
		}
		defer p.Close()

		preamble, err = io.ReadAll(p)
		if err != nil {
			log.Fatalf("error reading preamble file: %v", err)
		}
	}

	rd := renderer.NewRenderer(
		renderer.WithNodeRenderers(
			util.Prioritized(
				latex.NewRenderer(
					latex.Config{
						NoHeadingNumbering: false,
						Unsafe:             true,
						Preamble:           preamble,
						HeadingLevelOffset: 0,
					},
				),
				1000),
		),
	)

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRenderer(rd),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert(data, &buf); err != nil {
		log.Fatalf("error converting input Markdown: %v", err)
	}

	_, err = output.Write(buf.Bytes())
	if err != nil {
		log.Printf("error writing to output: %v", err)
	}
}
