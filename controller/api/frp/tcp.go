package frp

import (
	"net/http"

	"github.com/SpicyChickenFLY/gin-frp/model"
	"github.com/SpicyChickenFLY/gin-frp/service"
	"github.com/gin-gonic/gin"
)

// CreateFrpTCPService create new TCP service
func CreateFrpTCPService(c *gin.Context) {
	frpTCPServiceData := &model.FrpTCPUDPService{}
	if err := c.BindJSON(frpTCPServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	if err := service.AddFrpServiceConf(frpTCPServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetFrpTCPService get TCP service by name
func GetFrpTCPService(c *gin.Context) {
	serviceName := c.PostForm("ServiceName")
	result := &model.FrpTCPUDPService{}
	if err := service.GetFrpServiceConfByName(serviceName, result); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateFrpTCPService update TCP service
func UpdateFrpTCPService(c *gin.Context) {
	frpTCPServiceData := &model.FrpTCPUDPService{}
	if err := c.BindJSON(frpTCPServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	if err := service.DeleteFrpServiceConf(frpTCPServiceData.ServiceName); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteFrpTCPService delete TCP service
func DeleteFrpTCPService(c *gin.Context) {
	serviceName := c.PostForm("ServiceName")
	if err := service.DeleteFrpServiceConf(serviceName); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}
