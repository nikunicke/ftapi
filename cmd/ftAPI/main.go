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

	eventListCall := events.List()
	page1, err := eventListCall.P("campus_id", []string{"13"}).Do()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(page1.Header.Get("Link"))
	fmt.Println("Next page:", page1.NextPage)
	page2, err := eventListCall.PageToken(page1.NextPage).Do()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(page2.Header.Get("Link"))

	fmt.Println("DONE")
}
