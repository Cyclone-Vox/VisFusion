package ReversProxy

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"

	"sync"
)

type proxy struct {
	//key=Pattern;value=targetHost
	hostTarget map[string]string
	//key=Pattern;value=targetHost
	ProxyCache map[string]*httputil.ReverseProxy
}

//ProxySetUp(..) is external Func for main() Call
func ProxySetUp(ctx context.Context, ProxyMap *sync.Map, Port *string, caCertPath *string, CertPath *string, KeyPath *string) {
	p := &proxy{}
	p.proxySetUp(ctx, ProxyMap, *Port, *caCertPath, *CertPath, *KeyPath)
}

//this function will call a trail of functions to set up proxy serve
func (p *proxy) proxySetUp(ctx context.Context, ProxyMap *sync.Map, Port string, caCertPath string, CertPath string, KeyPath string) {
	p.hostTarget = make(map[string]string)
	p.loadProxyMap(ProxyMap)
	p.setUpProxy()
	//Creat Https Listener And SetUp Proxy Server
	lnTls, err := newTlsLn(Port, caCertPath, CertPath, KeyPath)
	CheckError(err)

	http.HandleFunc("/", p.HandlerWithCache)
	//wait Context done,and reset up proxy function
	//http.Serve(lnTls, nil)
	srv := &http.Server{Handler: nil}
	go ServerClose(ctx, srv)
	srv.Serve(lnTls)

	select {
	case <-ctx.Done():
		lnTls.Close()
		CheckError(err)
	}
}

//Range the ProxyMap，set the configurations it in hostTarget
func (p *proxy) loadProxyMap(ProxyMap *sync.Map) {
	ProxyMap.Range(func(key, value interface{}) bool {
		p.hostTarget[key.(string)] = value.(string)
		return true
	})
}

//Set Up the Proxy handler and save it in ProxyCache as cache
func (p *proxy) setUpProxy() {
	for pattern, targetHost := range p.hostTarget {
		remoteUrl, err := url.Parse(targetHost)
		CheckError(err)
		proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
		p.ProxyCache[pattern] = proxy
	}
}

//Main Handler for Proxy Serve, it will find ProxyHandler in the Map saves Cache
func (p *proxy) HandlerWithCache(w http.ResponseWriter, r *http.Request) {

	pattern := r.URL.Path
	if ProxyFn, ok := p.ProxyCache[pattern]; ok {
		ProxyFn.ServeHTTP(w, r)
	} else {
		_, err := w.Write([]byte("403: Host forbidden" + r.URL.String() ))
		CheckError(err)
	}

}

func ServerClose(ctx context.Context, server *http.Server) {
	select {
	case <-ctx.Done():
		server.SetKeepAlivesEnabled(false)
		server.Shutdown(ctx)
	}
}
