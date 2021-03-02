package frp

import (
	"net/http"

	"github.com/SpicyChickenFLY/gin-frp/app"
	"github.com/SpicyChickenFLY/gin-frp/model"
	"github.com/gin-gonic/gin"
)

// CreateFrpTCPService create new TCP service
func CreateFrpTCPService(c *gin.Context) {
	frpTCPServiceData := &model.FrpTCPService{}
	if err := c.BindJSON(frpTCPServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	if err := app.AddFrpServiceConf(frpTCPServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetFrpTCPService get TCP service by name
func GetFrpTCPService(c *gin.Context) {
	serviceName := c.PostForm("ServiceName")
	result := &model.FrpTCPService{}
	if err := app.GetFrpServiceConfByName(serviceName, result); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateFrpTCPService update TCP service
func UpdateFrpTCPService(c *gin.Context) {
	frpTCPServiceData := &model.FrpTCPService{}
	if err := c.BindJSON(frpTCPServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	if err := app.DeleteFrpServiceConf(frpTCPServiceData.ServiceName); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteFrpTCPService delete TCP service
func DeleteFrpTCPService(c *gin.Context) {
	serviceName := c.PostForm("ServiceName")
	if err := app.DeleteFrpServiceConf(serviceName); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}
