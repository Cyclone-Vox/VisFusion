package Director

import (
	"github.com/valyala/fasthttp"
)

func TokenError(ctx *fasthttp.RequestCtx) {
	//if string(ctx.Path())=="/upload" {
	//	fmt.Println(string(ctx.Path()),string(ctx.Request.Header.Peek("token")),"Error")
	//
	//}else {
	//	fmt.Println(string(ctx.Path()),string(ctx.Request.Header.Peek("token")),"Error",string(ctx.Request.Body()))
	//
	//}
	//
	//var sss CfgLoad.HttpsReg
	//sss.CFG.Pong.ReLogin=true
	//b,err:=json.Marshal(sss)
	//CheckError(err)
	//ctx.SetBody(b)
	//ctx.Request.ConnectionClose()
	//ctx.SetConnectionClose()
	//fmt.Println(string(b))
	//return
}