/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.

Входные данные

Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.

Выходные данные

Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения в первом числе.

Sample Input:

564 8954
Sample Output:

5 4
*/

package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a)

	fmt.Scan(&b)

	// Инвертируем порядок цифр в первом числе
	reversedA := 0
	tempA := a
	for tempA > 0 {
		digit := tempA % 10
		reversedA = reversedA*10 + digit
		tempA /= 10
	}

	for reversedA > 0 {
		d := reversedA % 10
		reversedA /= 10

		tempB := b
		for tempB > 0 {
			if tempB%10 == d {
				fmt.Print(d, " ")
				break
			}
			tempB /= 10
		}
	}
}
