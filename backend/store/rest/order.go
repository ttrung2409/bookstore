package rest

import (
	"net/http"
	module "store"
	"store/app/data"
	op "store/app/operation"
	"store/utils"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	controller := &orderController{}
	r.GET("/", controller.find())
	r.GET("/:id", controller.get())
	r.PUT("/:id/accept", controller.accept())
}

var orderQuery = module.Container.Get(utils.Nameof((*op.OrderQuery)(nil))).(op.OrderQuery)
var acceptOrder = module.Container.Get(utils.Nameof((*op.AcceptOrder)(nil))).(op.AcceptOrder)

type orderController struct{}

func (c *orderController) find() gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := orderQuery.FindByStatus(
			[]string{string(data.OrderStatusQueued), string(data.OrderStatusStockFilled)},
		)

		if err != nil {
			c.JSON(getHttpStatusByError(err), err)
			return
		}

		c.JSON(http.StatusOK, books)
	}
}

func (c *orderController) get() gin.HandlerFunc {
	return func(c *gin.Context) {
		order, err := orderQuery.Get(c.Query("id"))
		if err != nil {
			c.JSON(getHttpStatusByError(err), err)
			return
		}

		c.JSON(http.StatusOK, order)
	}
}

func (c *orderController) accept() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := acceptOrder.Accept(c.Query("id"))
		if err != nil {
			c.JSON(getHttpStatusByError(err), err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
