package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateStatistic(orders []int) (int, int, int, float64) {
	min := orders[0]
	max := orders[0]
	sum := 0

	for _, order := range orders {
		if order < min {
			min = order
		}
		if order > max {
			max = order
		}
		sum += order
	}

	average := float64(sum) / float64(len(orders))
	return min, max, sum, average
}

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	orders := make([]int, 0, len(strings.Fields(input)))

	for _, part := range strings.Fields(input) {
		order, _ := strconv.Atoi(part)
		orders = append(orders, order)
	}

	if len(orders) == 0 {
		fmt.Println("Неа, пустым нельзя")
		return
	}

	min, max, sum, average := calculateStatistic(orders)

	fmt.Println("минимум:", min)
	fmt.Println("максимум:", max)
	fmt.Println("сумма:", sum)
	fmt.Printf("среднее: %.4f\n", average)
}
