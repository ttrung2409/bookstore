package rest

import (
	"net/http"
	"store/app/book/query"

	"github.com/gin-gonic/gin"
)

func addBookRoutes(r *gin.RouterGroup) {
	controller := &bookController{}
	r.GET("/", controller.find())
}

type bookController struct{}

func (c *bookController) find() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := query.New()
		books, err := query.FindGoogleBooks(c.Query("term"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, books)
	}
}
