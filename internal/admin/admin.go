package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"internal/user"
)

// Функція для перегляду та редагування профілю адміністратора
func ViewAndEditProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Перегляд профілю
		profile := user.GetProfile("admin")
		json.NewEncoder(w).Encode(profile)
	} else if r.Method == http.MethodPost {
		// Редагування профілю
		var profile user.Profile
		err := json.NewDecoder(r.Body).Decode(&profile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user.UpdateProfile("admin", profile)
		fmt.Fprintf(w, "Profile updated successfully")
	}
}

// Функція для перегляду списку покупців
func ViewCustomers(w http.ResponseWriter, r *http.Request) {
	customers := user.GetAllCustomers()
	json.NewEncoder(w).Encode(customers)
}

// Функція для блокування покупців
func BlockCustomer(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CustomerID string `json:"customer_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.BlockCustomer(req.CustomerID)
	fmt.Fprintf(w, "Customer blocked successfully")
}
