package CfgLoad

import (
	"database/sql"
	"github.com/garyburd/redigo/redis"
)
type ServerCfg struct {

	// https Api url Cfg
	// Api url should be same as CaCert url
	HttpsServiceUrl string `ini:"HttpsServiceUrl"`

	HttpsServiceIP string `ini:"HttpsServiceIP"`
	//Instant Data on Tcp ip(transfer on Tcp)
	//
	InsDataIP string `ini:"InsDataIP"`

	// Set Instant Data ttl
	// When Instant Data
	InsDataTimeOut string `ini:"InsDataTimeOut"`

	// Instant can transfer by https when params of LinshiInsDataSwitch is set
	LinshiInsDataSwitch bool `ini:"LinshiInsDataSwitch"`

	// DownLoadPath means the physical path where is file set, clients download it from this path
	DownLoadPath string `ini:"DownLoadPath"`

	// clients upload purpose path
	VqUploadPath string `ini:"VqUploadPath"`

	// params of shellPath save the path where shell file save
	shellPath string `ini:"shellPath"`

	//Save all kinds of default configs
	devDefaultCfg []string `ini:"devDefaultCfg,allowshadow"`
}


type Cfg struct {
	//Mode Selected
	DirectMode bool `ini:"DirectMode"`
	DataMode   bool `ini:"DataMode"`

	//traefik port configurations
	TraefikPort string `ini:"traefikPort"`

	ServerCfg
	DataSourceCfg
	KeyCfg
	RaftNodes

	//Saving the shell file content by Map
	ShellMap map[string]string
	//ShellList
	ShellList string
	//Reg Struct map
	DevDefaultCfg map[string]HttpsReg
	//RegStruct HttpsReg
	//RegDevCfg DevCfg
	Topicbool map[string]bool

	//NodeMap Save {NodeID: [Node Client IP,Node Peer IP]}
	NodesMap map[string][]string
	//NodeList Save All Client IP
	NodesClientList []string

	//Etcd Service Client
	Clientdis *EtcdFunc.ClientDis

	RedisPool *redis.Pool
	//One Redis Pool for Redis subscribe
	SubRedisPool *redis.Pool
	Db           *sql.DB
}
