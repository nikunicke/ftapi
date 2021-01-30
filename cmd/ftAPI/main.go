package main

import (
	"fmt"
	"log"

	"github.com/nikunicke/ftapi"
	"github.com/nikunicke/ftapi/events"
)

// 3647

func main() {
	fmt.Println("Hello ft_api")
	client, err := ftapi.Client(ftapi.AuthorizationCode)
	if err != nil {
		log.Fatalf("Failed to authenticate client: %v", err)
	}
	events := events.Service(client)

	result, err := events.List().
		Q("cursus_id", []string{"1"}).
		Q("campus_id", []string{"13"}).
		Do()
	for _, event := range result.Events {
		fmt.Println(event.Name)
	}
	// result, err := events.Get("3647").Do()
	// if err != nil {
	// 	log.Fatalf("Failed to get event: %v", err)
	// }
	// fmt.Printf("%s\n", result.Name)
	fmt.Println("DONE")
}
