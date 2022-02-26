package main

import (
	"automobile/controller/http"
	"automobile/repository"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go repository.Init()
	//go rabbitmq.Init()
	go http.Run(":8001")
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		select {
		case <-sigc:
			// wg.Done()
			os.Exit(1)
		}
	}()
	wg.Wait()
}
