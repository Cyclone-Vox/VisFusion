package Director

import (


	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/valyala/fasthttp"
	"strings"
)

func ping(ctx *fasthttp.RequestCtx) {

	tk := ctx.Request.Header.Peek("token")
	r := CfgLoad.Conf.RedisPool.Get()
	var resp []byte
	//从redis中获取mc

	mc, err := redis.String(r.Do("GET", string(tk)))
	CheckError(err)
	sn_pid, err := redis.String(r.Do("GET", mc))
	CheckError(err)
	sn_pidSplit := strings.Split(sn_pid, "#")







	if v, ok := mission.JobCfgMap.Load(sn_pidSplit[0] + "#job"); ok {
		get := v.(cfgs.MValue)

		if get.Change {
			var mcfg CfgLoad.HttpsReg
			var jsoncfg CfgLoad.DevCfg
			//instant job list
			var list []CfgLoad.Commandlist
			mcfg.STATUS = 1
			mcfg.CFG.Post.Mp2tIndex.Switch = true
			mcfg.CFG.Post.Mp2tIndex.Cycle = 60
			mcfg.CFG.Post.Mp2tEvent.Switch = true
			mcfg.CFG.Post.HlsEvent.Switch = true
			mcfg.CFG.Post.HlsIndex.Switch = true
			mcfg.CFG.Post.HlsIndex.Cycle = 60
			mcfg.CFG.Post.IgmpEvent.Switch = true
			mcfg.CFG.Post.VqIndex.Switch = true
			mcfg.CFG.Post.VqIndex.Cycle = 60
			mcfg.CFG.Post.RtspEvent.Switch = true
			mcfg.CFG.Post.DeviceInfo.Cycle = 60
			mcfg.CFG.Post.DeviceInfo.Switch = true
			mcfg.CFG.Post.HlsEventPull.Switch=true
			mcfg.CFG.Post.AlarmData.Switch=true
			mcfg.CFG.Post.Posturl = "https://" + CfgLoad.Conf.ApiHttps + "/upload"
			mcfg.CFG.Post.Uploadurl = "https://:" + cfgs.Conf.ApiHttps + "/fileupload"

			//get the ShellV in map
			if get.ShellV != nil {

				//choose check LinshiConsDataSwitch
				if cfgs.Conf.LinshiInsDataSwitch {
					mcfg.CFG.Pong.Tcp = ""
					mcfg.CFG.Pong.LinshiTCP = "https://" + cfgs.Conf.ApiHttps + "/insdata"
				}else {
					mcfg.CFG.Pong.Tcp = cfgs.Conf.InsDataIP
					mcfg.CFG.Pong.LinshiTCP = ""
				}

				list = get.ShellV

				//del the shell data in get.ShellV and redis
				get.ShellV = nil
				_, err = r.Do("DEL", sn_pidSplit[0]+"#job#0")
				tool.CheckError(err)

			}
			// get the Cfg in map
			if get.CfgV != nil {
				err = json.Unmarshal(get.CfgV, &jsoncfg)
				tool.CheckError(err)
				mcfg.CFG.Pull = jsoncfg.Pull
				mcfg.CFG.Alarm.Alarmthresholds = jsoncfg.Alarm.Alarmthresholds

				tool.CheckError(err)

				get.CfgV = nil
			}

			// heartbeat and dataPost setting
			if get.PostV != nil {

				mcfg.CFG.Pong.Pongurl = "https://" + CfgLoad.Conf.HttpsServiceUrl + "/ping"
				mcfg.CFG.Pong.Cycle = 10

				//choose check LinshiConsDataSwitch
				if cfgs.Conf.LinshiInsDataSwitch {
					mcfg.CFG.Pong.Tcp = ""
					mcfg.CFG.Pong.LinshiTCP = "https://" + "CfgLoad.Conf.HttpsServiceUrl" + "/insdata"
				}else {
					mcfg.CFG.Pong.Tcp = cfgs.Conf.InsDataIP
					mcfg.CFG.Pong.LinshiTCP = ""
				}

				mcfg.CFG.Pong.Tcptimeout = 10

				var cmd cfgs.Commandlist

				//get each instant command into list
				for i := 0; i < len(get.PostV); i++ {
					//clear := 0
					//
					cmd.Commandtype = get.PostV[i]
					cmd.Command = ""
					list = append(list, cmd)


				}

				//remove "" in []string
				for i:=0;i<len(get.PostV);i++{
					if get.PostV[i]==""{
						//use append remove the member in get.PostV
						get.PostV=append(get.PostV[:i],get.PostV[i+1:]...)
						i--
					}
					if len(get.PostV) ==0{
						get.PostV=nil
					}
				}
			}

			//if there is null data in jobcfgmap ,set the 'change' to false
			if (get.PostV == nil) && (get.CfgV == nil) && (get.ShellV == nil) {
				get.Change = false
			}

			mcfg.CFG.Pong.Commandlist = list
			mission.JobCfgMap.Store(sn_pidSplit[0]+"#job", get)
			resp, err = json.Marshal(mcfg)
		}

	}

	r.Close()
	fmt.Println(string(ctx.Path()), string(tk), "sn:", sn_pidSplit[0]+"Req.Head:", string(tk), "Resp:", string(resp))
	ctx.Response.SetBody(resp)
	ctx.Request.ConnectionClose()
	ctx.Response.ConnectionClose()
	ctx.SetConnectionClose()
	return
}
