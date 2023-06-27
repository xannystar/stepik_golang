/*
Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.

# Входные данные

Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.

# Выходные данные

Выведите максимальную цифру, которая встречается во введенной строке.

Sample Input:

1112221112
Sample Output:

2
*/
package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Printf("%c", max(str))
}

func max(str string) rune {
	max := '0'
	for _, i := range str {
		if max < i {
			max = i
		}
	}
	return max
}
