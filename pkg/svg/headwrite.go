package svg

import (
	"fmt"
	"io"

	"github.com/PotatoMaster101/headsvg/pkg/head"
	"github.com/go-playground/colors"
)

// WriteSvg writes the given `playerHead` to SVG format, save to `output`.
func WriteSvg(playerHead head.PlayerHead, output io.Writer, scale int, percent bool) {
	dim := head.Dimension * scale
	output.Write([]byte("<?xml version=\"1.0\"?>\n"))
	output.Write([]byte(fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"%d\" height=\"%d\">\n", dim, dim)))
	for y := 0; y < head.Dimension; y++ {
		for x := 0; x < head.Dimension; x++ {
			colour := fmt.Sprintf("fill:%s", colors.FromStdColor(playerHead[x][y]).ToHEX().String())
			xScale, yScale := x * scale, y * scale

			if percent {
				xPercent, yPercent := float64(xScale) / float64(dim) * 100.0, float64(yScale) / float64(dim) * 100.0
				scalePercent := float64(scale) / float64(dim) * 100.0
				output.Write([]byte(fmt.Sprintf("<rect x=\"%.1f%%\" y=\"%.1f%%\" width=\"%.1f%%\" height=\"%.1f%%\" style=\"%s\" />\n", xPercent, yPercent, scalePercent, scalePercent, colour)))
			} else {
				output.Write([]byte(fmt.Sprintf("<rect x=\"%d\" y=\"%d\" width=\"%d\" height=\"%d\" style=\"%s\" />\n", xScale, yScale, scale, scale, colour)))
			}
		}
	}
	output.Write([]byte("</svg>"))
}
