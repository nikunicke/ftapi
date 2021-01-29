package main

import (
	"fmt"
	"log"

	"github.com/nikunicke/ftapi"
)

func main() {
	fmt.Println("Hello ft_api")
	_, err := ftapi.Client(ftapi.ClientCredentials)
	if err != nil {
		log.Fatalf("Failed to initiate client: %v", err)
	}
}
