/*
Количество минимумов
Найдите количество минимальных элементов в последовательности.

# Входные данные

Вводится натуральное число N, а затем N целых чисел последовательности.

# Выходные данные

Выведите количество минимальных элементов последовательности.

Sample Input:

3
21
11
4
Sample Output:

1
*/
package main

import "fmt"

func main() {
	var n, nums, min int
	fmt.Scan(&n)
	slice := make([]int, n)
	for index := range slice {
		fmt.Scan(&nums)
		slice[index] = nums
	}
	current := slice[0]
	min = 1
	for i := 1; i < len(slice); i++ {
		if current > slice[i] {
			current = slice[i]
			min = 1
		} else if current == slice[i] {
			min++
		}
	}
	fmt.Print("\n", min)
}
