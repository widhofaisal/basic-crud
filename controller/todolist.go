package controller

import (
	"basic-crud/config"
	"basic-crud/model"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

// ENdpoint 0: Hello_world
func Hello_world(c echo.Context) error {
	return c.JSON(http.StatusOK, model.ResponseError{
		StatusCode:   200,
		ErrorMessage: "Hello world!!! All right reserved. TEST_NO: 001",
	})
}

// Endpoint 1 : Get_all_list
// TODO:
func Get_all_list(c echo.Context) error {
	todolists := []model.ToDoList{}
	if err_find := config.DB.Find(&todolists).Error; err_find != nil {
		log.Print(color.RedString(err_find.Error()))
		return c.JSON(http.StatusInternalServerError, model.ResponseError{
			StatusCode:   500,
			ErrorMessage: err_find.Error(),
		})
	}

	responseSuccess := model.ResponseSuccessToDoListAll{
		StatusCode:   200,
		ErrorMessage: "success to get all to do list",
		Data:         todolists,
	}

	return c.JSON(http.StatusOK, responseSuccess)
}

// Endpoint 2 : Add_list
// TODO:
func Add_list(c echo.Context) error {
	// DEFINE: struct only for binding
	type bindJson struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		Priority    string `json:"priority"`
	}

	// BIND: reuqest body json
	bindingTodos := bindJson{}
	if err_bind := c.Bind(&bindingTodos); err_bind != nil {
		log.Print(color.RedString(err_bind.Error()))
		return c.JSON((http.StatusBadRequest), map[string]interface{}{
			"status":  400,
			"message": "bad request",
		})
	}

	// VALIDATION: check if request body empty
	if bindingTodos.Title == "" || bindingTodos.Status == "" || bindingTodos.Priority == "" {
		log.Print(color.RedString("there is an empty request body"))
		return c.JSON((http.StatusBadRequest), map[string]interface{}{
			"status":  400,
			"message": "bad request, there is an empty request body",
		})
	}

	// insert data
	var inputedList model.ToDoList
	inputedList.Title = bindingTodos.Title
	inputedList.Status = bindingTodos.Status
	inputedList.Priority = bindingTodos.Priority
	if err_save := config.DB.Save(&inputedList).Error; err_save != nil {
		log.Print(color.RedString(err_save.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  500,
			"message": "internal server error",
		})
	}

	responseSuccess := model.ResponseSuccessToDoListOne{
		StatusCode:   201,
		ErrorMessage: "success to add list",
		Data:         inputedList,
	}

	return c.JSON(http.StatusCreated, responseSuccess)
}

// // Endpoint 3 : Get_material_by_material_code
// // TODO:
// func Get_material_by_material_code(c echo.Context) error {
// 	token := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
// 	role, _ := utils.Get_role_from_token(token)
// 	if role != "admin" {
// 		log.Print(color.RedString("only admins are allowed"))
// 		return c.JSON((http.StatusUnauthorized), map[string]interface{}{
// 			"status":  401,
// 			"message": "unauthorized, only admins are allowed",
// 		})
// 	}

// 	material_code := c.Param("material_code")

// 	material := model.Material{}
// 	if err_first := config.DB.Where("material_code=?", material_code).First(&material).Error; err_first != nil {
// 		log.Print(color.RedString(err_first.Error()))
// 		return c.JSON((http.StatusBadRequest), map[string]interface{}{
// 			"status":  400,
// 			"message": "bad request, material_code = " + material_code + " not found",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"status":  200,
// 		"message": "success to get material by material_code = " + material_code,
// 		"data":    material,
// 	})
// }

// Endpoint 3 : Update_list_by_id
// TODO:
func Update_list_by_id(c echo.Context) error {
	listId := c.Param("id")
	var todos model.ToDoList

	// check is id exists
	if err_first := config.DB.Where("id=?", listId).First(&todos).Error; err_first != nil {
		log.Print(color.RedString(err_first.Error()))
		return c.JSON((http.StatusBadRequest), map[string]interface{}{
			"status":  400,
			"message": "bad request",
		})
	}

	// define struct only for binding
	type bindJson struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		Priority    string `json:"priority"`
	}

	// binding
	bindingTodos := bindJson{}
	if err_bind := c.Bind(&bindingTodos); err_bind != nil {
		log.Print(color.RedString(err_bind.Error()))
		return c.JSON((http.StatusBadRequest), map[string]interface{}{
			"status":  400,
			"message": "bad request",
		})
	}

	// VALIDATION: check if request body empty
	if bindingTodos.Title == "" || bindingTodos.Status == "" || bindingTodos.Priority == "" {
		log.Print(color.RedString("there is an empty request body"))
		return c.JSON((http.StatusBadRequest), map[string]interface{}{
			"status":  400,
			"message": "bad request, there is an empty request body",
		})
	}

	// save
	todos.Title = bindingTodos.Title
	todos.Description = bindingTodos.Description
	todos.Status = bindingTodos.Status
	todos.Priority = bindingTodos.Priority
	if err_save := config.DB.Save(&todos).Error; err_save != nil {
		log.Print(color.RedString(err_save.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  500,
			"message": "internal server error",
		})
	}

	responseSuccess := model.ResponseSuccessToDoListOne{
		StatusCode:   201,
		ErrorMessage: "success to update list by id",
		Data:         todos,
	}

	return c.JSON(http.StatusCreated, responseSuccess)
}

// Endpoint 4 : Delete_list_by_id
// TODO:
func Delete_list_by_id(c echo.Context) error {
	listId := c.Param("id")
	var todos model.ToDoList

	// check is id exists
	if err_first := config.DB.Where("id=?", listId).First(&todos).Error; err_first != nil {
		log.Print(color.RedString(err_first.Error()))
		return c.JSON((http.StatusBadRequest), map[string]interface{}{
			"status":  400,
			"message": "bad request",
		})
	}

	// delete
	if err_delete := config.DB.Where("id=?", listId).Delete(&todos).Error; err_delete != nil {
		log.Print(color.RedString(err_delete.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  500,
			"message": "internal server error",
		})
	}

	responseSuccess := model.ResponseSuccessToDoListOne{
		StatusCode:   202,
		ErrorMessage: "success to delete list by id",
		Data:         todos,
	}

	return c.JSON(http.StatusCreated, responseSuccess)
}
