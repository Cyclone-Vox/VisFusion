package Director

import (
	"context"
	"github.com/valyala/fasthttp"
)

func Director(ctx context.Context, ip string) {

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())
		tk := ctx.Request.Header.Peek("token")
		r := CfgLoad.Conf.RedisPool.Get()
		defer r.Close()
		//Check token in redis

		// if it not raises an error,it will
		//Select Data Mode Open
		switch path {
		case "devreg":
			CertCheckOrReg(ctx)
			return
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
	}

	
}

