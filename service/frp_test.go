package service

import (
	"reflect"
	"testing"

	"github.com/SpicyChickenFLY/gin-frp/model"
)

func TestLoadConfFromFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadConfFromFile(tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("LoadConfFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddFrpServiceConf(t *testing.T) {
	type args struct {
		serviceStruct interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddFrpServiceConf(tt.args.serviceStruct); (err != nil) != tt.wantErr {
				t.Errorf("AddFrpServiceConf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllFrpServiceConfsTypeAndName(t *testing.T) {
	tests := []struct {
		name                  string
		wantServiceBasicConfs []*model.FrpBasicService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotServiceBasicConfs := GetAllFrpServiceConfsTypeAndName(); !reflect.DeepEqual(gotServiceBasicConfs, tt.wantServiceBasicConfs) {
				t.Errorf("GetAllFrpServiceConfsTypeAndName() = %v, want %v", gotServiceBasicConfs, tt.wantServiceBasicConfs)
			}
		})
	}
}

func TestGetFrpServiceConfsByTag(t *testing.T) {
	type args struct {
		ServiceName   string
		serviceStruct interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetFrpServiceConfsByTag(tt.args.ServiceName, tt.args.serviceStruct); (err != nil) != tt.wantErr {
				t.Errorf("GetFrpServiceConfsByTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFrpServiceConfByName(t *testing.T) {
	type args struct {
		ServiceName   string
		serviceStruct interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetFrpServiceConfByName(tt.args.ServiceName, tt.args.serviceStruct); (err != nil) != tt.wantErr {
				t.Errorf("GetFrpServiceConfByName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteFrpServiceConf(t *testing.T) {
	type args struct {
		serviceName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteFrpServiceConf(tt.args.serviceName); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFrpServiceConf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateFrpServiceConf(t *testing.T) {
	type args struct {
		serviceName   string
		serviceStruct interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateFrpServiceConf(tt.args.serviceName, tt.args.serviceStruct); (err != nil) != tt.wantErr {
				t.Errorf("UpdateFrpServiceConf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
