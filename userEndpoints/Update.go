package userEndpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/chandanaavadhani/usermanagement/models"
	"github.com/chandanaavadhani/usermanagement/userDB"
	_ "github.com/go-sql-driver/mysql"
)

// Update the Password
func UpdateRequest(w http.ResponseWriter, r *http.Request) {
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
	db, err := userDB.Dbconnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Validate the Data
	query, err := db.Query(`select count(*), PWord from users.users where UserName = ? Group By UserName`, updateUser.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var count int
	var passw string
	for query.Next() {
		if err := query.Scan(&count, &passw); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if count == 0 {
		http.Error(w, "Username Not Found", http.StatusBadRequest)
		return
	}
	if passw != updateUser.OldPassword {
		http.Error(w, "Incorrect old Password", http.StatusBadRequest)
		fmt.Println(updateUser.OldPassword, updateUser.PWord)
		return
	}
	//fmt.Println(updateUser.OldPassword, updateUser.PWord)
	if len(updateUser.NewPassword) < 8 {
		http.Error(w, "Password is Not Valid", http.StatusBadRequest)
		return
	}
	if updateUser.ConfirmPassword != updateUser.NewPassword {
		http.Error(w, "Passwords don't match", http.StatusBadRequest)
		return
	}
	//Update the Data
	_, err = db.Query(`Update users.users set PWord = ? where UserName = ? `, updateUser.NewPassword, updateUser.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Updated the password for user: %s", updateUser.Username)
	w.WriteHeader(http.StatusAccepted)
}
