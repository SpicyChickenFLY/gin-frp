package app

import (
	"reflect"

	"gopkg.in/ini.v1"
)

const (
	commonSectionName = "common"
	cfgTag            = "ini"
)

var cfg *ini.File

// ServiceConf is a service configure struct
type ServiceConf struct {
	Name string
	Para map[string]string
}

// LoadConfFromFile load config from conf file
func LoadConfFromFile(filePath string) (err error) {
	cfg, err = ini.Load(filePath)
	if err != nil {
		return err
	}
	return nil
}

// AddFrpServiceConf add service config
func AddFrpServiceConf(service interface{}) error {
	reflectType := reflect.TypeOf(service)
	serviceName := reflectType.Field(0).Name
	sec, err := cfg.NewSection(serviceName)
	if err != nil {
		return err
	}
	for i := 1; i < reflectType.NumField(); i++ {
		key := reflect.TypeOf(service).Field(i).Tag.Get(cfgTag)
		val := reflect.ValueOf(service).Field(i).String()
		sec.NewKey(key, val)
	}
	return nil
}

// GetAllFrpServiceConf return all service config
// FIXME
func GetAllFrpServiceConf() (serviceConfs []*ServiceConf) {
	for _, sec := range cfg.Sections() {
		if sec.Name() == commonSectionName {
			continue
		}
		para := make(map[string]string)
		for _, key := range sec.Keys() {
			para[key.Name()] = key.String()
		}
		serviceConfs = append(serviceConfs, &ServiceConf{
			Name: sec.Name(),
			Para: para,
		})
	}
	return serviceConfs
}

// GetFrpServiceConfByType return specified config
func GetFrpServiceConfByType(serviceType string) *ServiceConf {

}

// GetFrpServiceConfByName return specified config
func GetFrpServiceConfByName(
	ServiceName string, serviceStruct interface{}) error {
	return cfg.Section(ServiceName).MapTo(serviceStruct)
}

// DeleteFrpServiceConf delete service config by name
func DeleteFrpServiceConf(secName string) error {
	cfg.DeleteSection(secName)
	return nil
}

// UpdateFrpServiceConf delete service config by name
func UpdateFrpServiceConf(service interface{}) error {
	reflectType := reflect.TypeOf(service)
	serviceName := reflectType.Field(0).Name
	sec, err := cfg.GetSection(serviceName)
	if err != nil {
		return err
	}
	for i := 1; i < reflectType.NumField(); i++ {
		key := reflect.TypeOf(service).Field(i).Tag.Get(cfgTag)
		val := reflect.ValueOf(service).Field(i).String()
		sec.NewKey(key, val)
	}
	return nil
}
