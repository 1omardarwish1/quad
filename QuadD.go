package piscine

import "fmt"


func QuadD(x, y int) {
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if i == 1 && j == 1 {
					fmt.Print("A")
				} else if i == 1 && j == x {
					fmt.Print("C")
				} else if i == y && j == 1 {
					fmt.Print("A")
				} else if i == y && j == x {
					fmt.Print("C")
				} else if i == 1 || i == y || j == 1 || j == x {
					fmt.Print("B")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
}