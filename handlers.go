package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
)

func EmailIndex(w http.ResponseWriter, r *http.Request) {
	// get request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading request body: %v", err)
		http.Error(w, "can't read request body", http.StatusBadRequest)
		return
	}

	// convert request body to mail object
	emailBody, err := mail.ReadMessage(bytes.NewBufferString(string(body)))
	if err != nil {
		log.Printf("error reading request body: %v", err)
		http.Error(w, "can't read request body", http.StatusBadRequest)
		return
	}

	header := emailBody.Header

	// create email from strut
	email := Email{
		To:        header.Get("To"),
		From:      header.Get("From"),
		Date:      header.Get("Date"),
		Subject:   header.Get("Subject"),
		MessageID: header.Get("Message-ID"),
	}

	// set response mime type and return json response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(email)
}

func HealthcheckIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "healthy")
	w.WriteHeader(http.StatusOK)
}
