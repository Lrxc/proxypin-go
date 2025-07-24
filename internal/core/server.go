package core

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	resources "proxypin-go/assets"
	"proxypin-go/internal/config"
)

func StartServer() {
	proxy := goproxy.NewProxyHttpServer()
	//proxy.Verbose = true
	//proxy.AllowHTTP2 = true
	proxy.Logger = SilentLog{}

	// 设置CA证书
	if err := SetCA(); err != nil {
		log.Fatalf("https certificate err: %v", err)
	}

	// 处理HTTPS连接
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	// 请求
	proxy.OnRequest().DoFunc(ReqHandler)
	// 相应
	proxy.OnResponse().DoFunc(ResHandler)

	addr := fmt.Sprintf("%s:%s", config.Conf.System.Host, config.Conf.System.Port)
	fmt.Println("server listen: ", addr)
	fmt.Println("start successful!")
	log.Fatal(http.ListenAndServe(addr, proxy))
}

func SetCA() error {
	certFile, _ := resources.ReadByte("server.crt")
	keyFile, _ := resources.ReadByte("server.key")

	goproxyCa, err := tls.X509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}
	if goproxyCa.Leaf, err = x509.ParseCertificate(goproxyCa.Certificate[0]); err != nil {
		return err
	}
	goproxy.GoproxyCa = goproxyCa

	//goproxy.OkConnect = &goproxy.ConnectAction{Action: goproxy.ConnectAccept, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	//goproxy.MitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	//goproxy.HTTPMitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectHTTPMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	//goproxy.RejectConnect = &goproxy.ConnectAction{Action: goproxy.ConnectReject, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}

	return nil
}

type SilentLog struct{}

func (e SilentLog) Printf(format string, v ...any) {
}
