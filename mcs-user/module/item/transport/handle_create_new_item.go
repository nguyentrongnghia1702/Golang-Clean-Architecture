package todotrpt

import (
	"net/http"
	"strings"

	todobiz "mcs-nghiadeptrai/mcs-user/module/item/business"
	todomodel "mcs-nghiadeptrai/mcs-user/module/item/model"
	todostorage "mcs-nghiadeptrai/mcs-user/module/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HanleCreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem todomodel.ToDoItem

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// preprocess title - trim all spaces
		dataItem.Title = strings.TrimSpace(dataItem.Title)

		// setup dependencies
		storage := todostorage.NewMySQLStorage(db)

		biz := todobiz.NewCreateToDoItemBiz(storage)

		if err := biz.CreateNewItem(c.Request.Context(), &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dataItem.Id})
	}
}
