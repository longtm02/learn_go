package controller

import (
	"GolangTodoApp/src/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string      `json:"message"`
	Params  interface{} `json:"params"`
}

func AddTodo(context *gin.Context) {
	body := store.Todo{}

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, Error{Message: "Incorrect format", Params: body})
		return
	}

	store.Id++
	body.Id = store.Id
	store.TodoManagement = append(store.TodoManagement, body)
	context.JSON(http.StatusOK, store.TodoManagement)
}

func ListTodo(context *gin.Context) {
	context.JSON(http.StatusOK, store.TodoManagement)
}

func UpdateTodo(context *gin.Context) {
	body := store.Todo{}
	query := context.Param("id")

	id, errorQuery := strconv.Atoi(query)

	if errorQuery != nil {
		context.JSON(http.StatusBadRequest, Error{Message: "Incorrect url", Params: query})
		return
	}

	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, Error{Message: "Incorrect format", Params: body})
		return
	}

	for index, value := range store.TodoManagement {
		if id == value.Id {
			store.TodoManagement[index] = store.Todo{
				Id:     id,
				Title:  body.Title,
				Status: body.Status,
			}
			context.JSON(http.StatusOK, store.TodoManagement)
			return
		}
	}

	context.JSON(http.StatusBadRequest, Error{Message: "not found item", Params: id})
}

func GetTodo(context *gin.Context) {
	query := context.Param("id")

	id, errorQuery := strconv.Atoi(query)

	if errorQuery != nil {
		context.JSON(http.StatusBadRequest, Error{Message: "Incorrect url", Params: query})
		return
	}

	for _, value := range store.TodoManagement {
		if id == value.Id {
			context.JSON(http.StatusOK, value)
			return
		}
	}

	context.JSON(http.StatusBadRequest, Error{Message: "not found item", Params: id})
}

func DeleteTodo(context *gin.Context) {

	query := context.Param("id")

	id, errorQuery := strconv.Atoi(query)

	if errorQuery != nil {
		context.JSON(http.StatusBadRequest, Error{Message: "Incorrect url", Params: query})
		return
	}

	for index, value := range store.TodoManagement {
		if id == value.Id {
			store.TodoManagement = removeSlice(store.TodoManagement, index)
			context.JSON(http.StatusOK, store.TodoManagement)
			return
		}
	}

	context.JSON(http.StatusBadRequest, Error{Message: "not found item", Params: id})

}

func removeSlice(slice []store.Todo, index int) []store.Todo {
	return append(slice[:index], slice[index+1:]...)
}
