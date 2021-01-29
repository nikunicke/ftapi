package main

import (
	"fmt"
	"log"

	"github.com/nikunicke/ftAPI"
)

func main() {
	fmt.Println("Hello ft_api")
	_, err := ftAPI.Client(ftAPI.ClientCredentials)
	if err != nil {
		log.Fatalf("Failed to initiate client: %v", err)
	}
}
