package main

import (
	"github.com/01-edu/z01"
	"fmt"
)

func main() {
	var num int = 123
	t := fmt.Sprintf("%d", num)
	// fmt.Printf(t) # also prints
	slice := []rune(t)
	for i := 0; i < len(slice); i++ {
		z01.PrintRune(slice[i])
	}
	z01.PrintRune('\n')
}
