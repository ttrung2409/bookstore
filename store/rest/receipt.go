package rest

import (
	"net/http"

	"store/app/receiving/command"

	"github.com/gin-gonic/gin"
)

func addReceiptRoutes(r *gin.RouterGroup) {
	controller := receiptController{}
	r.POST("/", controller.create())
}

type receiptController struct{}

func (c *receiptController) create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request command.ReceiveBooksRequest
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		command := command.New()
		err = command.Receive(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
