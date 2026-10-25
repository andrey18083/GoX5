package main

import (
	"errors"
	"fmt"
)

var ErrEmptyMessage = errors.New("пустое сообщение")

type Notifier interface {
	Send(message string) error
}

type BaseNotifier struct {
	Channel string
}

func (b BaseNotifier) Log(message string) {
	fmt.Printf("[%s] %s\n", b.Channel, message)
}

type EmailNotifier struct {
	BaseNotifier
	Address string
}

func (n EmailNotifier) Send(message string) error {
	if message == "" {
		return &CustomError{
			Channel: n.Channel,
			Code:    400,
			Err:     ErrEmptyMessage,
		}
	}
	n.Log(fmt.Sprintf("email to %s: %s", n.Address, message))
	return nil
}

type PushNotifier struct {
	BaseNotifier
	DeviceID string
}

func (n PushNotifier) Send(message string) error {
	if message == "panic" {
		panic("драйвер push-уведомлений поврежден")
	}
	n.Log(fmt.Sprintf("push to device %s: %s", n.DeviceID, message))
	return nil
}

func safeSend(n Notifier, message string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = &CustomError{
				Channel: "dispatcher_boundary",
				Code:    500,
				Err:     fmt.Errorf("восстановлено после паники: %v", r),
			}
		}
	}()

	return n.Send(message)
}

func dispatch(all []Notifier, message string) error {
	for _, notifier := range all {
		if err := safeSend(notifier, message); err != nil {
			return fmt.Errorf("ошибка в диспетчере: %w", err)
		}
	}
	return nil
}

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
