package service

import (
	"github.com/SpicyChickenFLY/auto-mycnf/app"
	"github.com/SpicyChickenFLY/gin-frp/model"
)

// ===================== TCP =====================

// CreateFrpTCPService create new TCP service
func CreateFrpTCPService(frpTCPServ model.FrpTCPService) error {
	return app.AddFrpServiceConf(frpTCPServ)
}

// GetAllFrpTCPService get all TCP service
func GetAllFrpTCPService() []*model.FrpTCPService {

}

// GetFrpTCPServiceByName get TCP service by name
func GetFrpTCPServiceByName(serviceName string) {

}

//UpdateFrpTCPService update TCP service
func UpdateFrpTCPService(frpTCPServ model.FrpTCPService) {

}

// DeleteFrpTCPService delete TCP service
func DeleteFrpTCPService(serviceName string) {

}

// ===================== File =====================

// GetAllFrpFileService get all File service
func GetAllFrpFileService() []*model.FrpFileService {

}

// GetFrpFileServiceByName get File service by name
func GetFrpFileServiceByName(serviceName string) {

}

//UpdateFrpFileService update File service
func UpdateFrpFileService(frpFileServ model.FrpFileService) {

}

// DeleteFrpFileService delete File service
func DeleteFrpFileService(serviceName string) {

}

// ===================== HTTPS =====================

// GetAllFrpHTTPSService get all HTTPS service
func GetAllFrpHTTPSService() []*model.FrpHTTPSService {

}

// GetFrpHTTPSServiceByName get HTTPS service by name
func GetFrpHTTPSServiceByName(serviceName string) {

}

//UpdateFrpHTTPSService update HTTPS service
func UpdateFrpHTTPSService(frpHTTPSServ model.FrpHTTPSService) {

}

// DeleteFrpHTTPSService delete HTTPS service
func DeleteFrpHTTPSService(serviceName string) {

}
