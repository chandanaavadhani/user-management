package userEndpoints

import (
	"encoding/json"
	"net/http"

	"github.com/chandanaavadhani/usermanagement/models"
)

func handlePage(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message models.Message
	err := json.NewDecoder(request.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Decoding Error", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		http.Error(w, "Encoding Error", http.StatusInternalServerError)
		return
	}

}
