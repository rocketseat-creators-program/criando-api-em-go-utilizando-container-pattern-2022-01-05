package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GuilhermeCaruso/expertscrud/src/handlers"
	"github.com/GuilhermeCaruso/expertscrud/src/repositories"
	"github.com/GuilhermeCaruso/expertscrud/src/services"
	"github.com/GuilhermeCaruso/expertscrud/src/structs"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	server := fiber.New()

	dialector := postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"<HOST",
		"<USER",
		"<PASSWORD>",
		"<DB>",
		"<PORT",
	))

	db, err := gorm.Open(dialector)

	if err != nil {
		log.Fatalf("banco nao conectou. Err=%v", err.Error())
	}

	db.AutoMigrate(&structs.Installment{})

	repositoryContainer := repositories.GetRepositories(db)
	servicesContainer := services.GetServices(repositoryContainer)
	handlers.NewHandlerContainer(server, servicesContainer)

	server.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("pong :)")
	})

	if err := server.Listen(":3001"); err != nil {
		log.Fatalf("aplicacao nao subiu. Err=%v", err.Error())
	}
}
