package main

import (
	"fmt"
	"math"
)

func main() {
	// input number
	n := 265149
	// determine ring
	ringNo := int(math.Ceil(math.Sqrt(float64(n)))) / 2
	// calc highest number in ring
	ringMax := int(math.Pow(float64(ringNo*2+1), 2))
	// calc hightest number in inner ring
	ringMin := int(math.Pow(float64((ringNo-1)*2+1), 2))
	// calc length between each midpoint in ring
	segment := (ringMax - ringMin) / 4
	// calc first midpoint in ring
	firstMid := ringMin + segment/2
	// calc all midpoints in ring
	midpoints := [4]int{firstMid, firstMid + segment, firstMid + segment*2, firstMid + segment*3}

	// determine closest midpoint from n
	// starting with first midpoint
	midpoint := midpoints[0]
	distance := math.Abs(float64(n - midpoint))
	for i := 1; i < 4; i++ {
		// update midpoint and distance if new distance < previous distance
		if math.Abs(float64(n-midpoints[i])) < distance {
			midpoint = midpoints[i]
			distance = math.Abs(float64(n - midpoints[i]))
		}
	}
	// print distance to midpoint + distance to center
	fmt.Println(int(distance) + ringNo)
}
