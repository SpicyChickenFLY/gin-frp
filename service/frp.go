package service

import (
	"github.com/SpicyChickenFLY/gin-frp/app"
	"github.com/SpicyChickenFLY/gin-frp/model"
)

// GetAllFrpService get all  service
func GetAllFrpService() []*model.FrpBasicService {
	return app.GetAllFrpServiceConfsTypeAndName()
}

// ===================== TCP =====================

// CreateFrpTCPService create new TCP service
func CreateFrpTCPService(frpTCPServiceData *model.FrpTCPService) error {
	return app.AddFrpServiceConf(frpTCPServiceData)
}

// GetFrpTCPServiceByName get TCP service by name
func GetFrpTCPServiceByName(
	serviceName string, result *model.FrpTCPService) error {
	return app.GetFrpServiceConfByName(serviceName, result)
}

//UpdateFrpTCPService update TCP service
func UpdateFrpTCPService(
	serviceName string, frpTCPServiceData model.FrpTCPService) error {
	return app.UpdateFrpServiceConf(serviceName, frpTCPServiceData)
}

// DeleteFrpTCPService delete TCP service
func DeleteFrpTCPService(serviceName string) error {
	return app.DeleteFrpServiceConf(serviceName)
}

// ===================== File =====================

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
