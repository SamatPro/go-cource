package handler

import (
	"WebApp/internal/core"
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserService interface {
	GetById(ctx context.Context, id string) (*core.User, error)
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (handler *UserHandler) InitRoutes(app *fiber.App) {
	app.Get("/users/:userId", handler.GetById)
}

func (handler *UserHandler) GetById(ctx *fiber.Ctx) error {
	user, err := handler.userService.GetById(ctx.UserContext(), ctx.Params("userId"))

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			fiber.Map{
				"error": err.Error(),
			})
	}

	return ctx.Status(http.StatusOK).JSON(
		fiber.Map{
			"User": user,
		})
}
