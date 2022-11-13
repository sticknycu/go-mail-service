package main

import (
	"go-mail-service/service"
)

func main() {
	s := service.NewService()
	s.StartWebService()
}
