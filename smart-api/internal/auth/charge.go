package auth

import (
	"encoding/json"
	"os"
)

const chargeFile = "./data/charges.json"

func LoadCharges() ([]Charge, error) {
	file, err := os.Open(chargeFile)
	if err != nil {
		return []Charge{}, err
	}
	defer file.Close()

	var charges []Charge
	err = json.NewDecoder(file).Decode(&charges)
	if err != nil {
		return []Charge{}, err
	}
	return charges, err
}

func SaveCharges(charges []Charge) error {
	file, err := os.Create(chargeFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(charges)
}

func CreateCharge(newCharge Charge) error {
	charges, err := LoadCharges()
	if err != nil {
		return err
	}
	
	newCharge.ID = len(charges) + 1
	charges = append(charges, newCharge)
	return SaveCharges(charges)
}