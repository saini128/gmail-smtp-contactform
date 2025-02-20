package mailing

import (
	"errors"
	"fmt"
	"os"
)

var (
	GMAIL_USERNAME string
	GMAIL_PASSWORD string
	GMAIL_HOST     string
)

func Init() error {

	GMAIL_USERNAME = os.Getenv("GMAIL_USERNAME")
	GMAIL_PASSWORD = os.Getenv("GMAIL_PASSWORD")
	GMAIL_HOST = os.Getenv("GMAIL_HOST")

	if GMAIL_USERNAME == "" || GMAIL_PASSWORD == "" || GMAIL_HOST == "" {
		fmt.Println("GMAIL_USERNAME or GMAIL_PASSWORD or GMAIL_HOST environment variable not found!")
		return errors.New("GMAIL_USERNAME or GMAIL_PASSWORD or GMAIL_HOST environment variable not found")
	}
	return nil
}
