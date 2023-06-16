/*
Из натурального числа удалить заданную цифру.

# Входные данные

Вводятся натуральное число и цифра, которую нужно удалить.

# Выходные данные

Вывести число без заданных цифр.

Sample Input:

38012732
3
Sample Output:

801272
*/
package main

import "fmt"

func main() {
	var num, del, a int
	fmt.Scan(&num, &del)
	sum := 0
	current := 1
	for num > 0 {
		a = num % 10
		if del != a {
			sum = sum + a*current
			current *= 10
		}
		num /= 10
	}
	fmt.Print(sum)
}
