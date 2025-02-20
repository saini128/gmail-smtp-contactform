package mailing

import (
	"errors"
	"fmt"
	"os"
)

var (
	SES_USERNAME string
	SES_PASSWORD string
	SES_HOST     string
)

func Init() error {

	SES_USERNAME = os.Getenv("GMAIL_USERNAME")
	SES_PASSWORD = os.Getenv("GMAIL_PASSWORD")
	SES_HOST = os.Getenv("GMAIL_HOST")

	if SES_USERNAME == "" || SES_PASSWORD == "" || SES_HOST == "" {
		fmt.Println("SES_USERNAME or SES_PASSWORD or SES_HOST environment variable not found!")
		return errors.New("SES_USERNAME or SES_PASSWORD or SES_HOST environment variable not found")
	}
	return nil
}
