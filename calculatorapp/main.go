package main

import (
	"fmt"
	"log"
	"github.com/lazycruise-sys/greeting"
	"github.com/lazycruise-sys/keyboard"
)

func main() {
	greeting.Hello()
	fmt.Println("Enter a Fahrenheit Temperature: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celcuis := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f degrees celcuis", celcuis)
}
