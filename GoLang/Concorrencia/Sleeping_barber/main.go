package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/fatih/color"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	rand.Seed(time.Now().UnixNano())
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("----------------------------")

	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		BarbersDone:     doneChan,
		ClientsChan:     clientChan,
		Open:            true,
	}

	color.Green("The shop is open for the day!")

	shop.addBarber("Frank")
	shop.addBarber("Gerard")
	shop.addBarber("milton")
	shop.addBarber("Susan")
	shop.addBarber("Kelly")
	shop.addBarber("Penis")

	shopClosing := make(chan bool)
	//closed := make(chan bool)

	var wgCliente sync.WaitGroup
	wgCliente.Add(1)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
	}()

	i := 1
	go func() {
		defer wgCliente.Done()
		for {
			randomMiliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMiliseconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	wgCliente.Wait()

	shop.closeShopForDay()
}
