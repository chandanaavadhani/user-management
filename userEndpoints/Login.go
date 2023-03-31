package userEndpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chandanaavadhani/usermanagement/models"
	"github.com/chandanaavadhani/usermanagement/userValidations"

	_ "github.com/go-sql-driver/mysql"
)

// Login with UserName and Password
func LoginRequest(w http.ResponseWriter, r *http.Request) {

	//Validate the Method
	if r.Method != "POST" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//Get the Data
	var loginData models.User
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, "Login Decoding Error", http.StatusInternalServerError)
		return
	}

	//Validate the Data
	StatusCode, err := userValidations.LoginValidation(loginData)
	if err != nil {
		http.Error(w, err.Error(), StatusCode)
		return
	}
	fmt.Fprintf(w, "Login Successful")
	w.WriteHeader(http.StatusCreated)
}
