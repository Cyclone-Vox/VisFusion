package ReversProxy

import (
	"context"
	"sync"

	"log"

	"net/http"
	"net/http/httputil"
	"net/url"
)

type ExtensionCfg struct{
	url string
	pattern string
}
type proxy struct{
	hostTarget  map[string]ExtensionCfg
}

func (p *proxy)ProxySetUp(ctx context.Context, ProxyMap *sync.Map,Port string, caCertPath string, CertPath string, KeyPath string)  {
	p.hostTarget= make(map[string]ExtensionCfg)
	p.loadProxyMap(ProxyMap)
	lnTls, err := newTlsLn(Port, caCertPath, CertPath, KeyPath)
	CheckError(err)
	http.Handle("/",p)
	http.Serve(lnTls,p)

	select {
	case <-ctx.Done():
		p.ProxySetUp(ctx,ProxyMap,Port,caCertPath,CertPath,KeyPath)
	}
}

func (p *proxy)loadProxyMap(ProxyMap *sync.Map)  {
	ProxyMap.Range(func(key, value interface{}) bool {
		p.hostTarget[key.(string)]=value.(ExtensionCfg)
		return true
	})
}

func (p *proxy)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host

	// 直接从缓存取出
	if fn, ok := hostProxy[host]; ok {
		fn.ServeHTTP(w, r)
		return
	}

	// 检查域名白名单
	if target, ok := hostTarget[host]; ok {
		remoteUrl, err := url.Parse(c.url)
		if err != nil {
			log.Println("target parse fail:", err)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
		hostProxy[host] = proxy
		proxy.ServeHTTP(w, r)
		return

		w.Write([]byte("403: Host forbidden " + host))
	}

}
//func ReversProxySetUp(ctx context.Context, Port string, caCertPath string, CertPath string, KeyPath string) {
//
//		ln, err := newTlsLn(Port, caCertPath, CertPath, KeyPath)
//		CheckError(err)
//		p:=&proxy{}
//		http.Handle("/",p)
//		http.Serve(ln,p)
//
//}
