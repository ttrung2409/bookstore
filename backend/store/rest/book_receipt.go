package rest

import (
	"net/http"
	module "store"
	op "store/app/operation"
	"store/utils"

	"github.com/gin-gonic/gin"
)

func BookReceiptRoutes(r *gin.Engine) {
	controller := bookReceiptController{}
	r.POST("/", controller.create())
}

var bookReceiving = module.Container().Get(utils.Nameof((*op.BookReceiving)(nil))).(op.BookReceiving)

type bookReceiptController struct{}

func (c *bookReceiptController) create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request op.ReceiveBooksRequest
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
