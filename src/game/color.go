package game

import (
	"image/color"
	"strconv"
)

// ParseHexColor converts a hex color string (like "#FF0000") to color.RGBA
func ParseHexColor(hex string) (color.RGBA, error) {
	// Remove '#' if present
	if hex[0] == '#' {
		hex = hex[1:]
	}

	// Parse the hex string
	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}

	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}

	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}

	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}, nil
}
