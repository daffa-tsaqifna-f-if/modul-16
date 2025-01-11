package main

import "fmt"

func main() {
	var x float64
	var y, z int
	for {
		fmt.Scan(&y)
		if y == 9999 {
			break
		}
		z++
		x = x + float64(y)
	}
	x = x / float64(z)
	fmt.Print(x)
}
