package rest

import "github.com/gin-gonic/gin"

func Start() {
	router := gin.Default()

	addBookRoutes(router.Group("/book"))
	addBookReceiptRoutes(router.Group("/book-receipt"))
	addOrderRoutes(router.Group("/order"))

	router.Run(":8080")
}
