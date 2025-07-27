package svg

import "fmt"

func Path(d string, style Style) string {
	return fmt.Sprintf(`<path d="%s" style="%s" />`, d, style)
}