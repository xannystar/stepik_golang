/*
Дано трехзначное число. Переверните его, а затем выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных
Выведите перевернутое число.

Sample Input:

843
Sample Output:

348
*/
package main

import "fmt"

func main() {
	for {
		var num int
		fmt.Scan(&num)
		if num > 99 && num < 1000 {
			fmt.Printf("%d%d%d", (num % 10), (num / 10 % 10), (num / 100))
			break
		}
	}
}
