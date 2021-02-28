package app

import "gopkg.in/ini.v1"

const (
	commonSectionName = "common"
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

// GetAllFrpServiceConf return all service config
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

// GetFrpServiceConfByName return specified config
func GetFrpServiceConfByName(secName string) (*ServiceConf, error) {
	sec, err := cfg.GetSection(secName)
	if err != nil {
		return nil, err
	}
	para := make(map[string]string)
	for _, key := range sec.Keys() {
		para[key.Name()] = key.String()
	}w
	return &ServiceConf{
		Name: sec.Name(),
		Para: para,
	}, nil
}

// AddFrpServiceConf add service config
func AddFrpServiceConf(secName string, kvMap map[string]string) error {
	if err := cfg.NewSections(secName); err != nil {
		return err
	}
	sec, err := cfg.GetSection(secName)
	if err != nil {
		return err
	}
	for key, val := range kvMap {
		sec.NewKey(key, val)
	}
	return nil
}

// DeleteFrpServiceConf delete service config by name
func DeleteFrpServiceConf(secName string) error {
	cfg.DeleteSection(secName)
	return nil
}

// UpdateFrpServiceConf delete service config by name
func UpdateFrpServiceConf(secName string, kvMap map[string]string) error {
	sec, err := cfg.GetSection(secName)
	if err != nil {
		return err
	}
	for key, val := range kvMap {
		if sec.HasKey(key) {
			sec.Key(key).SetValue(val)
		}
	}
	return nil
}
