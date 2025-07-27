package svg

import (
	"bytes"
	"fmt"
)

type SVG struct {
	width  float64
	height float64
	buf    *bytes.Buffer
}

func New(width, height float64) *SVG {
	s := &SVG{
		width:  width,
		height: height,
		buf:    bytes.NewBuffer(nil),
	}
	s.buf.WriteString(fmt.Sprintf(`<svg width="%.2f" height="%.2f" xmlns="http://www.w3.org/2000/svg">`, width, height))
	return s
}

func (s *SVG) String() string {
	s.buf.WriteString("</svg>")
	return s.buf.String()
}

func (s *SVG) AddElement(elem string) {
	s.buf.WriteString(elem)
}