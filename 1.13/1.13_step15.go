/*
Двоичная запись
Дано натуральное число N. Выведите его представление в двоичном виде.

# Входные данные

# Задано единственное число N

# Выходные данные

Необходимо вывести требуемое представление числа N.

Sample Input:

12
Sample Output:

1100
*/
package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	binary := ""
	for num > 0 {
		bit := num % 2
		binary = fmt.Sprint(bit) + binary
		num /= 2
	}
	fmt.Println(binary)
}
