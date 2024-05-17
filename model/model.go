package model

import (
	"gorm.io/gorm"
)

type ToDoList struct {
	gorm.Model  `json:"-"`
	ID          uint   `json:"id" form:"id"`
	Title       string `json:"title" form:"title" gorm:"not null"`
	Description string `json:"description" form:"description" gorm:"not null"`
	Status      string `json:"status" form:"status" gorm:"not null"`
	Priority    string `json:"priority" form:"priority" gorm:"not null"`
}

type Book struct {
	gorm.Model  `json:"-"`
	ID          uint   `json:"id" form:"id"`
	Title       string `json:"title" form:"title" gorm:"not null"`
	Description string `json:"description" form:"description" gorm:"not null"`
	Genre       string `json:"genre" form:"genre" gorm:"not null"`
	TotalPage   int    `json:"total_page" form:"total_page" gorm:"not null"`
}

type ResponseError struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"error_message"`
}

type ResponseSuccessToDoListAll struct {
	StatusCode   int        `json:"status"`
	ErrorMessage string     `json:"error_message"`
	Data         []ToDoList `json:"data`
}

type ResponseSuccessToDoListOne struct {
	StatusCode   int      `json:"status"`
	ErrorMessage string   `json:"error_message"`
	Data         ToDoList `json:"data`
}

type ResponseSuccessBookAll struct {
	StatusCode   int        `json:"status"`
	ErrorMessage string     `json:"error_message"`
	Data         []Book `json:"data`
}

type ResponseSuccessBookOne struct {
	StatusCode   int      `json:"status"`
	ErrorMessage string   `json:"error_message"`
	Data         Book `json:"data`
}
