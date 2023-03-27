package userEndpoints

import (
	"fmt"
	"net/http"

	"github.com/chandanaavadhani/usermanagement/userDB"
	"github.com/chandanaavadhani/usermanagement/userValidations"
	_ "github.com/go-sql-driver/mysql"
)

// Delete the Data
func DeleteRequest(w http.ResponseWriter, r *http.Request) {

	//Validate the Method
	if r.Method != "DELETE" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//Get the Query Params
	username := r.URL.Query().Get("username")

	//Validate the data
	StatusCode, err := userValidations.DeleteUsername(username)
	if err != nil {
		http.Error(w, err.Error(), StatusCode)
		return
	}

	//Delete the data
	err = userDB.Deleteuser(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Deleted the username from DataBase")
	w.WriteHeader(http.StatusAccepted)
}
