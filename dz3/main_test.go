package main

import (
	"errors"
	"testing"
)

type MockNotifier struct {
	SendFunc func(message string) error
}

func (m MockNotifier) Send(message string) error {
	if m.SendFunc != nil {
		return m.SendFunc(message)
	}
	return nil
}

func TestDispatch_Success(t *testing.T) {
	called := 0
	mock := MockNotifier{
		SendFunc: func(message string) error {
			called++
			if message != "тест" {
				t.Errorf("должно было быть сообщение тест, а вышло '%s'", message)
			}
			return nil
		},
	}

	notifiers := []Notifier{mock}
	err := dispatch(notifiers, "тест")

	if err != nil {
		t.Fatalf("ошибок не должно было быть: %v", err)
	}

	if called != 1 {
		t.Errorf("должно было быть 1 вызов Send, а вызовов: %d", called)
	}
}

func TestDispatch_Error(t *testing.T) {
	expectedErr := errors.New("тестовая ошибка")
	mock := MockNotifier{
		SendFunc: func(message string) error {
			return expectedErr
		},
	}

	notifiers := []Notifier{mock}
	err := dispatch(notifiers, "тест")

	if err != expectedErr {
		t.Fatalf("ожидалась ошибка %v, вышла %v", expectedErr, err)
	}
}
