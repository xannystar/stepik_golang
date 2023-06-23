/*
На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

Sample Input:

ihgewlqlkot
Sample Output:

hello
*/
package main

import "fmt"

func main() {
	var input, result string
	fmt.Scan(&input)

	for i, char := range input {
		if i%2 != 0 {
			result += string(char)
		}
	}
	fmt.Println(result)
}
