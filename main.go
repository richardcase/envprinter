package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	configFile = "./config.json"
)

func main() {
	fmt.Printf("Starting Environment Printer.......\n")

	stop := make(chan struct{})
	go dumpenvironment()

	if _, err := os.Stat(configFile); err == nil {
		go dumpconfigfile()
	} else {
		fmt.Printf("Couldn't find config file\n")
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown signal received, exiting...")
	close(stop)
}

func dumpenvironment() {
	fmt.Printf("\n")
	fmt.Printf("*******************************\n")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Printf("%s = %s\n", pair[0], pair[1])
	}
	time.Sleep(30 * time.Second)
}

func dumpconfigfile() {
	config, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error occured reading config file: %v", err)
		return
	}
	fmt.Print(string(config))
	fmt.Print("\n")

	time.Sleep(30 * time.Second)
}
