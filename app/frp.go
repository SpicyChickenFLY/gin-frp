package app

import (
	"reflect"

	"gopkg.in/ini.v1"
)

const (
	commonSectionName = "common"
	cfgStructTag      = "ini"
)

const (
	serviceTypeTag      = "type"
	udpServiceTypeTag   = "tcp"
	tcpServiceTypeTag   = "tcp"
	stcpServiceTypeTag  = "stcp"
	xtcpServiceTypeTag  = "xtcp"
	httpServiceTypeTag  = "http"
	httpsServiceTypeTag = "https"
)

var cfg *ini.File

// ServiceBasicConf is a service configure struct
type ServiceBasicConf struct {
	Name string
	Type string
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

// GetAllFrpServiceConfTypeAndName return all service config
// FIXME:
func GetAllFrpServiceConfTypeAndName() (serviceBasicConfs []*ServiceBasicConf) {
	for _, sec := range cfg.Sections() {
		if sec.Name() == commonSectionName {
			continue
		}
		serviceBasicConfs = append(serviceBasicConfs,
			&ServiceBasicConf{
				Name: sec.Name(),
				Type: sec.Key(serviceTypeTag).String(),
			})
	}
	return serviceBasicConfs
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
		key := reflect.TypeOf(service).Field(i).Tag.Get(cfgStructTag)
		val := reflect.ValueOf(service).Field(i).String()
		sec.NewKey(key, val)
	}
	return nil
}
