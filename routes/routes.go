package routes

import (
	"sentinal-contactform/controllers"

	"github.com/labstack/echo/v4"
)

func ConactRoutes(group *echo.Group) {

	group.POST("/submit", controllers.ContactFormController)

}
