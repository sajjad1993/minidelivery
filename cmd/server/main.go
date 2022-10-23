package main

import (
	"fmt"
	"log"
	"minidelivery/internal/adapter/main_control"
	"minidelivery/internal/interactor"
	"minidelivery/internal/ui"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	count := 2
	logger := log.Logger{}
	fmt.Println("start service")
	repository := main_control.New(count, &logger)
	go repository.HandleFreeDeliveries()
	service := interactor.New(repository, &logger)
	listener := ui.New(service)
	fmt.Println("enter your destination in this format 'x y' for example '-1 3' 1 ")

	go func() {
		for {
			listener.Listen()

		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	fmt.Println("stopped Contract server")

}
