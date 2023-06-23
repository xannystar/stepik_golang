/*
На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").

Sample Input:

топот
Sample Output:

Палиндром
*/
package main

import "fmt"

func main() {
	var str, revStr string
	fmt.Scan(&str)
	for _, a := range str {
		revStr = string(a) + revStr
	}
	if str == revStr {
		fmt.Print("Палиндром")
	} else {
		fmt.Print("Нет")
	}
}