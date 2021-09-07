package handler

import (
	"log"
	"net/http"
)

type sleepHandler struct{}

func SleepHandler() *sleepHandler {
	return &sleepHandler{}
}

func (sleepHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//log.Println("before sleep")
	//sleep
	//time.Sleep(120 * time.Second)
	//log.Println("after sleep")

	log.Println("Receiving request:" + r.RequestURI + " method:" + r.Method)
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}