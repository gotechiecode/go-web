package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main()  {
	log.Println("My Server Started")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":6767", nil))
}

func handler(w http.ResponseWriter, r *http.Request)  {
	log.Println("Handler started")
	defer log.Println("handler ended")
	// making the server to slow down, so that we can see the results
	time.Sleep(5 * time.Second)
	fmt.Fprintln(w, "You are running a server")
}