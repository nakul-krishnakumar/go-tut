package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/nakul-krishnakumar/go-tut/auth"
	"github.com/nakul-krishnakumar/go-tut/user"
)

func main() {
	auth.LoginWithCreds("nakul", "123")
	u := user.User{
		Email: "nakul@gmail.com",
	}
	fmt.Println(u)

	godotenv.Write(map[string]string{"MY_NAME": "nakul"}, ".env")
}