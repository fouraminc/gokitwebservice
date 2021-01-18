package main

import (
	"context"
	"flag"
	"fmt"
	"gokitwebservice"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)

	flag.Parse()
	ctx := context.Background()
	srv := gokitwebservice.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	endpoints := gokitwebservice.Endpoints{
		GetEndPoint: gokitwebservice.MakeGetEndpoint(srv),
		StatusEndPoint: gokitwebservice.MakeStatusEndpoint(srv),
		ValidateEndpoint: gokitwebservice.MakeValidateEndpoint(srv),
	}

	go func() {
		log.Println("gokitwebservice is listening on port:", *httpAddr)
		handler := gokitwebservice.NewHttpServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
