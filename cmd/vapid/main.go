package main

import (
	"fmt"

	"github.com/SherClockHolmes/webpush-go"
)

func GenerateVapidKey() {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		fmt.Println("Error generating VAPID keys:", err)
		return
	}

	fmt.Println("Private Key:", privateKey)
	fmt.Println("Public Key:", publicKey)
}

func main() {
	GenerateVapidKey()
}
