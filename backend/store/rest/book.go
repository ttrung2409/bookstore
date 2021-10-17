package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BookRoutes(r *gin.Engine) {
	controller := &bookController{}
	r.GET("/", controller.find())
}

type bookController struct{}

func (c *bookController) find() gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := receivingQuery.SearchGoogleBooks(c.Query("term"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, books)
	}
}
