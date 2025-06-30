package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := input()
	nums := convertStringToSlice(input)
	res := sumAll(nums...)
	fmt.Println(res)
}

func sumAll(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите числа через пробел: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	return input
}

func convertStringToSlice(s string) []int {
	s = strings.TrimSpace(s)
	strNums := strings.Split(s, " ")
	var nums []int
	for _, strNum := range strNums {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Printf("Ошибка преобразования строки %s в число, этот элемент будет пропущен\n", strNum)
			continue
		}
		nums = append(nums, num)
	}
	return nums
}
