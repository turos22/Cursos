package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDone     chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clientes", barber)

		for {
			if len(shop.ClientsChan) == 0{
				isSleeping = true
				color.Yellow("There is nothing to do, so %s goes to sleep", barber)				
			}

			client, shopOpen := <- shop.ClientsChan
			if shopOpen {
				if isSleeping{
					color.Yellow("%s wakes %s up", client, barber)
					isSleeping = false					
				}
				shop.cutHair(barber, client)
			} else {
				shop.sendBarberHome(barber)			
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s`s hair.", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finished cutting %s`s hair", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string){
	color.Cyan("%s is going home", barber)
	//caso dois barbeiros tentem enviar, por nao ser um canal bufferizado
	//ele fica bloqueado ate que o valor seja lido
	shop.BarbersDone <- true
}

func (shop *BarberShop) closeShopForDay(){
	color.Cyan("Closing shop fot the day")

	close(shop.ClientsChan)
	shop.Open = false

	for a:=1; a <= shop.NumberOfBarbers; a++{
		<-shop.BarbersDone
	}

	close(shop.BarbersDone)

	color.Green("----------------------------")
	color.Green("Shop is now closed for the day")
}

func (shop *BarberShop) addClient(client string){
	color.Green(" *** %s arrives!", client)

	if shop.Open{
		select {
			case shop.ClientsChan <- client:
				color.Yellow("%s takes a seat in the waiting room", client)
			default:
				color.Red("Waiting is full, so %s leaves", client)
		}
	} else {
		color.Red(" *** Shop is closed, %s leaves", client)
	}
}