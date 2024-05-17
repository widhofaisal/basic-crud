package controller

import (
	"basic-crud/config"
	"basic-crud/model"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

// Endpoint 1 : Get_all_book
// TODO:
func Get_all_book(c echo.Context) error {
	books := []model.Book{}
	if err_find := config.DB.Find(&books).Error; err_find != nil {
		log.Print(color.RedString(err_find.Error()))
		return c.JSON(http.StatusInternalServerError, model.ResponseError{
			StatusCode:   500,
			ErrorMessage: err_find.Error(),
		})
	}

	responseSuccess := model.ResponseSuccessBookAll{
		StatusCode:   200,
		ErrorMessage: "success to get all book",
		Data:         books,
	}

	return c.JSON(http.StatusOK, responseSuccess)
}

// Endpoint 2 : Add_book
// TODO:
func Add_book(c echo.Context) error {
	// DEFINE: struct only for binding
	type bindJson struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Genre       string `json:"genre"`
		TotalPage   int    `json:"total_page"`
	}

	// BIND: request body json
	bindingBook := bindJson{}
	if err_bind := c.Bind(&bindingBook); err_bind != nil {
		log.Print(color.RedString(err_bind.Error()))
		return c.JSON((http.StatusBadRequest), map[string]interface{}{
			"status":  400,
			"message": "bad request",
		})
	}

	// VALIDATION: check if request body empty
	if bindingBook.Title == "" || bindingBook.Description == "" || bindingBook.Genre == "" || bindingBook.TotalPage == 0 {
		log.Print(color.RedString("there is an empty request body"))
		return c.JSON((http.StatusBadRequest), map[string]interface{}{
			"status":  400,
			"message": "bad request, there is an empty request body",
		})
	}

	// insert data
	var inputedBook model.Book
	inputedBook.Title = bindingBook.Title
	inputedBook.Description = bindingBook.Description
	inputedBook.Genre = bindingBook.Genre
	inputedBook.TotalPage = bindingBook.TotalPage
	if err_save := config.DB.Save(&inputedBook).Error; err_save != nil {
		log.Print(color.RedString(err_save.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  500,
			"message": "internal server error",
		})
	}

	responseSuccess := model.ResponseSuccessBookOne{
		StatusCode:   201,
		ErrorMessage: "success to add book",
		Data:         inputedBook,
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

// Endpoint 3 : Update_book_by_id
// TODO:
func Update_book_by_id(c echo.Context) error {
	bookId := c.Param("id")
	var theBook model.Book

	// check is id exists
	if err_first := config.DB.Where("id=?", bookId).First(&theBook).Error; err_first != nil {
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
		Genre       string `json:"genre"`
		TotalPage   int    `json:"total_page"`
	}

	// BIND: request body json
	bindingBook := bindJson{}
	if err_bind := c.Bind(&bindingBook); err_bind != nil {
		log.Print(color.RedString(err_bind.Error()))
		return c.JSON((http.StatusBadRequest), map[string]interface{}{
			"status":  400,
			"message": "bad request",
		})
	}

	// VALIDATION: check if request body empty
	if bindingBook.Title == "" || bindingBook.Description == "" || bindingBook.Genre == "" || bindingBook.TotalPage == 0 {
		log.Print(color.RedString("there is an empty request body"))
		return c.JSON((http.StatusBadRequest), map[string]interface{}{
			"status":  400,
			"message": "bad request, there is an empty request body",
		})
	}

	// save
	theBook.Title = bindingBook.Title
	theBook.Description = bindingBook.Description
	theBook.Genre = bindingBook.Genre
	theBook.TotalPage = bindingBook.TotalPage
	if err_save := config.DB.Save(&theBook).Error; err_save != nil {
		log.Print(color.RedString(err_save.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  500,
			"message": "internal server error",
		})
	}

	responseSuccess := model.ResponseSuccessBookOne{
		StatusCode:   201,
		ErrorMessage: "success to update book by id",
		Data:         theBook,
	}

	return c.JSON(http.StatusCreated, responseSuccess)
}

// Endpoint 4 : Delete_book_by_id
// TODO:
func Delete_book_by_id(c echo.Context) error {
	bookId := c.Param("id")
	var theBook model.Book

	// check is id exists
	if err_first := config.DB.Where("id=?", bookId).First(&theBook).Error; err_first != nil {
		log.Print(color.RedString(err_first.Error()))
		return c.JSON((http.StatusBadRequest), map[string]interface{}{
			"status":  400,
			"message": "bad request",
		})
	}

	// delete
	if err_delete := config.DB.Where("id=?", bookId).Delete(&theBook).Error; err_delete != nil {
		log.Print(color.RedString(err_delete.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  500,
			"message": "internal server error",
		})
	}

	responseSuccess := model.ResponseSuccessBookOne{
		StatusCode:   202,
		ErrorMessage: "success to delete book by id",
		Data:         theBook,
	}

	return c.JSON(http.StatusCreated, responseSuccess)
}
