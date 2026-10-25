package main

import (
	"errors"
	"testing"
)

func TestDispatch_Success(t *testing.T) {
	notifiers := []Notifier{
		EmailNotifier{BaseNotifier: BaseNotifier{Channel: "email"}, Address: "test@test.com"},
		PushNotifier{BaseNotifier: BaseNotifier{Channel: "push"}, DeviceID: "device-1"},
	}

	err := dispatch(notifiers, "привет")
	if err != nil {
		t.Fatalf("ошибка: %v", err)
	}
}

func TestEmailNotifier_EmptyMessage(t *testing.T) {
	email := EmailNotifier{BaseNotifier: BaseNotifier{Channel: "email"}, Address: "test@test.com"}
	err := email.Send("")

	if err == nil {
		t.Fatal("нет ошибки, но должна быть, пусто же")
	}

	if !errors.Is(err, ErrEmptyMessage) {
		t.Errorf("ожидалась ошибка ErrEmptyMessage, а она другая: %v", err)
	}
}

func TestDispatch_PanicIsolation(t *testing.T) {
	notifiers := []Notifier{
		PushNotifier{BaseNotifier: BaseNotifier{Channel: "push"}, DeviceID: "dev-panic"},
	}

	err := dispatch(notifiers, "panic")
	if err == nil {
		t.Fatal("диспетчер должен был вернуть ошибку после паники, но вернул nil")
	}

	var customErr *CustomError
	if !errors.As(err, &customErr) {
		t.Fatalf("ожидался *CustomError, вышло: %T", err)
	}

	if customErr.Channel != "dispatcher_boundary" {
		t.Errorf("ожидался канал dispatcher_boundary, вышло: %s", customErr.Channel)
	}
}

func TestDescribe(t *testing.T) {
	email := EmailNotifier{BaseNotifier: BaseNotifier{Channel: "email"}, Address: "a@b.ru"}
	push := PushNotifier{BaseNotifier: BaseNotifier{Channel: "push"}, DeviceID: "123"}

	if desc := describe(email); desc != "email: a@b.ru" {
		t.Errorf("неверное описание почты: %s", desc)
	}

	if desc := describe(push); desc != "push: 123" {
		t.Errorf("неверное описание push: %s", desc)
	}
}
