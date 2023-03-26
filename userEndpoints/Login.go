package userEndpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	types "user-management/models"

	_ "github.com/go-sql-driver/mysql"
)

// Login with UserName and Password
func LoginRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	//Get the Data
	var loginData types.User
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, "Login Decoding Error", http.StatusInternalServerError)
		return
	}
	fmt.Printf(loginData.Username)
	//Validate the Data
	query, err := db.Query(`select count(*), PWord from users.users where UserName = ? Group By UserName`, loginData.Username)
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
	if loginData.PWord != passw {
		http.Error(w, "Incorrect Password", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Login Successful")
	w.WriteHeader(http.StatusCreated)
}
