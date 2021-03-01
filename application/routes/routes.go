package routes

import (
	_ "github.com/jailtonjunior94/go-challenge-appointments/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// @title Go Challenge Appointments
// @version 1.0
// @description Go Challenge Appointments
// @contact.name API Support
// @contact.email jailton.junior94@outlook.com
// @BasePath /
func RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Get("/me", me)

	RegisterUserRouter(v1)
	RegisterExamRouter(v1)
	RegisterAppointmentRouter(v1)

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

}

// Me godoc
// @Summary Me
// @Description Me
// @Accept  json
// @Produce  json
// @Tags Me
// @Success 200 {object} Me
// @Failure 400,404 {object} HttpError
// @Failure 500 {object} HttpError
// @Failure default {object} HttpError
// @Router /api/v1/me [get]
func me(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(Me{Message: "I'm a project, Go Challenge Appointments"})

}

type Me struct {
	Message string `json:"message"`
}

type HttpError struct {
	Status  string
	Message string
}
