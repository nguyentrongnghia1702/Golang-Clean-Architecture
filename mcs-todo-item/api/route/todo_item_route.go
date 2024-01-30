package route

import (
	"mcs-nghiadeptrai/mcs-todo-item/api/controller"
	"mcs-nghiadeptrai/mcs-todo-item/repository"
	"mcs-nghiadeptrai/mcs-todo-item/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTodoItemRouter(timeout time.Duration, db *gorm.DB, gin *gin.RouterGroup) {
	tr := repository.NewTodoRepository(db)
	tc := &controller.TodoItemController{
		TodoUsecase: usecase.NewTodoUsecase(tr, timeout),
	}

	gin.POST("/items", tc.CreateItem)       // create item
	gin.GET("/items", tc.GetAllItem)        // list items
	gin.PUT("/items/:id", tc.UpdateItem)    // edit an item by ID
	gin.DELETE("/items/:id", tc.DeleteItem) // delete an item by ID

}
