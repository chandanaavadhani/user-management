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

// Update the Password
func UpdateRequest(w http.ResponseWriter, r *http.Request) {

	//Validate the Method
	if r.Method != "PUT" {
		fmt.Println(r.Method)
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//Get the Data
	var updateUser models.User
	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, "Updating Error", http.StatusInternalServerError)
		return
	}

	//Validate the data
	StatusCode, err := userValidations.UpdateValidation(updateUser)
	if err != nil {
		http.Error(w, err.Error(), StatusCode)
		return
	}

	//Update the Data
	err = userDB.Updatepassword(updateUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Updated the password for user: %s", updateUser.Username)
	w.WriteHeader(http.StatusAccepted)
}
