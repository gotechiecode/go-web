package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main()  {
	resp, err := http.Get("http://localhost:6767")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}
	io.Copy(os.Stdout, resp.Body)
}
