package main

import (
	"github.com/SekiguchiKai/clean-architecture-with-go/server/infra/router"
)

func main() {
	if err := router.G.Run(":8080"); err != nil {
		panic(err.Error())
	}
}
