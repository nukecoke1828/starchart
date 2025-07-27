package svg

type Style struct {
	Fill        string
	Stroke      string
	StrokeWidth float64
	Opacity     float64
}

func (s Style) String() string {
	style := ""
	if s.Fill != "" {
		style += fmt.Sprintf("fill:%s;", s.Fill)
	}
	if s.Stroke != "" {
		style += fmt.Sprintf("stroke:%s;", s.Stroke)
	}
	if s.StrokeWidth > 0 {
		style += fmt.Sprintf("stroke-width:%.2f;", s.StrokeWidth)
	}
	if s.Opacity > 0 {
		style += fmt.Sprintf("opacity:%.2f;", s.Opacity)
	}
	return style
}