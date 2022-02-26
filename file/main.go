package main

import (
	"file/controller/http"
	"file/minio"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go minio.Run()
	go http.Run(":8003")
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
