package handler

import (
	"WebApp/internal/core"
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	GetAll() []*core.User
	GetById(id int) *core.User
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (handler *UserHandler) InitRoutes(app *fiber.App) {
	app.Get("/users", handler.GetAll)
}

func (handler *UserHandler) GetAll(ctx *fiber.Ctx) error {
	users := handler.userService.GetAll()

	return ctx.Status(200).JSON(
		fiber.Map{
			"Users": users,
		})
}
