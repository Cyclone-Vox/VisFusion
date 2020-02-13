package ReversProxy

import (
	"context"
	"log"

	"net/http"
	"net/http/httputil"
	"net/url"
)

type ExtensionCfg struct{
	url string
	pattern string
}


var hostTarget = make(map[string]*ExtensionCfg)
var hostProxy map[string]*httputil.ReverseProxy


func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host

	// 直接从缓存取出
	if fn, ok := hostProxy[host]; ok {
		fn.ServeHTTP(w, r)
		return
	}

	// 检查域名白名单
	if target, ok := hostTarget[host]; ok {
		remoteUrl, err := url.Parse(target)
		if err != nil {
			log.Println("target parse fail:", err)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
		hostProxy[host] = proxy
		proxy.ServeHTTP(w, r)
		return
	}
	w.Write([]byte("403: Host forbidden " + host))
}

func ReversProxySetUp(ctx context.Context, Port string, caCertPath string, CertPath string, KeyPath string) {
	if ln, err := newTlsLn(Port, caCertPath, CertPath, KeyPath); err==nil {
		for i:=0;i<1;i++{
			http.Handle("/", h)
		}

		http.Serve(ln,)
	}
}
