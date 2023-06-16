/*
Заданы три числа - a,b,c (a<b<c) - длины сторон треугольника.
Нужно проверить, является ли треугольник прямоугольным. Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"

Sample Input:

6 8 10
Sample Output:

Прямоугольный
*/
package main

import "fmt"

func main() {
	for {
		var a, b, c uint
		fmt.Scan(&a, &b, &c)
		switch {
		case a < b && b < c:
			switch {
			case a*a+b*b == c*c:
				fmt.Println("Прямоугольный")
			default:
				fmt.Println("Непрямоугольный")
			}
		}
		break
	}
}
