package model

// ServerInfo is a frp server info struct
type ServerInfo struct {
	address string
	port    int
	token   string
}

// FrpCommonService is a general frp service struct
type FrpCommonService struct {
	Name string
	Type string
}

// FrpTCPService is a tcp type frp service struct
type FrpTCPService struct {
	FrpCommonService
	LocalIP   string
	LocalPort int
	SecretKey string
}

// NewFrpTCPService initialize new FrpTCPService struct
func NewFrpTCPService(
	name, serviceType string,
	localIP string, localPort int, secretKey string) *FrpTCPService {
	return &FrpTCPService{
		FrpCommonService: FrpCommonService{
			Name: name,
			Type: serviceType,
		},
		LocalIP:   localIP,
		LocalPort: localPort,
		SecretKey: secretKey,
	}
}

// FrpFilePlugin is a file frp plugin struct
type FrpFilePlugin struct {
	Name         string
	LocalPath    string
	StripPrefix  string
	HTTPUser     string
	HTTPPassword string
}

// FrpFileService is a file type frp service struct
type FrpFileService struct {
	FrpCommonService
	FrpFilePlugin
	RemotePort int
}

// NewFrpFileService initialize new FrpFileService struct
func NewFrpFileService(
	serviceName, serviceType string,
	pluginName, pluginPath, pluginPrefix, pluginUser, pluginPwd string,
	remotePort int) *FrpFileService {
	return &FrpFileService{
		FrpCommonService: FrpCommonService{
			Name: serviceName,
			Type: serviceType,
		},
		FrpFilePlugin: FrpFilePlugin{
			Name:         pluginName,
			LocalPath:    pluginPath,
			StripPrefix:  pluginPrefix,
			HTTPUser:     pluginUser,
			HTTPPassword: pluginPwd,
		},
		RemotePort: remotePort,
	}
}

// FrpHTTPSPlugin is a https frp plugin struct
type FrpHTTPSPlugin struct {
	Name              string
	LocalAddress      string
	CertificatePath   string
	KeyPath           string
	HostHeaderRewrite string
	HeaderXFromWhere  string
}

// FrpHTTPSService is a https type frp service struct
type FrpHTTPSService struct {
	FrpCommonService
	FrpHTTPSPlugin
}

// NewFrpHTTPSService initialize new FrpHTTPSService struct
func NewFrpHTTPSService(
	serviceName, serviceType string,
	pluginName, pluginAddr, pluginCrt, pluginKey,
	pluginRewrite, pluginWhere string,
	remotePort int) *FrpHTTPSService {
	return &FrpHTTPSService{
		FrpCommonService: FrpCommonService{
			Name: serviceName,
			Type: serviceType,
		},
		FrpHTTPSPlugin: FrpHTTPSPlugin{
			Name:              pluginName,
			LocalAddress:      pluginAddr,
			CertificatePath:   pluginCrt,
			KeyPath:           pluginKey,
			HostHeaderRewrite: pluginRewrite,
			HeaderXFromWhere:  pluginWhere,
		},
	}
}
