package controllers

import (
	"net/http"
	"sentinal-contactform/mailing"
	"sentinal-contactform/models"

	"github.com/labstack/echo/v4"
)

func ContactFormController(c echo.Context) error {
	var reqBody models.ContactFormBody
	if err := c.Bind(&reqBody); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}
	err := mailing.SendContactInfoPage(reqBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to send email",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Message Sent Successfully"})
}
