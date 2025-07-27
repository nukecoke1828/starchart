package chart

type Box struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (b Box) Contains(x, y float64) bool {
	return x >= b.X && x <= b.X+b.Width &&
		y >= b.Y && y <= b.Y+b.Height
}

func (b Box) Center() (float64, float64) {
	return b.X + b.Width/2, b.Y + b.Height/2
}