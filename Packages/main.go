package main

import (
	"fmt"

	"github.com/fatih/color"
	auth "github.com/krishsingh120/go/auth"
	user "github.com/krishsingh120/go/user"
)

func main() {
	auth.LoginWithCredentials("krish@123", "282005@krish")

	session := auth.GetSession()
	fmt.Println(session)

	u := &user.UserModel{}


	email := "krish@example.com"

	u.UpdateDetails(&email, nil, nil)
	fmt.Println(u.Email)
	fmt.Println(u.Username)
	fmt.Println(u.Mobile)


	// third package use in go : color package

	color.Green("Hello Babe!!")
	color.Cyan(u.Email)



}
