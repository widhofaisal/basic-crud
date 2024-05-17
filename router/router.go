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

	eApi := e.Group("/todolists/api/v1")
	eApi.GET("/hello", controller.Hello_world)
	eApi.GET("", controller.Get_all_list)
	eApi.POST("", controller.Add_list)
	eApi.PUT("/:id", controller.Update_list_by_id)
	eApi.DELETE("/:id", controller.Delete_list_by_id)

	
	return e
}

