package main

import (
	"fmt"
)

//Паттерн адаптер позволяет объектам с несовместимыми интерфейсами работать вместе.

//Pазъём HDMI
type connector struct {
}

//Интерфей для подключения кабеля HDMI
type Connector interface {
	insertHDMI()
}

//Вставляем HDMI кабель/адаптер в разъём с помощью интерфейса Connector
func (c *connector) Insert(cable Connector) {
	fmt.Println("Inserting cable")
	cable.insertHDMI()
}

//HDMI кабель
type HDMI struct {
}

//Вставляет кабель в разъём HDMI
func (h *HDMI) insertHDMI() {
	fmt.Println("HDMI plugged into computer\n")
}

//VGA кабель
type VGA struct {
}

//Вставляет кабель в разъём VGA
func (v *VGA) insertVGA() {
	fmt.Println("VGA plugged into computer")
}

//адаптер, который позволяет подключать VGA в HDMI интерфейс
type VGAtoHDMI struct {
	vga *VGA
}


func (v *VGAtoHDMI) insertHDMI() {
	fmt.Println("Connecting VGA cable into computer via adapter ")
	v.vga.insertVGA()
}

func main() {
	connector := connector{}

	hdmi := HDMI{}
	vga := VGA{}

	adapter := VGAtoHDMI{vga: &vga}

	connector.Insert(&hdmi)
	connector.Insert(&adapter)
}
