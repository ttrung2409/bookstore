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

var googleBookQuery = module.Container().Get(utils.Nameof((*query.GoogleBookQuery)(nil))).(query.GoogleBookQuery)

type bookController struct{}

func (c *bookController) find() gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := googleBookQuery.Find(c.Query("term"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, books)
	}
}
