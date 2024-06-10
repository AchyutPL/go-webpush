package main

import (
	"fmt"
	"go-webpush/internal/server"
)

func main() {

	err := server.NewServer()

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
