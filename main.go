package main

import (
	"books-api/config"
	"books-api/controllers"
	"books-api/repository"
	"books-api/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	db := config.InitializeDatabase(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	bookRepository := repository.NewBookRepository(db)
	bookService := services.NewBookService(bookRepository)
	bookController := controllers.NewBookController(bookService)

	routes := gin.Default()

	bookRoutes := routes.Group("api/books")
	{
		bookRoutes.GET("/", bookController.All)
		bookRoutes.GET("/:id", bookController.FindByID)
		bookRoutes.POST("/", bookController.Insert)
		bookRoutes.PUT("/:id", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}

	routes.Run()
}
