package frp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexPage return index.html page
func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
