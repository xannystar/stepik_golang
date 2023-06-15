/*
Напишите программу, принимающая на вход число N(N≥4), а затем N целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

Sample Input:

5
41 -231 24 49 6
Sample Output:

49
*/
package main

import (
	"fmt"
)

func main() {
	var n int
	for n < 4 {
		fmt.Scan(&n)
	}
	slice := make([]int, n)
	for idx := range slice {
		fmt.Scan(&slice[idx])
	}
	fmt.Println(slice[3])
}
