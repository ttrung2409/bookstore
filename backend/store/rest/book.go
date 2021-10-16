package rest

import (
	"net/http"
	module "store"
	query "store/app/receiving/query"
	"store/utils"

	"github.com/gin-gonic/gin"
)

func BookRoutes(r *gin.Engine) {
	controller := &bookController{}
	r.GET("/", controller.find())
}

var bookQuery = module.Container().Get(utils.Nameof((*query.BookQuery)(nil))).(query.BookQuery)

type bookController struct{}

func (c *bookController) find() gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := bookQuery.SearchGoogleBooks(c.Query("term"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, books)
	}
}
