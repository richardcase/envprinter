package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	fmt.Printf("Starting Environment Printer.......")

	stop := make(chan struct{})
	go dumpenvironment()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown signal received, exiting...")
	close(stop)
}

func dumpenvironment() {
	fmt.Printf("")
	fmt.Printf("*******************************")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Printf("%s = %s\n", pair[0], pair[1])
	}
	time.Sleep(30 * time.Second)
}
