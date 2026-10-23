package main

import (
	"errors"
)

// позиция в заказе
type Item struct {
	Name  string
	Price int64
}

// описание заказа
type Order struct {
	ID     int
	Items  []Item
	Status string
}

func NewOrder(id int) (Order, error) {
	if id <= 0 {
		return Order{}, errors.New("id должно быть положительным")
	}
	return Order{
		ID:     id,
		Items:  []Item{},
		Status: "новый",
	}, nil
}

func AddItem(o *Order, item Item) error {
	if o == nil {
		return errors.New("nil order")
	}
	if item.Name == "" {
		return errors.New("item name не может быть пустым")
	}
	if item.Price < 0 {
		return errors.New("item price не может быть отрицательной")
	}

	o.Items = append(o.Items, item)
	return nil
}

func Total(o Order) int64 {
	var sum int64 = 0
	for _, item := range o.Items {
		sum += item.Price
	}
	return sum
}

var statuses = map[int]string{
	1: "сделан",
	2: "в обработке",
	3: "доставка закза",
	4: "исполнено",
}

func GetStatusDescription(statusCode int) string {
	if name, exists := statuses[statusCode]; exists {
		return name
	}
	return "неизвестный статус"
}

/*
Поставьте breakpoint в AddItem и проследите, как меняются len и cap среза:

До добавления элементов len = 0, cap = 0 --- память под базовый массив еще не выделена
После добавления 1-го элемента len = 1, cap = 1 --- длина увеличилась на 1, выделилась память, но только первая ячейка
После добавления 2-го элемента len = 2, cap = 2 --- длина среза автоматически удволиась,
так как старый массив был полностью заполнен, и язык сам перераспределил память под новый массив
*/
