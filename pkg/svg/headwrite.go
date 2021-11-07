package svg

import (
	"fmt"
	"io"

	"github.com/PotatoMaster101/headsvg/pkg/head"
	"github.com/ajstarks/svgo"
	"github.com/go-playground/colors"
)

// WriteSvg writes the given `playerHead` to SVG format, save to `output`.
func WriteSvg(playerHead head.PlayerHead, output io.Writer, scale int) {
	canvas := svg.New(output)
	canvas.Start(head.Dimension * scale, head.Dimension * scale)
	for y := 0; y < head.Dimension; y++ {
		for x := 0; x < head.Dimension; x++ {
			color := colors.FromStdColor(playerHead[x][y]).ToHEX().String()
			canvas.Rect(x * scale, y * scale, scale, scale, fmt.Sprintf("fill:%s", color))
		}
	}
	canvas.End()
}
