package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/luthfi/inventory-pos-app/backend/config"
	"github.com/luthfi/inventory-pos-app/backend/handlers"
	"github.com/luthfi/inventory-pos-app/backend/middleware"
	"github.com/luthfi/inventory-pos-app/backend/repositories"
	"github.com/luthfi/inventory-pos-app/backend/services"
)

func main() {
	config.ConnectDB()
	app := fiber.New()

	userRepo := repositories.NewUserRepository(config.DB)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	productRoutes := app.Group("/products")
	productRoutes.Use(middleware.AuthRequired(userRepo))

	productRoutes.Get("/", handlers.GetProducts)
	productRoutes.Post("/", handlers.CreateProduct)
	productRoutes.Put("/:id", handlers.UpdateProduct)
	productRoutes.Delete("/:id", handlers.DeleteProduct)

	app.Post("/logout", middleware.AuthRequired(userRepo), authHandler.Logout)

	log.Println("Server berjalan di :8080")
	log.Fatal(app.Listen(":8080"))

	select {}
}
