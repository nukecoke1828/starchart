package svg

import "math"

func PolarToCartesian(centerX, centerY, radius, angleInDegrees float64) (float64, float64) {
	angleInRadians := (angleInDegrees - 90) * math.Pi / 180.0
	x := centerX + (radius * math.Cos(angleInRadians))
	y := centerY + (radius * math.Sin(angleInRadians))
	return x, y
}

func DescribeArc(x, y, radius, startAngle, endAngle float64) string {
	start := PolarToCartesian(x, y, radius, endAngle)
	end := PolarToCartesian(x, y, radius, startAngle)

	largeArcFlag := "0"
	if endAngle-startAngle > 180 {
		largeArcFlag = "1"
	}

	d := []string{
		"M", fmt.Sprintf("%.2f", start[0]), fmt.Sprintf("%.2f", start[1]),
		"A", fmt.Sprintf("%.2f", radius), fmt.Sprintf("%.2f", radius),
		"0", largeArcFlag, "0",
		fmt.Sprintf("%.2f", end[0]), fmt.Sprintf("%.2f", end[1]),
	}

	return strings.Join(d, " ")
}