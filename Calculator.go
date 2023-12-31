package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Калькулятор")
	fmt.Println("Нажмите q для выхода")
	for {
		var s string
		fmt.Println("Знак (+, -, *, /, **)")
		fmt.Scanln(&s)
		if s == "q" {
			break
		}
		if s == "+" || s == "-" || s == "*" || s == "/" || s == "**" {
			var x float64
			var y float64
			fmt.Println("x=")
			fmt.Scanln(&x)
			fmt.Println("y=")
			fmt.Scanln(&y)

			if s == "+" {
				fmt.Println(x + y)
			} else if s == "-" {
				fmt.Println(x - y)
			} else if s == "*" {
				fmt.Println(x * y)
			} else if s == "**" {
				fmt.Println(math.Pow(x, y))
			} else if s == "/" {
				if y != 0 {
					fmt.Println(x / y)
				} else {
					fmt.Println("Ошибка")
				}
			}
		} else {
			fmt.Println("Неверный ввод")
		}
	}
}
