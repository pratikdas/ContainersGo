package main

import (
	"fmt"
	"inventory/service"
)

func main() {
	fmt.Println("Starting api cserver")
	service.StartServer("3000")
}
