package main

import "fmt"

func main() {
	var inputData [3]int

	fmt.Scan(&inputData[0], &inputData[1], &inputData[2])

	students := inputData[0]
	candies := inputData[1]
	firstStudent := inputData[2]
	fmt.Println(findLastPerson(students, candies, firstStudent))
}

func findLastPerson(students, candies, firstStudent int) int {
	last := students - firstStudent + 1
	distCandy := candies - last
	lastPerson := distCandy % students
	if lastPerson == 0 {
		return students
	} else {
		return lastPerson
	}
}
