package main

import (
	_ "WebApp/docs"
	"WebApp/internal/repository/memory"
	"WebApp/internal/service"
	"WebApp/internal/transport/rest/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Panic(err)
	}

	port := viper.Get("PORT").(string)

	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	userRepository := memory.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	userHandler.InitRoutes(app)
	if err := app.Listen(":" + port); err != nil {
		log.Panic(err)
	}
}
