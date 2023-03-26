package userEndpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	types "user-management/models"

	_ "github.com/go-sql-driver/mysql"
)

// SignUp Request
func SignUpRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	//Get the Data
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Decoding Error", http.StatusInternalServerError)
		return
	}
	fmt.Println(user.Fullname)
	//Validate the data
	if user.Fullname == " " || user.Fullname == "" {
		http.Error(w, "FullName is Required", http.StatusBadRequest)
		return
	}
	if user.Username == " " || user.Username == "" {
		http.Error(w, "UserName is Required", http.StatusBadRequest)
		return
	}
	if user.PWord == " " || user.PWord == "" {
		http.Error(w, "Password is Required", http.StatusBadRequest)
		return
	}
	if len(user.PWord) < 8 {
		http.Error(w, "Password is Not Valid", http.StatusBadRequest)
		return
	}
	if user.ConfirmPassword != user.PWord {
		http.Error(w, "Passwords don't match", http.StatusBadRequest)
		return
	}
	//Execute the query
	query, err := db.Prepare(`INSERT INTO users VALUES(?,?,?,?,?)`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = query.Exec(user.Fullname, user.Contactno, user.Email, user.Username, user.PWord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Successfully Created new user: %s", user.Username)
	w.WriteHeader(http.StatusCreated)

}
