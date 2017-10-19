package main

import (
	"net/http"
	"sync"
	"log"
	"github.com/manachyn/go-examples/microservices/microservice-eureka-client/service"
	"github.com/manachyn/go-examples/microservices/microservice-eureka-client/eureka"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Graceful shutdown on Ctrl+C or kill
	handleSigterm()

	// Starts HTTP service  (async)
	go startWebServer()

	// Performs Eureka registration
	eureka.Register()

	// Performs Eureka heartbeating (async)
	go eureka.StartHeartbeat()

	// Use a WaitGroup to block main() exit
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func handleSigterm() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		eureka.Deregister()
		os.Exit(1)
	}()
}

func startWebServer() {
	router := service.NewRouter()
	log.Println("Starting HTTP service at 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port 8080")
		log.Println("Error: " + err.Error())
	}
}
