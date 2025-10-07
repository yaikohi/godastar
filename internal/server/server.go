package server

import (
	"github.com/gofiber/fiber/v2"

	"godastar/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "godastar",
			AppName:      "godastar",
		}),

		db: database.New(),
	}

	return server
}
