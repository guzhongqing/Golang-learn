package main

import (
	"bufio"
	"fmt"
	"os"
)

func acm_mode() {
	// fmt.Println("1 2")
	// fmt.Println("2 3")
	// fmt.Println("4 5")
	// fmt.Println("6 7")

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 2; i++ {
		scanner.Scan()
		fmt.Println(scanner.Text())
	}

}
