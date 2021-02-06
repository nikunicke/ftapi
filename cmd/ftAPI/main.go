package main

import (
	"fmt"
	"log"

	"github.com/nikunicke/ftapi"
)

func main() {
	fmt.Println("Hello ft_api")
	s, err := ftapi.NewService(ftapi.AuthorizationCode)
	if err != nil {
		log.Fatalf("Failed to authenticate client: %v", err)
	}
	fmt.Println("Authentication succesful")

	me, err := s.Me()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println("Hello,", me.FirstName, me.LastName)
	events := ftapi.Events(s)
	eventList, err := events.List().P("campus_id", []string{"13"}).Do()
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, event := range eventList.Events {
		fmt.Println(event.Name)
	}

	fmt.Println("DONE")
}
