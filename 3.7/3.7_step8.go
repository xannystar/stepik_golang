/*
На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей преобразования равны 0).

Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.

Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.

Sample Input:

12 мин. 13 сек.
Sample Output:

Fri May 15 19:28:18 UTC 2020
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Константа Unix-Time в виде int64
	const baseUnixTime = 1589570165

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Разбиваем строку на части (минуты и секунды)
	parts := strings.Split(input, " ")
	if len(parts) != 4 {
		panic("Invalid input format")
	}

	// Получаем значение минут и секунд из строк и преобразуем их в числа
	minutes, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}

	// Преобразуем продолжительность периода в тип time.Duration
	duration := time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second

	// Преобразуем базовую дату в тип time.Time и переносим её в зону UTC
	baseTime := time.Unix(baseUnixTime, 0).UTC()

	// Добавляем период к базовой дате
	resultTime := baseTime.Add(duration)

	// Выводим полученную дату и время в формате UnixDate
	fmt.Println(resultTime.Format(time.UnixDate))
}
