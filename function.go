package ipo

import (
	"encoding/json"
	"log"
	"net/http"
)

func ListIpos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := GetIpoData()
	if err != nil {
		log.Fatal(err)
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
