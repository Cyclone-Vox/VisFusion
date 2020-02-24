package Director

import (
	"VisFusion/DataConn"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/spaolacci/murmur3"
	"github.com/valyala/fasthttp"
	"strconv"
	"strings"
	"time"
)


//CertCheckOrReg负责完成认证注册功能,
//注册成功后会从

//var HttpReg CfgLoad.HttpsReg

func CertCheckOrReg(ctx *fasthttp.RequestCtx) {
	Req := ctx.PostBody()
	var jsonget map[string]string
	var resp []byte
	err := json.Unmarshal(Req, &jsonget)
	CheckError(err)
	r := DataConn.RedisPool.Get()
	defer r.Close()

	//判断是否含有machineCode,不含有machineCode时返回
	mc, hasmc := HasKeyAndGet(jsonget, "machineCode")
	fmt.Println(string(ctx.Path()), string(ctx.Request.Header.Peek("token")), string(Req), "hasmc:", hasmc)
	if !hasmc {
		ctx.Error("this request don't Send Mc", fasthttp.StatusForbidden)
		ctx.Request.ConnectionClose()
		ctx.SetConnectionClose()
		return
	}

	//当设备内容带有machineCode时
	fmt.Println(mc)
	str, err := redis.String(r.Do("GET", mc))
	if err != nil {
		fmt.Println(err)
		//判断是否有此键
		if err.Error() == "redigo: nil returned" {
			fmt.Println("mc:", mc, "初次设备申请平台许可")
			_, err = redis.String(r.Do("SET", mc, ""))
			CheckError(err)

			_, err = CfgLoad.Conf.Db.Exec("INSERT into "+CfgLoad.Conf.MySqlDB+".MACHINE_CODE_LIST(MACHINE_CODE,IP) value (?,?)", mc, ctx.RemoteIP().String())
			CheckError(err)
			resp, err = json.Marshal(CfgLoad.HttpsReg{
				STATUS: 0,
			})
			CheckError(err)
			fmt.Println(string(resp))
			ctx.Response.SetBody(resp)
			ctx.Request.ConnectionClose()
			ctx.SetConnectionClose()

			return
		}

	}

	//即平台工作人员是否同意设备的注册
	switch str {
	case "":
		fmt.Println("平台还未同意注册")
		resp, err = json.Marshal(CfgLoad.HttpsReg{
			STATUS: 0,
		})
		CheckError(err)
		fmt.Println(string(resp))
		ctx.Response.SetBody(resp)
		ctx.Request.ConnectionClose()
		ctx.SetConnectionClose()
		r.Close()
		return

	default:

		println("已注册设备", str)
		//////为设备分配token
		hash := murmur3.Sum32([]byte(mc))
		hashstr := "token#" + strconv.Itoa(int(hash)) + "#" + mc
		fmt.Println(hashstr)
		MsgSplit := strings.Split(str, "#")

		///////////返回数据的组装
		//判断是否需要初始化

		if RespawnCfg(MsgSplit[0]) {
			//如果需要初始化，则给RegStruct赋回初始化设备的内容，并且将设置改为不需要初始化
			cfgs.RegStruct.CFG = cfgs.RegDevCfg
			_, err := r.Do("SET", MsgSplit[0]+"#Respawn", "1")
			CheckError(err)
		} else {
			//如果需要不需要初始化，则不给RegStruct赋回初始化设备的内容
			cfgs.RegStruct.CFG.Pull = nil
		}

		//choose check LinshiConsDataSwitch
		if CfgLoad.Conf.LinshiInsDataSwitch {
			cfgs.RegStruct.CFG.Pong.Tcp = ""
			cfgs.RegStruct.CFG.Pong.LinshiTCP = "https://" + cfgs.Conf.ApiHttps + "/insdata"
		} else {
			cfgs.RegStruct.CFG.Pong.Tcp = cfgs.Conf.InsDataIP
			cfgs.RegStruct.CFG.Pong.LinshiTCP = ""
		}

		cfgs.RegStruct.TOKEN = hashstr
		cfgs.RegStruct.SN = MsgSplit[0]
		cfgs.RegStruct.PID = MsgSplit[1]
		resp, err = json.Marshal(cfgs.RegStruct)
		CheckError(err)

		///////////////////////
		_, err = r.Do("SET", hashstr, mc)
		CheckError(err)
		//成功登陆注册后维护状态信息
		_, err = CfgLoad.Conf.Db.Exec("INSERT into "+CfgLoad.Conf.MySqlDB+".STB_EVENT_LOG(SN,PROBE_ID,TIME,TYPE) value (?,?,?,?)", cfgs.RegStruct.SN, cfgs.RegStruct.PID, time.Now().Format("2006-01-02 15:04:05"), 1)
		CheckError(err)
		_, err = CfgLoad.Conf.Db.Exec("UPDATE "+CfgLoad.Conf.MySqlDB+".VISQUAL_PROBES set ACTIVE=? where SN=?", "Y", cfgs.RegStruct.SN)
		CheckError(err)
		_, err = r.Do("EXPIRE", hashstr, 60)
		CheckError(err)
		var job cfgs.MValue
		mission.JobCfgMap.Store(MsgSplit[0]+"#job", job)

	}

	//LogSet.Info("Https接口发送的注册请求数据：",fmt.Sprintf("%s",resp))

	fmt.Println(string(ctx.Path()), string(ctx.Request.Header.Peek("token")), "Req:", string(Req), "Resp", string(resp), "hasmc:", hasmc)
	ctx.Response.SetBody(resp)
	ctx.Request.ConnectionClose()
	ctx.SetConnectionClose()
	r.Close()

	return

}

//判断设备Http请求json中是否带有某个Key
func HasKeyAndGet(jsonmap map[string]string, key string) (string, bool) {
	if v, ok := jsonmap[key]; ok {
		return v, true
	} else {
		return "NoExistKey", false
	}
}

//判断设备是否需要初始化配置
func RespawnCfg(sn string) bool {
	r := CfgLoad.Conf.RedisPool.Get()
	v, err := redis.String(r.Do("GET", sn+"#Respawn"))
	//判断是否为不存在的错误
	if err != nil {
		fmt.Println(err)
		if err.Error() == "redigo: nil returned" {
			_, err := r.Do("SET", sn+"#Respawn", "0")
			CheckError(err)
			r.Close()
			return true
		}
	} else {
		r.Close()
		if v == "0" {
			return true
		} else {
			return false
		}

	}

	return true
	//键值为0的时候需要刷新初始化
	//键值为非0的时候不需要刷新初始化
}
