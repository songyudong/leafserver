package utils

import "math"

type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func IsIntersect(r1 *Rect, r2 *Rect) bool {

	w := math.Abs((r1.X + r1.X + r1.Width) - (r2.X + r2.X + r2.Width))
	ww := r1.X + r1.Width + r2.X + r2.Width - r1.X - r2.X

	h := math.Abs((r1.Y + r1.Y + r1.Height) - (r2.Y + r2.Y + r2.Height))
	hh := r1.Y + r1.Height + r2.Y + r2.Height - r1.Y - r2.Y
	if w < ww && h < hh {
		return true
	}

	return false
}
