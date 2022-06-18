package piscine

import "fmt"

func QuadD(x int, y int) {

	//these will be x-2 and y-2
	var widthBody int
	var lengthBody int

	//check if x and y are positive and >0
	if x <= 0 || y <= 0 {
		return
	}

	//top line - always print once
	fmt.Print("A")

	if x > 2 {
		widthBody = x - 2

		for i := widthBody; i > 0; i-- {
			fmt.Print("B")
		}
	}

	if x > 1 {
		fmt.Print("C")
	}

	fmt.Print("\n")

	// body of the rectangle: lines between 1 and last
	printMiddleLine := func() {
		fmt.Print("B")

		if x > 2 {
			widthBody = x - 2

			for i := widthBody; i > 0; i-- {
				fmt.Print(" ")
			}
		}
		fmt.Print("B")
		fmt.Print("\n")
	}

	if y > 2 {
		lengthBody = y - 2

		for j := lengthBody; j > 0; j-- {
			printMiddleLine()
		}

	}

	//bottom line, print if y is at least 2
	if y > 1 {
		fmt.Print("A")

		if x > 2 {
			widthBody = x - 2

			for i := widthBody; i > 0; i-- {
				fmt.Print("B")
			}
		}

		if x > 1 {
			fmt.Print("C")
		}

		fmt.Print("\n")
	}
}
