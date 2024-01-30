package controller

import (
	//import domain
	"mcs-nghiadeptrai/mcs-todo-item/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoItemController struct {
	TodoUsecase domain.TodoItemUseCase
}

func (controller *TodoItemController) CreateItem(c *gin.Context) {
	var todoItem domain.ToDoItem
	err := c.ShouldBindJSON(&todoItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = controller.TodoUsecase.CreateItem(c, &todoItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Create Successfully",
	})
}

func (controller *TodoItemController) DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = controller.TodoUsecase.DeleteItem(c, map[string]interface{}{"id": id})
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Delete Successfully",
	})
}

func (controller *TodoItemController) UpdateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	var todoItem domain.ToDoItem
	err = c.ShouldBindJSON(&todoItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	oldData, err := controller.TodoUsecase.GetItemById(c, map[string]interface{}{"id": id})
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = controller.TodoUsecase.UpdateItem(c, map[string]interface{}{"id": id}, &todoItem, oldData.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Update Successfully",
	})

}

func (controller *TodoItemController) GetAllItem(c *gin.Context) {
	var paging domain.DataPaging

	if err := c.ShouldBind(&paging); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paging.Process()

	result, err := controller.TodoUsecase.GetAllItem(c.Request.Context(), &paging)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})
}
