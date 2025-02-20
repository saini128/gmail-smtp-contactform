package mailing

import (
	"bytes"
	"sentinal-contactform/models"
	"sentinal-contactform/utils"

	"fmt"
	"log"
	"mime"
	"mime/multipart"
	"net/smtp"
	"net/textproto"
)

func SendContactInfoPage(msg models.ContactFormBody) error {
	auth := smtp.PlainAuth("", GMAIL_USERNAME, GMAIL_PASSWORD, GMAIL_HOST)
	formattedName := mime.QEncoding.Encode("utf-8", ContactName)
	formattedAddress := fmt.Sprintf("\"%s\" <%s>", formattedName, ContactEmail)
	subject := "Contact Form Submission - " + msg.Name
	body := utils.ContactFormEmail(msg)

	var b bytes.Buffer
	m := multipart.NewWriter(&b)

	// 1. Plain text part (fallback) - VERY IMPORTANT
	part, err := m.CreatePart(textproto.MIMEHeader{"Content-Type": []string{"text/plain; charset=UTF-8"}})
	if err != nil {
		return err
	}
	_, err = part.Write([]byte(body)) // Fallback text
	if err != nil {
		return err
	}

	// 2. HTML part
	part, err = m.CreatePart(textproto.MIMEHeader{"Content-Type": []string{"text/html; charset=UTF-8"}})
	if err != nil {
		return err
	}
	_, err = part.Write([]byte(body))
	if err != nil {
		return err
	}

	err = m.Close()
	if err != nil {
		return err
	}

	msgg := []byte("From: " + formattedAddress + "\r\n" +
		"To: " + ContactEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: multipart/alternative; boundary=\"" + m.Boundary() + "\"\r\n" +
		"\r\n" + b.String())

	err = smtp.SendMail(GMAIL_HOST+":587", auth, ContactEmail, []string{ContactEmail}, msgg)
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		return err
	}
	return nil
}

func SendForgotPasswordLink(email string) error {
	return nil
}
