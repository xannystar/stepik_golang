/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные

Вводится четыре числа.

Выходные данные

Необходимо вернуть из функции наименьшее из 4-х данных чисел

Sample Input:

4 5 6 7
Sample Output:

4
*/
func minimumFromFour() int {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if a < b && a < c && a < d {
		return a
	} else if b < a && b < c && b < d {
		return b
	} else if c < a && c < b && c < d {
		return c
	} else {
		return d
	}
}