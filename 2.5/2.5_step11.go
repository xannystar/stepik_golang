/*
Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку

Sample Input:

zaabcbd
Sample Output:

zcd
*/
package main

import "fmt"

func main() {
	var str, result string
	fmt.Scan(&str)
	freq := make([]int, 256)
	for _, ch := range str {
		freq[ch]++
	}
	for _, ch := range str {
		if freq[ch] == 1 {
			result += string(ch)
		}
	}
	fmt.Println(result)
}
