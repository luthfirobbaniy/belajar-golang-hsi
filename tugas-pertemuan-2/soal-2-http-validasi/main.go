package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Query struct {
	email string
	age   int
}

type ResponseError struct {
	Error string
}

type ResponseSuccess struct {
	Status string
}

var errInputInvalid = errors.New("email kosong atau umur kurang dari 18")

func logRequest(r *http.Request) {
	log.WithFields(log.Fields{
		"method": r.Method,
		"url":    r.URL.String(),
	}).Info()
}

func validateQuery(query Query) error {
	if query.email == "" || query.age < 18 {
		err := fmt.Errorf("%w", errInputInvalid)
		return err
	}

	return nil
}

func validate(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	query := r.URL.Query()

	email := query.Get("email")
	age, _ := strconv.Atoi(query.Get("age"))

	err := validateQuery(Query{email, age})

	w.Header().Add("content-type", "application/json")

	if errors.Is(err, errInputInvalid) {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(
			ResponseError{
				Error: err.Error(),
			},
		)

		return
	}

	json.NewEncoder(w).Encode(
		ResponseSuccess{
			Status: "ok",
		},
	)
}

func main() {
	http.HandleFunc("/validate", validate)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
