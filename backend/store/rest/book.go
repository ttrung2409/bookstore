package rest

import (
	"net/http"
	ReceivingQuery "store/app/receiving/query"
	"store/container"
	"store/utils"

	"github.com/gin-gonic/gin"
)

func addBookRoutes(r *gin.RouterGroup) {
	controller := &bookController{}
	r.GET("/", controller.find())
}

type bookController struct{}

func (c *bookController) find() gin.HandlerFunc {
	return func(c *gin.Context) {
		receivingQuery := container.Instance().Get(utils.Nameof((*ReceivingQuery.Query)(nil))).(ReceivingQuery.Query)

		books, err := receivingQuery.FindGoogleBooks(c.Query("term"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, books)
	}
}
