package main

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 靜態資源
	r.Use(static.Serve("/", static.LocalFile("./resource", false)))

	// Route
	r.GET("/api", handler) //

	r.Run(":8080")
}

func handler(c *gin.Context) {

	c.JSON(http.StatusOK, "gaetest")
}
