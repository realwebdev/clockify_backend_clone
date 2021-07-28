package main

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/realwebdev/Bilal/clockify3/models"
)

func main() {

	// Database connection string
	//give

	log.Println("Clockify")
	var op string
	fmt.Println(`Signup = a        SignIn = b`)
	fmt.Scanln(&op)

	var (
		id       uint
		username string
		email    string
		password string
	)
	switch {

	case op == "a":

		fmt.Println("Enter ID")

		fmt.Scanln(id)
		fmt.Println("Enter username")

		fmt.Scanln(username)
		fmt.Println("Enter unique email")

		fmt.Scanln(email)
		fmt.Println("Enter password")

		fmt.Scanln(password)

		// p1:= models.SignUp(1, "haseeb", "haseeb@email1.com", "password1")

		models.SignUp(models.User{User_id: id, Username: username, Email: email, Password: password})

	case op == "b":
		fmt.Println("Enter  email")

		fmt.Scanln(email)
		fmt.Println("Enter password")

		fmt.Scanln(password)
		models.SignIn(email, password)

	}

}
