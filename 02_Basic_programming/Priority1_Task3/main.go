
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%4 == 0 && i%7 == 0:
			fmt.Println("buzz")
		case i%4 == 0:
			fmt.Println(i * i)
		case i%7 == 0:
			fmt.Println(i * i * i)
		default:
			fmt.Println(i)
		}
	}
}
