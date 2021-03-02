package app

import (
	"reflect"

	"github.com/SpicyChickenFLY/auto-mycnf/model"
	"gopkg.in/ini.v1"
)

const (
	commonSectionName = "common"
	serviceTypeTag    = "type"
	cfgStructTag      = "ini"
)

var cfg *ini.File

// LoadConfFromFile load config from conf file
func LoadConfFromFile(filePath string) (err error) {
	cfg, err = ini.Load(filePath)
	if err != nil {
		return err
	}
	return nil
}

// AddFrpServiceConf add service config
func AddFrpServiceConf(serviceStruct interface{}) error {
	reflectType := reflect.TypeOf(serviceStruct)
	serviceName := reflectType.Field(0).Name
	sec, err := cfg.NewSection(serviceName)
	if err != nil {
		return err
	}
	for i := 1; i < reflectType.NumField(); i++ {
		key := reflect.TypeOf(serviceStruct).Field(i).Tag.Get(cfgStructTag)
		val := reflect.ValueOf(serviceStruct).Field(i).String()
		sec.NewKey(key, val)
	}
	return nil
}

// GetAllFrpServiceConfsTypeAndName return all service config
func GetAllFrpServiceConfsTypeAndName() (serviceBasicConfs []*model.FrpBasicService) {
	for _, sec := range cfg.Sections() {
		if sec.Name() == commonSectionName {
			continue
		}
		serviceBasicConfs = append(serviceBasicConfs,
			&model.FrpBasicService{
				Name: sec.Name(),
				Type: sec.Key(serviceTypeTag).String(),
			})
	}
	return serviceBasicConfs
}

// GetFrpServiceConfsByTag return specified config
func GetFrpServiceConfsByTag(
	ServiceName string, serviceStruct interface{}) error {
	return cfg.Section(ServiceName).MapTo(serviceStruct)
}

// GetFrpServiceConfByName return specified config
func GetFrpServiceConfByName(
	ServiceName string, serviceStruct interface{}) error {
	return cfg.Section(ServiceName).MapTo(serviceStruct)
}

// DeleteFrpServiceConf delete service config by name
func DeleteFrpServiceConf(serviceName string) error {
	cfg.DeleteSection(serviceName)
	return nil
}

// UpdateFrpServiceConf delete service config by name
func UpdateFrpServiceConf(serviceName string, serviceStruct interface{}) error {
	reflectType := reflect.TypeOf(serviceStruct)
	sec, err := cfg.GetSection(serviceName)
	if err != nil {
		return err
	}
	for i := 1; i < reflectType.NumField(); i++ {
		key := reflect.TypeOf(serviceStruct).Field(i).Tag.Get(cfgStructTag)
		val := reflect.ValueOf(serviceStruct).Field(i).String()
		sec.NewKey(key, val)
	}
	return nil
}
