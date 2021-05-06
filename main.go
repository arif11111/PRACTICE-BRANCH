package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type operands struct {
	Operand1 float64 `json:"operand1,omitempty"`
	Operand2 float64 `json:"operand2,omitempty"`
}

func add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var request operands
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
	}
	response := request.Operand1 + request.Operand2
	json.NewEncoder(w).Encode(response)
}

func subtract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var request operands
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
	}
	response := request.Operand1 - request.Operand2
	json.NewEncoder(w).Encode(response)
}

func multiply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var request operands
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
	}
	response := request.Operand1 * request.Operand2
	json.NewEncoder(w).Encode(response)
}

func divide(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var request operands
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
	}
	response := request.Operand1 / request.Operand2
	json.NewEncoder(w).Encode(response)
}

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/status", status).Methods("GET")
	router.HandleFunc("/add", add).Methods("POST")
	router.HandleFunc("/subtract", subtract).Methods("POST")
	router.HandleFunc("/multiply", multiply).Methods("POST")
        router.HandleFunc("/divide", divide).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
