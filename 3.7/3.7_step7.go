/*
На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).

Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.

Sample Input:

13.03.2018 14:00:15,12.03.2018 14:00:15
Sample Output:

24h0m0s
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	dateStrings := strings.Split(input, ",")
	if len(dateStrings) != 2 {
		panic("Invalid input format")
	}

	layout := "02.01.2006 15:04:05"
	date1, err := time.Parse(layout, dateStrings[0])
	if err != nil {
		panic(err)
	}

	date2, err := time.Parse(layout, dateStrings[1])
	if err != nil {
		panic(err)
	}

	var duration time.Duration
	if date1.After(date2) {
		duration = date1.Sub(date2)
	} else {
		duration = date2.Sub(date1)
	}

	fmt.Println(duration)
}
