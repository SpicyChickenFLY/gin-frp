package api

import (
	"github.com/SpicyChickenFLY/gin-frp/model"
	"github.com/SpicyChickenFLY/gin-frp/service"
	"github.com/gin-gonic/gin"
)

func GetAllFrpServices(c *gin.Context) {

}

func CreateFrpTCPService(c *gin.Context) {
	newFrpTCPServiceData := &model.FrpTCPService{}
	if err := c.BindJSON(newFrpTCPServiceData); err != nil {

	}
	service.CreateFrpTCPService(newFrpTCPServiceData)
}

func GetFrpTCPService(c *gin.Context) {

}

func UpdateFrpTCPService(c *gin.Context) {

}

func DeleteFrpTCPService(c *gin.Context) {

}
