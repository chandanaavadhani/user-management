package userEndpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chandanaavadhani/usermanagement/models"
	"github.com/chandanaavadhani/usermanagement/userDB"
	"github.com/chandanaavadhani/usermanagement/userValidations"
	_ "github.com/go-sql-driver/mysql"
)

// SignUp Request
func SignUpRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//Get the Data
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Decoding Error", http.StatusInternalServerError)
		return
	}

	//Validate the data
	StatusCode, err := userValidations.SignUpValidation(user)
	if err != nil {
		http.Error(w, err.Error(), StatusCode)
		return
	}

	//Insert the details
	userDB.Insertdetails(user)
	fmt.Fprintf(w, "Successfully Created new user: %s", user.Username)
	w.WriteHeader(http.StatusCreated)
}
