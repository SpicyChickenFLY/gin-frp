package frp

import (
	"net/http"

	"github.com/SpicyChickenFLY/gin-frp/model"
	"github.com/SpicyChickenFLY/gin-frp/service"
	"github.com/gin-gonic/gin"
)

// CreateFrpFileService create new File service
func CreateFrpFileService(c *gin.Context) {
	frpFileServiceData := &model.FrpFileService{}
	if err := c.BindJSON(frpFileServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	if err := service.AddFrpServiceConf(frpFileServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// GetFrpFileService get File service by name
func GetFrpFileService(c *gin.Context) {
	serviceName := c.PostForm("ServiceName")
	result := &model.FrpFileService{}
	if err := service.GetFrpServiceConfByName(serviceName, result); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateFrpFileService update File service
func UpdateFrpFileService(c *gin.Context) {
	frpFileServiceData := &model.FrpFileService{}
	if err := c.BindJSON(frpFileServiceData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	if err := service.DeleteFrpServiceConf(frpFileServiceData.ServiceName); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteFrpFileService delete File service
func DeleteFrpFileService(c *gin.Context) {
	serviceName := c.PostForm("ServiceName")
	if err := service.DeleteFrpServiceConf(serviceName); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}
