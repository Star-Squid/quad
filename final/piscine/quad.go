package piscine

import "fmt"

func QuadA(x int, y int) {

	//these will be x-2 and y-2
	var widthBody int
	var lengthBody int

	//check if x and y are positive and >0
	if x <= 0 || y <= 0 {
		return
	}

	//top line - always print once, maybe repeat on bottom
	printLineOne := func() {
		fmt.Print("o")

		if x > 2 {
			widthBody = x - 2

			for i := widthBody; i > 0; i-- {
				fmt.Print("-")
			}
		}

		if x > 1 {
			fmt.Print("o")
		}

		fmt.Print("\n")
	}

	printLineOne()

	// body of the rectangle: lines between 1 and last
	printMiddleLine := func() {
		fmt.Print("|")

		if x > 2 {
			widthBody = x - 2

			for i := widthBody; i > 0; i-- {
				fmt.Print(" ")
			}
		}

		if x > 1 {
			fmt.Print("|")
		}
		fmt.Print("\n")
	}

	if y > 2 {
		lengthBody = y - 2

		for j := lengthBody; j > 0; j-- {
			printMiddleLine()
		}

	}

	//bottom line
	if y > 1 {
		printLineOne()
	}
}

func QuadB(x int, y int) {

	//these will be x-2 and y-2
	var widthBody int
	var lengthBody int

	//check if x and y are positive and >0
	if x <= 0 || y <= 0 {
		return
	}

	//top line - always print once
	fmt.Print("/")

	if x > 2 {
		widthBody = x - 2

		for i := widthBody; i > 0; i-- {
			fmt.Print("*")
		}
	}

	if x > 1 {
		fmt.Print("\\")
	}

	fmt.Print("\n")

	// body of the rectangle: lines between 1 and last
	printMiddleLine := func() {
		fmt.Print("*")

		if x > 2 {
			widthBody = x - 2

			for i := widthBody; i > 0; i-- {
				fmt.Print(" ")
			}
		}

		if x > 1 {
			fmt.Print("*")
		}
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
		fmt.Print("\\")

		if x > 2 {
			widthBody = x - 2

			for i := widthBody; i > 0; i-- {
				fmt.Print("*")
			}
		}

		if x > 1 {
			fmt.Print("/")
		}

		fmt.Print("\n")
	}
}

func QuadC(x int, y int) {

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
		fmt.Print("A")
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
		if x > 1 {
			fmt.Print("B")
		}
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
		fmt.Print("C")

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
		if x > 1 {
			fmt.Print("B")
		}
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

func QuadE(x int, y int) {

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
		if x > 1 {
			fmt.Print("B")
		}
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
		fmt.Print("C")

		if x > 2 {
			widthBody = x - 2

			for i := widthBody; i > 0; i-- {
				fmt.Print("B")
			}
		}

		if x > 1 {
			fmt.Print("A")
		}

		fmt.Print("\n")
	}
}
