package svg

import "fmt"

func Rect(x, y, width, height float64, style Style) string {
	return fmt.Sprintf(`<rect x="%.2f" y="%.2f" width="%.2f" height="%.2f" style="%s" />`, 
		x, y, width, height, style)
}