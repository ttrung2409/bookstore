package rest

import (
	"net/http"
	module "store"
	op "store/app/operation"
	"store/utils"

	"github.com/gin-gonic/gin"
)

func BookReceiptRoutes(r *gin.Engine) {
	r.POST("/", create())
}

var bookReceiver = module.Container.Get(utils.Nameof((*op.ReceiveBooks)(nil))).(op.ReceiveBooks)

func create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request op.ReceiveBooksRequest
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		err = bookReceiver.Receive(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
