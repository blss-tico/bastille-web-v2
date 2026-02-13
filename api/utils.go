package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondOkWithJSONUtil(w http.ResponseWriter, payload string) {
	log.Println("respondOkWithJSONUtil")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	if payload != "" {
		response := map[string]string{"msg": payload}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]string{"msg": "command executed"}
	json.NewEncoder(w).Encode(response)
}

func respondErrorWithJSONUtil(w http.ResponseWriter, code int, payload string) {
	log.Println("RespondErrorWithJSONUtil")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)

	response := map[string]string{"msg": "with errors", "err": payload}
	json.NewEncoder(w).Encode(response)
}
