package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numberOfArgs := len(os.Args)

	// If number of args is 3, then we calculate
	// the ratio between first and second argument
	if numberOfArgs == 3 {
		x, y := calculateRatio(os.Args)
		// Needs to be converted to float to work
		p := (float64(y) / float64(x)) * float64(100)

		// Format
		line1 := fmt.Sprintf("The ratio is: %d:%d", x, y)
		line2 := fmt.Sprintf("The percentage is: %.2f%%", p)

		// Print
		fmt.Println(line1)
		fmt.Println(line2)
	}

	// If number of args is 5, this means the users
	// wants to know the proportion for a given number
	// according to the calculated ratio
	if numberOfArgs == 5 {
		x, y := calculateRatio(os.Args)

		newValue, _ := strconv.Atoi(os.Args[4])

		var line string
		if os.Args[3] == "w" {
			newWidth := float64(newValue) * float64(y) / float64(x)
			line = fmt.Sprintf("The proportional height is: %d", int64(newWidth))
		} else if os.Args[3] == "h" {
			newHeight := float64(newValue) * float64(x) / float64(y)
			line = fmt.Sprintf("The proportional width is: %d", int64(newHeight))
		}

		fmt.Println(line)
	}
}

// Calculates the ratio between the first and second argument
func calculateRatio(args []string) (xx int, yy int) {
	width, _ := strconv.Atoi(args[1])
	height, _ := strconv.Atoi(args[2])

	x := width
	y := height
	for y > 0 {
		t := y
		y = x % y
		x = t
	}

	return width / x, height / x
}
