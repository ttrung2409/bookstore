package rest

import (
	"net/http"
	module "store"
	command "store/app/operation/command"
	"store/utils"

	"github.com/gin-gonic/gin"
)

func BookReceiptRoutes(r *gin.Engine) {
	controller := bookReceiptController{}
	r.POST("/", controller.create())
}

var bookReceiving = module.Container().Get(utils.Nameof((*command.BookReceiving)(nil))).(command.BookReceiving)

type bookReceiptController struct{}

func (c *bookReceiptController) create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request command.ReceiveBooksRequest
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		err = bookReceiving.Receive(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
