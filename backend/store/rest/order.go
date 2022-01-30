package rest

import (
	"net/http"
	OrderCommand "store/app/order/command"
	OrderQuery "store/app/order/query"
	"store/container"
	"store/utils"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	controller := &orderController{}
	r.GET("/", controller.find())
	r.GET("/:id", controller.get())
	r.PUT("/:id/accept", controller.accept())
	r.PUT("/:id/place-as-back-order", controller.placeAsBackOrder())
	r.PUT("/:id/reject", controller.reject())
}

type orderController struct{}

func (c *orderController) find() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderQuery := container.Instance().Get(utils.Nameof((*OrderQuery.Query)(nil))).(OrderQuery.Query)

		orders, err := orderQuery.FindDeliverableOrders()
		if err != nil {
			c.JSON(getHttpStatusByError(err), err)
			return
		}

		c.JSON(http.StatusOK, orders)
	}
}

func (c *orderController) get() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderQuery := container.Instance().Get(utils.Nameof((*OrderQuery.Query)(nil))).(OrderQuery.Query)

		order, err := orderQuery.GetOrderDetails(c.Query("id"))
		if err != nil {
			c.JSON(getHttpStatusByError(err), err)
			return
		}

		c.JSON(http.StatusOK, order)
	}
}

func (c *orderController) accept() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderCommand := container.Instance().Get(utils.Nameof((*OrderCommand.Command)(nil))).(OrderCommand.Command)

		err := orderCommand.AcceptOrder(c.Query("id"))
		if err != nil {
			c.JSON(getHttpStatusByError(err), err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}

func (c *orderController) placeAsBackOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderCommand := container.Instance().Get(utils.Nameof((*OrderCommand.Command)(nil))).(OrderCommand.Command)

		err := orderCommand.PlaceAsBackOrder(c.Query("id"))
		if err != nil {
			c.JSON(getHttpStatusByError(err), err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}

func (c *orderController) reject() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderCommand := container.Instance().Get(utils.Nameof((*OrderCommand.Command)(nil))).(OrderCommand.Command)

		err := orderCommand.RejectOrder(c.Query("id"))
		if err != nil {
			c.JSON(getHttpStatusByError(err), err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
