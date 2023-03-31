package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chandanaavadhani/usermanagement/userEndpoints"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	http.HandleFunc("/home", jwt.verifyJWT(userEndpoints.handlePage))
	http.HandleFunc("/SignUp", userEndpoints.SignUpRequest)
	http.HandleFunc("/Login", userEndpoints.LoginRequest)
	http.HandleFunc("/Update", userEndpoints.UpdateRequest)
	http.HandleFunc("/Delete", userEndpoints.DeleteRequest)
	log.Printf("Starting the server at the port 9999")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println("There was an error listening on port :9999", err)
	}
}
