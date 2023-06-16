/*
Количество нулей
По данным числам, определите количество чисел, которые равны нулю.

Входные данные
Вводится натуральное число N, а затем N чисел.

Выходные данные
Выведите количество чисел, которые равны нулю.

Sample Input:

5
1
8
100
0
12
Sample Output:

1
*/
package main

import "fmt"

func main() {
	var n, nums, counter int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums)
		if nums == 0 {
			counter++
		}
	}
	fmt.Print(counter)
}
