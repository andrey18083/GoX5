package main

import (
	"testing"
)

func TestNewOrderZeroTotal(t *testing.T) {
	order, err := NewOrder(1)
	if err != nil {
		t.Fatalf("ошибка: %v", err)
	}

	if Total(order) != 0 {
		t.Errorf("должна была быть сумма, а вышло %d", Total(order))
	}
}

func TestOrderTotalAfterAddItems(t *testing.T) {
	order, _ := NewOrder(2)

	_ = AddItem(&order, Item{Name: "капуста", Price: 322})
	_ = AddItem(&order, Item{Name: "куст", Price: 223})

	expectedSum := int64(545)
	if Total(order) != expectedSum {
		t.Errorf("ожидалась сумма %d, а получили %d", expectedSum, Total(order))
	}
}

func TestAddItemValidation(t *testing.T) {
	var nilOrder *Order
	err := AddItem(nilOrder, Item{Name: "арбуз", Price: 2026})
	if err == nil {
		t.Error("ошибка для nil заказа не высветилась")
	}

	order, _ := NewOrder(3)

	err = AddItem(&order, Item{Name: "", Price: 100})
	if err == nil {
		t.Error("ошибка для пустого имени товара не высветилась")
	}

	err = AddItem(&order, Item{Name: "вода", Price: -10})
	if err == nil {
		t.Error("ошибка для отрицательной цены товара не высветилась")
	}
}

func TestOrderCopyImmutability(t *testing.T) {
	original, _ := NewOrder(4)
	_ = AddItem(&original, Item{Name: "молочко", Price: 100})

	copyOrder := original
	copyOrder.ID = 999

	_ = AddItem(&copyOrder, Item{Name: "творожок", Price: 145})

	if original.ID == 999 {
		t.Error("изменение id в копии повлияло на оригинал")
	}
	if len(original.Items) != 1 {
		t.Errorf("ожидался 1 элемент, а получилось %d", len(original.Items))
	}
}

func TestMapInitialization(t *testing.T) {
	var customStatuses map[string]string

	customStatuses = make(map[string]string)
	customStatuses["new"] = "сделан"

	if customStatuses["new"] != "сделан" {
		t.Error("ошибка в инициализации map")
	}
}
