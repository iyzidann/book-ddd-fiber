package http

import (
	domainBook "book-ddd/domain/book"
	infraBook "book-ddd/infrastructure/book"
	"book-ddd/interfaces/http/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *fiber.App {
	app := fiber.New()

	app.Use(cors.New())

	repo := infraBook.NewMySQLRepo(db)
	service := domainBook.NewService(repo)
	handler := handler.NewBookHandler(service)

	api := app.Group("/books")
	api.Get("/", handler.GetBooks)
	api.Get("/:id", handler.GetBook)
	api.Post("/", handler.CreateBook)
	api.Put("/:id", handler.UpdateBook)
	api.Delete("/:id", handler.DeleteBook)

	return app
}
