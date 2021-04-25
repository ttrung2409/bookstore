package rest

import "github.com/gin-gonic/gin"

func Start() {
	r := gin.Default()
	r.Group("/book")
	{
		BookRoutes(r)
	}

	r.Group("/book-receipt")
	{
		BookReceiptRoutes(r)
	}

	r.Run(":8080")
}
