package main

import (
	"time"

	"github.com/fatih/color"
)

func server1(ch chan string) {
	for {
		time.Sleep(time.Second * 2)
		ch <- "this is from server 1"
	}
}
func server2(ch chan string) {
	for {
		time.Sleep(time.Second * 4)
		ch <- "this is from server 2"
	}
}
func Demo() {
	color.Cyan("Select with Channels")
	color.Cyan("---------- Main Started---------")
	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)

	go server2(channel2)

	time.AfterFunc(30*time.Second, func() {
		println("30 seconds timeout")
	})

	for {
		select {
		case s1 := <-channel1:
			color.Green("%s is from channel1 ", s1)
		case s2 := <-channel1:
			color.Green("%s is from channel1", s2)
		case s3 := <-channel2:
			color.Red("%s is from channel1", s3)
		case s4 := <-channel2:
			color.Red("%s is from channel1", s4)
		default:
			color.Magenta(" In Default case")
			time.Sleep(1 * time.Second)
		}

	}
	close(channel1)
	close(channel2)
	color.Cyan("---------- Main Started---------")

}
