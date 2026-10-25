package main

import (
	"errors"
	"fmt"
)

func main() {
	notifiers := []Notifier{
		EmailNotifier{
			BaseNotifier: BaseNotifier{Channel: "email"},
			Address:      "student@example.com",
		},
		PushNotifier{
			BaseNotifier: BaseNotifier{Channel: "push"},
			DeviceID:     "device-42",
		},
	}

	if err := dispatch(notifiers, "вперёд заказ"); err != nil {
		fmt.Println("ошибка рассылки:", err)
	}

	if err := dispatch(notifiers, ""); err != nil {
		fmt.Println("ошибка рассылки:", err)

		if errors.Is(err, ErrEmptyMessage) {
			fmt.Println("ожидаемая ошибка перехвачена ErrEmptyMessage!")
		}
	}

	if err := dispatch(notifiers, "panic"); err != nil {
		fmt.Println("ошибка рассылки:", err)
	}

	for _, notifier := range notifiers {
		fmt.Println(describe(notifier))
	}
}
