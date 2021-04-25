package rest

import (
	"net/http"
	module "store"
	op "store/app/operation"
	"store/utils"

	"github.com/gin-gonic/gin"
)

func BookRoutes(r *gin.Engine) {
	r.GET("/", find())
}

var bookFinder = module.Container.Get(utils.Nameof((*op.FindBooks)(nil))).(op.FindBooks)

func find() gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := bookFinder.Find(c.Query("term"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, books)
	}
}
