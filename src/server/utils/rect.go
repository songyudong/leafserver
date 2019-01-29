package utils

import "math"

type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func IsIntersect(r1 *Rect, r2 *Rect) bool {

	w := math.Abs((r1.X+r1.X+r1.Width)/2 - (r2.X+r2.X+r2.Width)/2)
	ww := ((r1.X + r1.Width + r2.X + r2.Width - r1.X - r2.X) / 2)

	h := math.Abs((r1.Y+r1.Y+r1.Height)/2 - (r2.Y+r2.Y+r2.Height)/2)
	hh := ((r1.Y + r1.Height + r2.Y + r2.Height - r1.Y - r2.Y) / 2)
	if w < ww && h < hh {
		return true
	}

	return false
}
