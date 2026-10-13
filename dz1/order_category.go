
package main

import "fmt"

func getCategory(amount int) string {
	if amount < 10000 {
		return "small"
	} else if amount < 50000 {
		return "regular"
	} else {
		return "large"
	}
}

func main() {
	var amount int
	fmt.Scan(&amount)

	if amount < 0 {
		fmt.Println("Неа, отрицательные нельзя")
		return
	}

	category := getCategory(amount)
	fmt.Println("Категория заказа:", category)
}
