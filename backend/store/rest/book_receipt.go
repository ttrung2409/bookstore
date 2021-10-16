package rest

import (
	"net/http"
	module "store"
	command "store/app/receiving/command"
	"store/utils"

	"github.com/gin-gonic/gin"
)

func BookReceiptRoutes(r *gin.Engine) {
	controller := bookReceiptController{}
	r.POST("/", controller.create())
}

var receiveBookCommand = module.Container().Get(utils.Nameof((*command.ReceiveBookCommand)(nil))).(command.ReceiveBookCommand)

type bookReceiptController struct{}

func (c *bookReceiptController) create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request command.ReceiveBooksRequest
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		err = receiveBookCommand.Execute(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
