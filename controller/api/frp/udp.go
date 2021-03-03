package frp

import (
	"net/http"

	"github.com/SpicyChickenFLY/gin-frp/model"
	"github.com/SpicyChickenFLY/gin-frp/service"
	"github.com/gin-gonic/gin"
)

// CreateFrpUDPService create new UDP service
func CreateFrpUDPService(c *gin.Context) {
	frpUDPServiceData := &model.FrpTCPUDPService{}
	if err := c.BindJSON(frpUDPServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	if err := service.AddFrpServiceConf(frpUDPServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetFrpUDPService get UDP service by name
func GetFrpUDPService(c *gin.Context) {
	serviceName := c.PostForm("ServiceName")
	result := &model.FrpTCPUDPService{}
	if err := service.GetFrpServiceConfByName(serviceName, result); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateFrpUDPService update UDP service
func UpdateFrpUDPService(c *gin.Context) {
	frpUDPServiceData := &model.FrpTCPUDPService{}
	if err := c.BindJSON(frpUDPServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	if err := service.DeleteFrpServiceConf(frpUDPServiceData.ServiceName); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteFrpUDPService delete UDP service
func DeleteFrpUDPService(c *gin.Context) {
	serviceName := c.PostForm("ServiceName")
	if err := service.DeleteFrpServiceConf(serviceName); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}
