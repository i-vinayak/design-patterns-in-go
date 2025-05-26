package main

// This design pattern is very rare in GO. But still a good to know

import "fmt"

type Light struct{}

func (l *Light) On() {
	fmt.Println("Light is ON")
}

func (l *Light) Off() {
	fmt.Println("Light is OFF")
}

type Command interface {
	Execute()
}

type TurnOnCommand struct {
	light *Light
}

func (c *TurnOnCommand) Execute() {
	c.light.On()
}

type TurnOffCommand struct {
	light *Light
}

func (c *TurnOffCommand) Execute() {
	c.light.Off()
}

type RemoteControl struct {
	commands []Command
}

func (r *RemoteControl) Submit(command Command) {
	r.commands = append(r.commands, command)
	command.Execute()
}

func main() {
	light := &Light{}

	onCommand := &TurnOnCommand{light: light}
	offCommand := &TurnOffCommand{light: light}

	remote := &RemoteControl{}
	remote.Submit(onCommand)  // Light is ON
	remote.Submit(offCommand) // Light is OFF
}
