package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"smart-api/internal/auth"
)

func ChargesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	charges, err := auth.LoadCharges()
	if err != nil {
		fmt.Println("Failed to load charges:", err)
		http.Error(w, "Failed to load charges", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(charges)
}

func CreateChargeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newCharge auth.Charge
	if err := json.NewDecoder(r.Body).Decode(&newCharge); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	charges, err := auth.LoadCharges()
	if err != nil {
		http.Error(w, "Не удалось загрузить начисления", http.StatusInternalServerError)
		return
	}

	newCharge.ID = len(charges) + 1
	charges = append(charges, newCharge)

	err = auth.SaveCharges(charges)
	if err != nil {
		http.Error(w, "Не удалось сохранить начисления", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newCharge)
}