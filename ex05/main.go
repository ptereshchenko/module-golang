package main

import (
	"./barbershop"
	"fmt"
	"time"
)

func main() {
	barber := barbershop.NewBarber("John")

	shop := barbershop.NewShop(barber, 4)

	go barber.ManageShop(shop)

	time.Sleep(time.Second * 2)

	clients := []string{"Alain"}
	clients = append(clients, "Bernard")
	clients = append(clients, "Claude")
	clients = append(clients, "Dan")
	clients = append(clients, "Eric")
	clients = append(clients, "Franck")
	clients = append(clients, "Gerard")
	clients = append(clients, "Henri")
	clients = append(clients, "Ignasse")

	for _, c := range clients {
		client := barbershop.NewClient(c)
		go client.EnterShop(shop)

	}
	fmt.Scanln()

}
