package routes

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	v1.Get("/me", me)

	RegisterUserRouter(v1)
	RegisterExamRouter(v1)
	RegisterAppointmentRouter(v1)
}

func me(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "I'm a project, Go Challenge Appointments"})
}
