package main

type Email struct {
	To        string `json:"To"`
	From      string `json:"From"`
	Date      string `json:"Date"`
	Subject   string `json:"Subject"`
	MessageID string `json:"Message-ID"`
}

type Emails []Email
