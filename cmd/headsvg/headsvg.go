package main

import (
	"fmt"
	"os"

	"github.com/PotatoMaster101/headsvg/pkg/head"
	"github.com/PotatoMaster101/headsvg/pkg/svg"
	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("headsvg", "Converts Minecraft player head to SVG")
	name := parser.String("n", "name", &argparse.Options{Required: true, Help: "Player username"})
	scale := parser.Int("s", "scale", &argparse.Options{Default: 100, Help: "SVG pixel scale"})
	percent := parser.Flag("p", "percent", &argparse.Options{Default: false, Help: "Whether to output x/y/width/height as a percentage"})
	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	img, err := head.GetHeadFromNet(*name, true)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	svg.WriteSvg(img, os.Stdout, *scale, *percent)
}
