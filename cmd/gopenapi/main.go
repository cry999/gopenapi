package main

import (
	"log"

	"github.com/cry999/gopenapi/pkg/iface/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
