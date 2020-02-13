package Director

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/valyala/fasthttp"
	"net"
)

func Director(ctx context.Context, ip string, RedisPool *redis.Pool) {

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())
		tk := ctx.Request.Header.Peek("token")
		r := RedisPool.Get()
		defer r.Close()

		switch path {

		//if Path equals "devreg",it Means this device's request is login or register
		case "devreg":
			CertCheckOrReg(ctx)
			return

		//if Path not equals "devreg",it Means this device's request is licgen or heartbeat
		default:
			if err := TokenCheck(string(tk), r); err == nil {
				switch path {
				case "/licgen":
					licgen(ctx)
				case "/ping":
					ping(ctx)
				default:
					ctx.Error("Unsupported path", fasthttp.StatusNotFound)
				}
			} else {
				TokenError(ctx)
			}

			return
		}
		return
	}
	ln, err := net.Listen("tcp4", ":"+ip)
	if err != nil {
		panic(fasthttp.Serve(ln, requestHandler))
	}

	select {
	case <-ctx.Done():
		ln.Close()
		Director(ctx, ip, RedisPool)
	}

}
