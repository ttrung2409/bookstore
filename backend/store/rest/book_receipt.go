package rest

import (
	"net/http"

	ReceivingCommand "store/app/receiving/command"

	"github.com/gin-gonic/gin"
)

func BookReceiptRoutes(r *gin.Engine) {
	controller := bookReceiptController{}
	r.POST("/", controller.create())
}

type bookReceiptController struct{}

func (c *bookReceiptController) create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request ReceivingCommand.ReceiveBooksRequest
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		err = receivingCommand.Receive(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
