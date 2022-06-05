package main

import (
	"fmt"
)

// 	// if it doesn't containt first line use this
// func numbers(s string) []int {
// 	var n []int
// 	for _, f := range strings.Fields(s) {
// 		i, err := strconv.Atoi(f)
// 		if err == nil {
// 			n = append(n, i)
// 		}
// 	}
// 	return n
// }

func main() {
	// // if it doesn't containt first line use this
	// var lines []int
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// lines = numbers(scanner.Text())
	// fmt.Println(lines)

	// if there's first line, use this
	length := 0
	fmt.Scanln(&length)
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanln(&numbers[i])
	}
	// fmt.Println(numbers)
	fmt.Println(compareTwoSide(numbers, length))
}

func compareTwoSide(arr []int, n int) string {
	const (
		Yes = "YES"
		No  = "NO"
	)

	for i := 1; i < n; i++ {
		leftElem := 0
		for j := i - 1; j >= 0; j-- {
			leftElem += arr[j]
		}

		rightElem := 0
		for k := i + 1; k < n; k++ {
			rightElem += arr[k]
		}
		// fmt.Println("Debug right", rightElem)
		// fmt.Println("Debug left", leftElem)
		if leftElem == rightElem {
			return Yes
		}
	}
	return No
}
