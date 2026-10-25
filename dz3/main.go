package main

import (
	"errors"
	"fmt"
)

// 1 пункт
type Notifier interface {
	Send(message string) error
}

// проверка интерфейса на этапе компиляции
var _ Notifier = EmailNotifier{}
var _ Notifier = PushNotifier{}

// 7 пункт
type BaseNotifier struct {
	Channel string
}

func (b BaseNotifier) Log(message string) {
	fmt.Printf("[%s] %s\n", b.Channel, message)
}

// 2 пункт
type EmailNotifier struct {
	BaseNotifier
	Address string
}

func (n EmailNotifier) Send(message string) error {
	// ошибка пустого сообщения
	if message == "" {
		return errors.New("ошибка: пустое сообщение для email")
	}
	fmt.Printf("email to %s: %s\n", n.Address, message)
	return nil
}

// 3 пункт
type PushNotifier struct {
	BaseNotifier
	DeviceID string
}

func (n PushNotifier) Send(message string) error {
	fmt.Println("push:", message)
	return nil
}

// 5 пункт
func dispatch(all []Notifier, message string) error {
	for _, notifier := range all {
		if err := notifier.Send(message); err != nil {
			return err
		}
	}
	return nil
}

// 6 пункт
func describe(n Notifier) string {
	switch value := n.(type) {
	case EmailNotifier:
		return "email: " + value.Address
	case PushNotifier:
		return "push: " + value.DeviceID
	default:
		return "неизвестный канал"
	}
}

func main() {
	// 4 пункт
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

	err := dispatch(notifiers, "вперёд заказ")
	if err != nil {
		fmt.Println("ошибка: ", err)
	}
	for _, notifier := range notifiers {
		fmt.Println(describe(notifier))
	}
}

// паттерны ооп освоил
