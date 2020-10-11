package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(makeStars(5))
}

func makeStars( val int ) (stars string) {
	stars = ""
	average:= (int(math.Ceil(float64(val) /float64(2))))

	if val % 2 == 1 {
		for x := 1; x <= val; x++ {
			for y := 0; y < val; y++ {
				if y == 0 || y == val-1 || x == average {
					stars += "* "
				} else {
					stars += "= "
				}
			}
			fmt.Println(stars)
			stars = ""
		}
	} else {
		fmt.Println("error : value must be an odd number")
	}
		return
}