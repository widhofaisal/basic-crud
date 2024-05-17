package router

import (
	"basic-crud/controller"
	"basic-crud/middleware"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()

	e.Use(mid.CORS())

	e.Use(middleware.MiddlewareLogging)

	eTodos := e.Group("/todolists/api/v1")
	eTodos.GET("/hello", controller.Hello_world)
	eTodos.GET("", controller.Get_all_list)
	eTodos.POST("", controller.Add_list)
	eTodos.PUT("/:id", controller.Update_list_by_id)
	eTodos.DELETE("/:id", controller.Delete_list_by_id)

	eBooks := e.Group("/books/api/v1")
	eBooks.GET("", controller.Get_all_book)
	eBooks.POST("", controller.Add_book)
	eBooks.PUT("/:id", controller.Update_book_by_id)
	eBooks.DELETE("/:id", controller.Delete_book_by_id)

	return e
}
