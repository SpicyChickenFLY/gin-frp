package model

type ServerInfo struct {
	address string
	port    int
	token   string
}

type FrpCommonService struct {
	Name string
	Type string
}

type FrpTcpService struct {
	FrpCommonService
	LocalIP   string
	LocalPort int
	SecretKey string
}

type FrpFilePlugin struct {
	Name         string
	LocalPath    string
	StripPrefix  string
	HttpUser     string
	HttpPassword string
}

type FrpFileService struct {
	FrpCommonService
	FrpFilePlugin
	RemotePort int
}

type FrpHttpsPlugin struct {
	Name              string
	LocalAddress      string
	CertificatePath   string
	KeyPath           string
	HostHeaderRewrite string
	HeaderXFromWhere  string
}

type FrpHttpsService struct {
	FrpCommonService
	FrpHttpsPlugin
}
