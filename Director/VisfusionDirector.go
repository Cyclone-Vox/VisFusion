package Director

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/valyala/fasthttp"
	"net"
)


func Director(ctx context.Context, ip string, RedisPool *redis.Pool) {

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		r := RedisPool.Get()
		defer r.Close()
		CertCheckOrReg(ctx)
		//if Path not equals "devreg",it Means this device's request is licgen or heartbeat
		return
	}

	ln, err := net.Listen("tcp4", ":"+ip)
	CheckError(err)
	ser:=&fasthttp.Server{Handler:requestHandler}
	go ServerClose(ctx,ser)
	ser.Serve(ln)
	//wait context cancel,then close the listener and set up a new http service
	select {
	case <-ctx.Done():
		ln.Close()
		Director(ctx, ip, RedisPool)
	}

}
func ServerClose(ctx context.Context, server *fasthttp.Server) {
	select {
	case <-ctx.Done():
		server.Shutdown()
	}
}