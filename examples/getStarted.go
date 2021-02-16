package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nikunicke/ftapi"
)

func main() {
	// Create a new service client. Choose AuthorizationCode or
	// ClientCredentials grant type.
	service, err := ftapi.NewService(ftapi.AuthorizationCode)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	// Get your user data
	me, err := service.Me()
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	fmt.Printf("%d - %s %s\n", me.ID, me.FirstName, me.LastName)

	// RAW response data from me.ServerResponse.Header | .Body | .StatusCode.
	// ServerResponse is included in all direct responses.
	fmt.Println(me.ServerResponse.StatusCode)

	// Get user by ID
	usersService := ftapi.Users(service)
	aUser, err := usersService.Get("59634").Do()
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	// Get user levels (piscine/42/...)
	for _, cursus := range aUser.CursusUsers {
		fmt.Println(cursus.Cursus.Name, cursus.Level)
	}

	// Get page of users from Hive Helsinki
	usersListCaller := usersService.List()
	page, err := usersListCaller.P("campus_id", "13").Do()
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	// Get following pages
	for len(page.Users) > 0 {
		fmt.Println(len(page.Users))
		page, err = usersListCaller.PageToken(page.NextPage).Do()
		if err != nil {
			log.Fatalf("Request failed: %v", err)
		}
		time.Sleep(1 * time.Second) // Rate-limit
	}
}
