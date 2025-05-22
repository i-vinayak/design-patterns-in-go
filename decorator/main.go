package main

import "fmt"

type Notifier interface {
	Send(string)
}

// struct with Send functionality implemented
type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Println(message)
}

type LogNotifier struct {
	notifier Notifier //Composition
}

func (l *LogNotifier) Send(message string) {
	l.notifier.Send(message)
}

type RandomNotifier struct {
	notifier Notifier
}

// Extend the functionality of this using other struct by embedding it in this struct
func (r *RandomNotifier) Send(message string) {
	r.notifier.Send(message)
}

func main() {
	notifier := &EmailNotifier{}
	logNotifier := &LogNotifier{notifier: notifier}
	randomNotifier := &RandomNotifier{notifier: notifier}

	notifier.Send("Email notifier")
	logNotifier.Send("log")
	randomNotifier.Send(("Random"))
}
