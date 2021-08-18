package controllers

import (
	"books-api/dto"
	"books-api/entity"
	"books-api/helpers"
	"books-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type bookController struct {
	bookService services.BookService
}

func NewBookController(bookServ services.BookService) BookController {
	return &bookController{
		bookService: bookServ,
	}
}

func (c *bookController) All(context *gin.Context) {
	var books []entity.Book = c.bookService.All()
	res := helpers.BuildResponse(true, "OK", books)
	context.JSON(http.StatusOK, res)
}

func (c *bookController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helpers.BuildErrorResponse("No param id was found", err.Error(), helpers.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var book entity.Book = c.bookService.FindByID(id)
	if (book == entity.Book{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", book)
		context.JSON(http.StatusOK, res)
	}
}

func (c *bookController) Insert(context *gin.Context) {
	var bookCreateDTO dto.BookCreateDTO
	errDTO := context.ShouldBind(&bookCreateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {

		result := c.bookService.Insert(bookCreateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *bookController) Update(context *gin.Context) {
	var bookUpdateDTO dto.BookUpdateDTO
	errDTO := context.ShouldBind(&bookUpdateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
	context.JSON(http.StatusForbidden, response)

}

func (c *bookController) Delete(context *gin.Context) {
	var book entity.Book
	// id, err := strconv.ParseUint(context.Param("id"), 10, 32)
	id, err := strconv.Atoi(context.Param("id"))
	// id, err := uint(context.Param("id"))
	if err != nil {
		response := helpers.BuildErrorResponse("Failed tou get id", "No param id were found", helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	book.ID = uint(id)

	c.bookService.Delete(book)
	res := helpers.BuildResponse(true, "Deleted", helpers.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
