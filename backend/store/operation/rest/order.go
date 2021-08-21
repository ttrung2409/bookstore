package rest

import (
	"net/http"
	module "store"
	command "store/app/operation/command"
	query "store/app/operation/query"
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

var orderQuery = module.Container().Get(utils.Nameof((*query.OrderQuery)(nil))).(query.OrderQuery)

var acceptOrderCommand = module.Container().Get(utils.Nameof((*command.AcceptOrderCommand)(nil))).(command.AcceptOrderCommand)

var placeAsBackOrderCommand = module.Container().Get(utils.Nameof((*command.PlaceAsBackOrderCommand)(nil))).(command.PlaceAsBackOrderCommand)

var rejectOrderCommand = module.Container().Get(utils.Nameof((*command.RejectOrderCommand)(nil))).(command.RejectOrderCommand)

type orderController struct{}

func (c *orderController) find() gin.HandlerFunc {
	return func(c *gin.Context) {
		orders, err := orderQuery.FindOrdersToDeliver()
		if err != nil {
			c.JSON(getHttpStatusByError(err), err)
			return
		}

		c.JSON(http.StatusOK, orders)
	}
}

func (c *orderController) get() gin.HandlerFunc {
	return func(c *gin.Context) {
		order, err := orderQuery.GetOrderToView(c.Query("id"))
		if err != nil {
			c.JSON(getHttpStatusByError(err), err)
			return
		}

		c.JSON(http.StatusOK, order)
	}
}

func (c *orderController) accept() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := acceptOrderCommand.Execute(c.Query("id"))
		if err != nil {
			c.JSON(getHttpStatusByError(err), err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}

func (c *orderController) placeAsBackOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := placeAsBackOrderCommand.Execute(c.Query("id"))
		if err != nil {
			c.JSON(getHttpStatusByError(err), err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}

func (c *orderController) reject() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := rejectOrderCommand.Execute(c.Query("id"))
		if err != nil {
			c.JSON(getHttpStatusByError(err), err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
