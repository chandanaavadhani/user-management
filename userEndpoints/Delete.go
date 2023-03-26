package userEndpoints

import (
	"fmt"
	"net/http"

	"github.com/chandanaavadhani/usermanagement/userDB"
	_ "github.com/go-sql-driver/mysql"
)

// Delete the Data
func DeleteRequest(w http.ResponseWriter, r *http.Request) {
	//Validate the Method
	if r.Method != "DELETE" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	db, err := userDB.Dbconnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Get the Query Params
	username := r.URL.Query().Get("username")
	query, err := db.Query(`select count(*), UserName from users.users where UserName = ? Group By UserName`, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Execute the Query
	var count int
	var uname string
	for query.Next() {
		if err := query.Scan(&count, &uname); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if count == 0 {
		//fmt.Println(count, username, uname)
		http.Error(w, "User not Found", http.StatusBadRequest)
		return
	}
	//Delete the data
	_, err = db.Query(`DELETE FROM users.users where UserName = ?`, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Deleted the username from DataBase")
	w.WriteHeader(http.StatusAccepted)
}
