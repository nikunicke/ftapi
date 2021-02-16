package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nikunicke/ftapi"
)

func main() {
	fmt.Println("Hello ft_api")
	service, err := ftapi.NewService(ftapi.AuthorizationCode)
	if err != nil {
		log.Fatalf("Failed to authenticate client: %v", err)
	}
	fmt.Println("Authentication succesful")

	me, err := service.Me()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println("Hello,", me.FirstName, me.LastName)

	eventsService := ftapi.Events(service)
	eventsListCaller := eventsService.List()
	page1, err := eventsListCaller.P("campus_id", "13").Do()
	if err != nil {
		log.Fatalf("%v", err)
	}
	respJSON, err := json.MarshalIndent(page1.ServerResponse, "", "\t")
	fmt.Println(string(respJSON))
	fmt.Println("DONE")
}
