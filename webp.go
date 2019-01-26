package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/chai2010/webp"
	"github.com/spf13/pflag"
)

const (
	ExitDecodeError = 64
	ExitEncodeError = 65
)

func main() {
	options := new(webp.Options)
	pflag.BoolVar(&options.Lossless, "lossless", false,
		"encode in lossless way")
	pflag.Float32VarP(&options.Quality, "quality", "q", 95,
		"quality values in [0, 100]")
	pflag.Parse()

	im, format, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(ExitDecodeError)
	}
	_ = format

	err = webp.Encode(os.Stdout, im, options)
	if err != nil {
		os.Exit(ExitEncodeError)
	}

	os.Exit(0)
}
