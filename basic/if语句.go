package main

func main16() {
	if true {
		println("This is true")
	} else {
		println("This is false")
	}

	if b := 8; b > 5 {
		println("b is greater than 5")
	}
	if c, d, e := 1, 4, 5; c < d && (c < e || d < e) {
		println("c is the smallest")
	}

}
