package frp

import (
	"net/http"

	"github.com/SpicyChickenFLY/gin-frp/service"
	"github.com/gin-gonic/gin"
)

// GetAllFrpServices return name and type of  all service
func GetAllFrpServices(c *gin.Context) {
	result := service.GetAllFrpServiceConfsTypeAndName()
	c.JSON(http.StatusOK, gin.H{"result": result})
}
