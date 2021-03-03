package model

const (
	frpFilePluginName  = "static_file"
	frpHTTPSPluginName = "https2http"
)

// ServerInfo is a frp server info struct
type ServerInfo struct {
	Address string `ini:"server_addr" json:"Address"`
	Port    int    `ini:"server_port" json:"Port"`
	Token   string `ini:"token" json:"Token"`
}

// FrpBasicService is a service configure struct
type FrpBasicService struct {
	Name string `json:"ServiceName"`
	Type string `json:"ServiceType"`
}

// FrpTransportService is a TCP/UDP type frp service struct
type FrpTransportService struct {
	ServiceName string `json:"ServiceName"`
	ServiceType string `ini:"type" json:"ServiceType"`
	LocalIP     string `ini:"local_ip" json:"LocalIP"`
	LocalPort   int    `ini:"local_port" json:"LocalPort"`
}

// FrpSecureTransportService is a STCP/SUDP/XTCP type frp service struct
type FrpSecureTransportService struct {
	ServiceName string `json:"ServiceName"`
	ServiceType string `ini:"type" json:"ServiceType"`
	LocalIP     string `ini:"local_ip" json:"LocalIP"`
	LocalPort   int    `ini:"local_port" json:"LocalPort"`
	SecretKey   string `ini:"sk" json:"SecretKey"`
}

// FrpFileService is a file type frp service struct
type FrpFileService struct {
	ServiceName  string `json:"ServiceName"`
	ServiceType  string `ini:"type" json:"ServiceType"`
	PluginName   string `ini:"plugin"`
	LocalPath    string `ini:"plugin_local_path"`
	StripPrefix  string `ini:"plugin_strip_prefix"`
	HTTPUser     string `ini:"plugin_http_user"`
	HTTPPassword string `ini:"plugin_http_passwd"`
	RemotePort   int    `ini:"remote_port"`
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
