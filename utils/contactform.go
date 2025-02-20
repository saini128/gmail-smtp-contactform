package utils

import "sentinal-contactform/models"

func ContactFormEmail(msg models.ContactFormBody) string {

	message := "Name: " + msg.Name + "<br>" + "Email: " + msg.Email + "<br>" + "Subject: " + msg.Subject + "<br>" + "Message: " + msg.Message
	return message
}
