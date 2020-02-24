package RegStruct
import "time"

type MValue struct {
	Change bool
	CfgV   []byte
	PostV  []string
	ShellV []Commandlist `json:"commandlist"`
}

type HttpsReg struct {
	STATUS int    `json:"status"`
	CFG    DevCfg `json:"CfgFile"`
	TOKEN  string `json:"token"`
	SN     string `json:"sn"`
	PID    string `json:"pid"`
}

type DevCfg struct {
	Startup []startup   `json:"startup"` // 启动配置
	Pull    []Pl        `json:"pull"`    // 拉流配置
	Post    Post        `json:"post"`    // 上报配置
	Pong    Pong        `json:"pong"`    // 心跳配置
	Alarm   alarmconfig `json:"alarm"`   // 告警配置
	Command []command   `json:"command"` // 实时命令
}

// 启动程序
type startup struct {
	Name    string            `json:"name"`    // 程序名
	Downurl string            `json:"downurl"` // 程序下载地址
	Params  map[string]string `json:"params"`  // 启动参数
	Config  string            `json:"config"`  // 程序配置文件
	//Licensekey string `json:"license"` //
	//License string `json:"license"` // License 存储路径 只要这个字不为空，就需要去通过 key去获取
}

// 拉流配置
type Pl struct {
	Type     string   `json:"type"` // 类型，组播地址 或者HLS
	Url      []string `json:"url"`  // 拉流地址
	App      string   `json:"app"`  // 拉流软件
	Filename string   `json:"filename"`
	Isclear  bool     `json:"isclear"`
}
type Params struct {
	CMCC_Sn    string `json:"cmcc_sn"`
	CMCC_Token string `json:"cmcc_token"`
	CMCC_Pid   string `json:"cmcc_pid"`
}

// 心跳规则
type Pong struct {
	Cycle       time.Duration `json:"cycle"` // 与平台通信心跳周期
	Pongurl     string        `json:"pongurl"`
	ReLogin     bool          `json:"relogin"`
	LinshiTCP   string        `json:"linshi_tcp"`
	Tcp         string        `json:"tcp"`
	Tcptimeout  time.Duration `json:"tcptimeout"`  // tcp超时
	Commandlist []Commandlist `json:"commandlist"` // command 列表

}

// 0、shell指令
// 1、hls_index 实时数据上报开启
// 2、hls_event 实时数据上报开启
// 3、Mp2t_index 实时数据上报开启
// 4、Mp2t_event 实时数据上报开启
// 5、Vq实时数据上报开启
// 6、rtsp_event
// 7、igmp_event

type Commandlist struct {
	Commandtype string            `json:"commandtype"` //  实时命令类型 0、shell脚本类型 1...、其他类型
	Command     string            `json:"command"`     // 如果是0，这里是具体的脚本内容
	Params      map[string]string `json:"params"`	   //
}

// 告警配置

type alarmconfig struct {
	Alarmthresholds map[string][]map[string]interface{} `json:"alarmthresholds"` // 告警阈值结构
}

// 命令
type command struct {
	Command  string `json:"command"`  // 命令
	Response string `json:"response"` // 执行后是否响应
}

// 上报数据配置
type Post struct {
	Posturl      string `json:"posturl"` // 数据上报地址
	Uploadurl    string `json:"uploadurl"`
	DeviceInfo   `json:"device_info"`    // 设备信息
	Mp2tIndex    `json:"mp_2_t_index"`   // Mp2t index
	HlsIndex     `json:"hls_index"`      // Hls index
	VqIndex      `json:"vq_index"`       // Vq index
	RtspEvent    `json:"rtsp_event"`     // Rtsp event
	Mp2tEvent    `json:"mp_2_t_event"`   // mp2t event
	HlsEvent     `json:"hls_event"`      // hls event
	IgmpEvent    `json:"igmp_event"`     // igmp event
	HlsEventPull `json:"hls_event_pull"` //hls_event_pull
	AlarmData	 `json:"alarm_data"`
}

type HlsEventPull struct {
	Switch bool `json:"switch"` // 开启上报
}

type Mp2tIndex struct {
	Cycle  time.Duration `json:"cycle"`  // 上报汇聚周期
	Switch bool          `json:"switch"` // 开启上报
}
type HlsIndex struct {
	Cycle  time.Duration `json:"cycle"`  // 上报汇聚周期
	Switch bool          `json:"switch"` // 开启上报
}
type VqIndex struct {
	Cycle  time.Duration `json:"cycle"`  // 上报汇聚周期
	Switch bool          `json:"switch"` // 开启上报
}
type RtspEvent struct {
	Switch bool `json:"switch"` // 开启上报
}
type Mp2tEvent struct {
	Switch bool `json:"switch"` // 开启上报
}
type HlsEvent struct {
	Switch bool `json:"switch"` // 开启上报
}
type IgmpEvent struct {
	Switch bool `json:"switch"` // 开启上报
}
type DeviceInfo struct {
	Cycle  time.Duration `json:"cycle"`  // 上报汇聚周期
	Switch bool          `json:"switch"` // 开启上报
}
type  AlarmData struct {
	Switch bool `json:"switch"` // 开启上报
}