package api

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func CreateFrpService(c *gin.Context) {
	cfg, err := ini.Load("my.ini")
	if err != nil {
	}
}

func GetAllFrpServices(c *gin.Context) {
	cfg, err := ini.Load("my.ini")
	if err != nil {
	}
}

func GetFrpServiceByName(c *gin.Context) {

}

func UpdateFrpService(c *gin.Context) {

}

func DeleteFrpService(c *gin.Context) {

}
