/*
Дано трехзначное число. Найдите сумму его цифр.

Формат входных данных
На вход дается трехзначное число.

Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.

Sample Input:

745
Sample Output:

16
*/
package main

import "fmt"

func main() {
	for {
		var num int
		fmt.Scan(&num)
		if num > 99 && num < 1000 {
			fmt.Print((num / 100) + (num / 10 % 10) + (num % 10))
			break
		}
	}
}
