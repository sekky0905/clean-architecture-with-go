package main

import (
	"fmt"

	"github.com/SekiguchiKai/clean-architecture-with-go/server/infra/router"
)

func main() {
	if err := router.G.Run(":8080"); err != nil {
		fmt.Println("~~~:")
		panic(err.Error())
	}
}
