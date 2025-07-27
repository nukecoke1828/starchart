package svg

import "fmt"

func Text(x, y float64, content string, style Style) string {
	return fmt.Sprintf(`<text x="%.2f" y="%.2f" style="%s">%s</text>`, 
		x, y, style, content)
}