package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float32
	var height float32

	// input radius
	fmt.Print("Enter the radius: ")
	fmt.Scanln(&radius)

	//  input height
	fmt.Print("Enter the height: ")
	fmt.Scanln(&height)

	// Calculate the surface area of the tube
	area := math.Pi * radius * (radius + height)

	fmt.Printf("Surface area of the tube: %.1f\n", area)
}
