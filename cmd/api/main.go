package main

import (
	_ "WebApp/docs"
	"WebApp/internal/config"
	"WebApp/internal/repository/mongo"
	"WebApp/internal/service"
	"WebApp/internal/transport/rest/handler"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
	"time"
)

func main() {
	if err := SetupViper(); err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New()

	config.SetupSwagger(app)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	mongoDatabase, err := config.SetupMongoDatabase(ctx, cancel)

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := mongo.NewUserRepository(mongoDatabase.Collection("go_study_wa"))
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	userHandler.InitRoutes(app)

	port := viper.GetString("http.port")
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}

func SetupViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil

}
