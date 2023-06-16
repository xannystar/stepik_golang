/*
Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.

# Входные данные

Вводится натуральное число.

# Выходные данные

Выведите ответ на задачу.

Sample Input:

8
Sample Output:

6
*/
package main

import "fmt"

func main() {
	var a int
	for {
		fmt.Scan(&a)
		if a > 1 {
			break
		}
	}
	previous, current, i := 1, 1, 2
	for current < a {
		next := current + previous
		previous, current = current, next
		i++
	}
	if current == a {
		fmt.Print(i)
	} else if current > a {
		fmt.Print(-1)
	}
}
