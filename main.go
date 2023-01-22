package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatingCapacity = 10
var arriveRate = 500
var cutDuration = 2000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	//seed our random number generator
	rand.Seed(time.Now().UnixNano())
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	//create channels if any
	clientsChannel := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	//create the barbershop
	shop := &BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientsChannel,
		BarbersDoneChan: doneChan,
		Open:            true,
	}
	color.Green("The shop is open for the day!")

	//add barber
	shop.AddBarber("Mohan")

	// start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.CloseShopForDay()
		closed <- true
	}()

	//add clients
	i := 1

	go func() {
		for {
			//get a random number with average arrival rate
			randomMilliseconds := rand.Int() % (2 * arriveRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.AddClients(fmt.Sprintf("Client #%d", i))
				i++
			}
		}

	}()

	// time.Sleep(time.Second * 5)
	<-closed

}
