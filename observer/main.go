package main

import (
	"fmt"
	"slices"
)

type Observer interface {
	Update(temperature float64)
}

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyAll()
}

type WeatherStation struct {
	observers   []Observer
	temperature float64
}

func (ws *WeatherStation) Register(observer Observer) {
	ws.observers = append(ws.observers, observer)
}

func (ws *WeatherStation) Unregister(observer Observer) {
	for i, obs := range ws.observers {
		if obs == observer {
			ws.observers = slices.Delete(ws.observers, i, i+1)
			break
		}
	}
}

func (ws *WeatherStation) NotifyAll() {
	for _, obs := range ws.observers {
		obs.Update(ws.temperature)
	}
}

func (ws *WeatherStation) SetTemperature(temp float64) {
	fmt.Println("Setting temperature to", temp)
	ws.temperature = temp
	ws.NotifyAll()
}

type PhoneDisplay struct {
	name string
}

func (p *PhoneDisplay) Update(temp float64) {
	fmt.Printf("[%s] Temperature updated: %.2f°C\n", p.name, temp)
}

func main() {
	station := &WeatherStation{}

	phone1 := &PhoneDisplay{name: "Phone 1"}
	phone2 := &PhoneDisplay{name: "Phone 2"}

	station.Register(phone1)
	station.Register(phone2)

	station.SetTemperature(24.5)
	station.SetTemperature(27.0)

	station.Unregister(phone2)
	station.SetTemperature(30.0)
}
