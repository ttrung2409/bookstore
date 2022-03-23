package rest

import (
	"net/http"

	ReceivingCommand "store/app/receiving/command"
	"store/container"
	"store/utils"

	"github.com/gin-gonic/gin"
)

func addBookReceiptRoutes(r *gin.RouterGroup) {
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

		receivingCommand := container.Instance().Get(utils.Nameof((*ReceivingCommand.Command)(nil))).(ReceivingCommand.Command)

		err = receivingCommand.Receive(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
