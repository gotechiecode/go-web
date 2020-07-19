package main

import (
	"fmt"
	"github.com/gotechiecode/go-web/go-client-server/go-server/service"
)

var appName = "accountservice"

func main()  {
	fmt.Println("Starting %v\n", appName)
	service.StartWebServer("6767")
}
