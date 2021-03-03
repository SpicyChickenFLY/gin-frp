package frp

import (
	"net/http"

	"github.com/SpicyChickenFLY/gin-frp/model"
	"github.com/SpicyChickenFLY/gin-frp/service"
	"github.com/gin-gonic/gin"
)

// CreateFrpTransportService create new TCP service
func CreateFrpTransportService(c *gin.Context) {
	frpTransportServiceData := &model.FrpTransportService{}
	if err := c.BindJSON(frpTransportServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	if err := service.AddFrpServiceConf(frpTransportServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetFrpTransportService get TCP service by name
func GetFrpTransportService(c *gin.Context) {
	serviceName := c.PostForm("ServiceName")
	result := &model.FrpTransportService{}
	if err := service.GetFrpServiceConfByName(serviceName, result); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateFrpTransportService update TCP service
func UpdateFrpTransportService(c *gin.Context) {
	frpTransportServiceData := &model.FrpTransportService{}
	if err := c.BindJSON(frpTransportServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	if err := service.DeleteFrpServiceConf(frpTransportServiceData.ServiceName); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteFrpTransportService delete TCP service
func DeleteFrpTransportService(c *gin.Context) {
	serviceName := c.PostForm("ServiceName")
	if err := service.DeleteFrpServiceConf(serviceName); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// CreateFrpSecureTransportService create new secure Transport service
func CreateFrpSecureTransportService(c *gin.Context) {
	frpSecureTransportServiceData := &model.FrpSecureTransportService{}
	if err := c.BindJSON(frpSecureTransportServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	if err := service.AddFrpServiceConf(frpSecureTransportServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetFrpSecureTransportService get secure Transport service by name
func GetFrpSecureTransportService(c *gin.Context) {
	serviceName := c.PostForm("ServiceName")
	result := &model.FrpSecureTransportService{}
	if err := service.GetFrpServiceConfByName(serviceName, result); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateFrpSecureTransportService update secure Transport service
func UpdateFrpSecureTransportService(c *gin.Context) {
	frpSecureTransportServiceData := &model.FrpSecureTransportService{}
	if err := c.BindJSON(frpSecureTransportServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	if err := service.DeleteFrpServiceConf(frpSecureTransportServiceData.ServiceName); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteFrpSecureTransportService delete secure Transport service
func DeleteFrpSecureTransportService(c *gin.Context) {
	serviceName := c.PostForm("ServiceName")
	if err := service.DeleteFrpServiceConf(serviceName); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}
