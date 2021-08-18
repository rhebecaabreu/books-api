package dto

type BookUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Author      string `json:"author"`
	Description string `json:"description" form:"description" binding:"required"`
}

type BookCreateDTO struct {
	Author      string `json:"author"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}
