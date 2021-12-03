package head

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/PotatoMaster101/headsvg/pkg/skin"
)

const (
	// Dimension is the face area dimension.
	Dimension = 8
	// FrontXY is the X and Y coordinates of the face area.
	FrontXY = 8
	// HatX is the X coordinate of the hat area.
	HatX = 40
	// HatY is the Y coordinate of the hat area.
	HatY = 8
)

// PlayerHead represents a 8x8 array of `Color`.
type PlayerHead = [Dimension][Dimension]color.Color

// GetHeadFromNet retrieves the head portion from `username`'s skin. `hat` determines whether the hat portion should be included.
func GetHeadFromNet(username string, hat bool) (PlayerHead, error) {
	url, _ := skin.GetSkinURL(username)
	resp, err := http.Get(url)
	if err != nil {
		return PlayerHead {}, err
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			// ignored
		}
	}(resp.Body)

	data, _ := ioutil.ReadAll(resp.Body)
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return PlayerHead {}, err
	}
	return getHeadPixels(img, hat), nil
}

// getHeadPixels returns the pixels for head portion (8x8) of the given `img`.
func getHeadPixels(img image.Image, hat bool) PlayerHead {
	pixels := PlayerHead {}
	for row := FrontXY; row < FrontXY + Dimension; row++ {
		for col := FrontXY; col < FrontXY + Dimension; col++ {
			pixels[row - FrontXY][col - FrontXY] = img.At(row, col)
		}
	}
	if !hat {
		return pixels
	}

	for row := HatX; row < HatX + Dimension; row++ {
		for col := HatY; col < HatY + Dimension; col++ {
			pix := img.At(row, col)
			if _, _, _, a := pix.RGBA(); a != 0 {   // add hat pixel only if not 100% transparent
				pixels[row - HatX][col - HatY] = pix
			}
		}
	}
	return pixels
}
