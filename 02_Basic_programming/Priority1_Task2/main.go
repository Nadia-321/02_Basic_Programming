package main

import (
	"fmt"
)

func main() {
	var score int
	fmt.Print("Enter the score: ")
	fmt.Scanln(&score)

	var grade string

	if score >= 85 && score <= 100 {
		grade = "A"
	} else if score >= 70 && score <= 84 {
		grade = "B"
	} else if score >= 55 && score <= 69 {
		grade = "C"
	} else if score >= 40 && score <= 54 {
		grade = "D"
	} else if score >= 0 && score <= 39 {
		grade = "E"
	} else {
		grade = "Invalid score"
	}

	fmt.Println("Grade:", grade)
}
