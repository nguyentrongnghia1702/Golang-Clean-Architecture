package todotrpt

import (
	"net/http"

	todobiz "mcs-nghiadeptrai/mcs-user/module/item/business"
	todomodel "mcs-nghiadeptrai/mcs-user/module/item/model"
	todostorage "mcs-nghiadeptrai/mcs-user/module/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleListItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging todomodel.DataPaging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		paging.Process()

		storage := todostorage.NewMySQLStorage(db)

		biz := todobiz.NewListToDoItemBiz(storage)

		result, err := biz.ListItems(c.Request.Context(), nil, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})
	}
}
