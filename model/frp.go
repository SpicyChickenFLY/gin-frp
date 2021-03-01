package model

const (
	frpFilePluginName  = "static_file"
	frpHTTPSPluginName = "https2http"
)

// ServerInfo is a frp server info struct
type ServerInfo struct {
	Address string `ini:"server_addr"`
	Port    int    `ini:"server_port"`
	Token   string `ini:"token"`
}

// FrpTCPService is a tcp type frp service struct
type FrpTCPService struct {
	ServiceName string
	ServiceType string `ini:"type"`
	LocalIP     string `ini:"local_ip"`
	LocalPort   int    `ini:"local_port"`
	SecretKey   string `ini:"sk"`
}

// NewFrpTCPService initialize new FrpTCPService struct
func NewFrpTCPService(
	name, serviceType string,
	localIP string, localPort int, secretKey string) *FrpTCPService {
	return &FrpTCPService{
		ServiceName: name,
		ServiceType: serviceType,
		LocalIP:     localIP,
		LocalPort:   localPort,
		SecretKey:   secretKey,
	}
}

// FrpFileService is a file type frp service struct
type FrpFileService struct {
	ServiceName  string
	ServiceType  string `ini:"type"`
	PluginName   string `ini:"plugin"`
	LocalPath    string `ini:"plugin_local_path"`
	StripPrefix  string `ini:"plugin_strip_prefix"`
	HTTPUser     string `ini:"plugin_http_user"`
	HTTPPassword string `ini:"plugin_http_passwd"`
	RemotePort   int    `ini:"remote_port"`
}

// NewFrpFileService initialize new FrpFileService struct
func NewFrpFileService(
	serviceName, serviceType string,
	pluginPath, pluginPrefix, pluginUser, pluginPwd string,
	remotePort int) *FrpFileService {
	return &FrpFileService{
		ServiceName:  serviceName,
		ServiceType:  serviceType,
		PluginName:   frpFilePluginName,
		LocalPath:    pluginPath,
		StripPrefix:  pluginPrefix,
		HTTPUser:     pluginUser,
		HTTPPassword: pluginPwd,
		RemotePort:   remotePort,
	}
}

// FrpHTTPSService is a https type frp service struct
type FrpHTTPSService struct {
	ServiceName       string
	ServiceType       string `ini:"type"`
	PluginName        string `ini:"plugin"`
	LocalAddress      string `ini:"plugin_local_addr"`
	CertificatePath   string `ini:"plugin_crt_path"`
	KeyPath           string `ini:"plugin_key_path"`
	HostHeaderRewrite string `ini:"plugin_host_header_rewrite"`
	HeaderXFromWhere  string `ini:"plugin_header_X-From-Where"`
}

// NewFrpHTTPSService initialize new FrpHTTPSService struct
func NewFrpHTTPSService(
	serviceName, serviceType string,
	pluginAddr, pluginCrt, pluginKey,
	pluginRewrite, pluginWhere string,
	remotePort int) *FrpHTTPSService {
	return &FrpHTTPSService{
		ServiceName:       serviceName,
		ServiceType:       serviceType,
		PluginName:        frpHTTPSPluginName,
		LocalAddress:      pluginAddr,
		CertificatePath:   pluginCrt,
		KeyPath:           pluginKey,
		HostHeaderRewrite: pluginRewrite,
		HeaderXFromWhere:  pluginWhere,
	}
}
