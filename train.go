package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a*a+b*b == c*c {
		fmt.Println("Прямоугольный")
	} else if c*c+b*b == a*a {
		fmt.Println("Прямоугольный")
	} else if c*c+a*a == b*b {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
