package rest

import (
	"errors"
	"net/http"
	"store/app/order/command"
	"store/app/order/query"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func addOrderRoutes(r *gin.RouterGroup) {
	controller := &orderController{}
	r.GET("/", controller.find())
	r.GET("/:id", controller.get())
	r.PUT("/:id/deliver", controller.deliver())
}

type orderController struct{}

func (c *orderController) find() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := query.New()
		orders, err := query.FindDeliverableOrders()

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, err)
				return
			}

			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, orders)
	}
}

func (c *orderController) get() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := query.New()
		order, err := query.GetOrderDetails(c.Query("id"))

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, err)
				return
			}

			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, order)
	}
}

func (c *orderController) deliver() gin.HandlerFunc {
	return func(c *gin.Context) {
		command := command.New()
		err := command.DeliverOrder(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
