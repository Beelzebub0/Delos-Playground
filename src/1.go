package main

import "fmt"

func main() {
	var expectedDate [3]int
	var actualDate [3]int
	var fine int
	fmt.Scan(&expectedDate[0], &expectedDate[1], &expectedDate[2])
	fmt.Scan(&actualDate[0], &actualDate[1], &actualDate[2])

	for i := 0; i < 3; i++ {

		switch {
		case (actualDate[0] > expectedDate[0]) && (actualDate[1] == expectedDate[1]) && (actualDate[2] == expectedDate[2]):
			fine = (actualDate[0] - expectedDate[0]) * 15
		case (actualDate[1] > expectedDate[1]) && (actualDate[2] == expectedDate[2]):
			fine = (actualDate[1] - expectedDate[1]) * 500
		case (actualDate[2] > expectedDate[2]):
			fine = (actualDate[2] - expectedDate[2]) * 12000
		default:
			fine = 0
		}
	}
	fmt.Println(fine)
}
