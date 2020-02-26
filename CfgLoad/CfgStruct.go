package CfgLoad

import (
	"VisFusion/RegStruct"
)
type ServerCfg struct {

	// https Api url Cfg
	// Api url should be same as CaCert url
	HttpsServiceUrl string `ini:"HttpsServiceUrl"`

	//HttpsServiceIP string `ini:"HttpsServiceIP"`
	//Instant Data on Tcp ip(transfer on Tcp)
	//
	//InsDataIP string `ini:"InsDataIP"`

	// Set Instant Data ttl
	// When Instant Data
	InsDataTimeOut string `ini:"InsDataTimeOut"`

	// Instant can transfer by https when params of LinshiInsDataSwitch is set
	BackUpInsDataSwitch bool `ini:"BackUpInsDataSwitch"`

	// DownLoadPath means the physical path where is file set, clients download it from this path
	DownLoadPath string `ini:"DownLoadPath"`

	// clients upload purpose path
	VqUploadPath string `ini:"VqUploadPath"`

	// params of shellPath save the path where shell file save
	ShellPath string `ini:"shellPath"`

	//Save all kinds of default configs
	DevDefaultCfg []string `ini:"devDefaultCfg,omitempty,allowshadow"`
}

type DataSourceCfg struct {
	//Redis Configuration,set Redis Ip address and RedisPW
	RedisIP string `ini:"RedisIP"`
	RedisPW string `ini:"RedisPW"`

	//Mysql Configuration,set Redis Ip address and RedisPW
	MySqlIP   string `ini:"MySqlIP"`
	MySqlUser string `ini:"MySqlUser"`
	MySqlPW   string `ini:"MySqlPW"`
	MySqlDB   string `ini:"MySqlDB"`

}

//Keys(such as license,cert,private key) configurations
//type KeyCfg struct {
////
////	//Asymmetric encryption keys configurations
////	PubKeyPath string `ini:"PubKeyPath"`
////	PriKeyPath string `ini:"PriKeyPath"`
////
////	//tls keys and Cert configurations
////	HttpsKeyPath  string `ini:"HttpsKeyPath"`
////	HttpsCertPath string `ini:"HttpsCertPath"`
////	HttpsCaPath   string `ini:"HttpsCaPath"`
////}

//RaftNodes configurations
type RaftNodes struct {
	Nodes      []string `ini:"Nodes,omitempty,allowshadow"`
	SelectNode string   `ini:"SelectNode"`
}

type Cfg struct {
	//Mode Selected
	//DirectMode bool `ini:"DirectMode"`
	//DataMode   bool `ini:"DataMode"`

	//traefik port configurations
	TraefikPort string `ini:"traefikPort"`

	ServerCfg
	DataSourceCfg
	//KeyCfg
	RaftNodes

	//Saving the shell file content by Map
	ShellMap map[string]string
	//ShellList
	ShellList string
	//Reg Struct map
	DevCfgMap map[string]RegStruct.HttpsReg
	//RegStruct HttpsReg
	//RegDevCfg DevCfg
	Topicbool map[string]bool


}
