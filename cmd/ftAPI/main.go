package main

import (
	"fmt"
	"log"

	"github.com/nikunicke/ftapi"
)

// 3647

func main() {
	fmt.Println("Hello ft_api")
	s, err := ftapi.NewService(ftapi.AuthorizationCode)
	if err != nil {
		log.Fatalf("Failed to authenticate client: %v", err)
	}
	fmt.Println("Authentication succesful")
	me, err := s.Me()
	fmt.Println(me.FirstName, me.LastName)
	// users := ftapi.Users(s)
	// me, err := users.Get("59634").Do()
	// if err != nil {
	// 	log.Fatalf("Failed to get me: %v", err)
	// }
	// fmt.Println(me.FirstName)
	fmt.Println("DONE")
}
