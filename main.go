package main

import (
	"http-server/config"
	"http-server/handler"
	"log"
	"net/http"
	"time"
)

func main() {
	configFile, err := config.ParseConfigFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:         configFile.Listen,
		ReadTimeout:  configFile.ReadTimeout * time.Second,
		WriteTimeout: configFile.WriteTimeout * time.Second,
	}

	http.Handle("/", handler.SleepHandler())

	log.Println("Listen on", configFile.Listen)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error: cannot listen on %s: %v", configFile.Listen, err)
	}
}