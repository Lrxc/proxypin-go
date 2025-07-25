package server

import (
	"fmt"
	"github.com/elazarl/goproxy"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"proxypin-go/internal/config"
	"strings"
)

func ReqHandler(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	log.Info("req: %s %s\n", r.Method, r.URL.String())

	for _, s := range config.Conf.Rule {
		if !s.Enable {
			continue
		}

		prefixAddr := EnsurePort(s.Source)
		replaceAddr := EnsurePort(s.Target)

		reqUrl := r.URL.String()

		if strings.HasPrefix(reqUrl, prefixAddr) {
			newUrl := strings.ReplaceAll(reqUrl, prefixAddr, replaceAddr)
			newURL, _ := url.Parse(newUrl)

			// 修改请求URL
			r.URL = newURL
			r.Host = newURL.Host

			log.Info("redirect: %s %s\n", r.Method, r.URL.String())
			return r, nil
		}
	}
	return r, nil
}

// EnsurePort 确保URL有端口，如果没有则添加默认端口
func EnsurePort(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}

	if u.Port() != "" {
		return rawURL
	}

	switch strings.ToLower(u.Scheme) {
	case "http":
		u.Host = fmt.Sprintf("%s:80", u.Hostname())
	case "https":
		u.Host = fmt.Sprintf("%s:443", u.Hostname())
	}
	return u.String()
}

func ResHandler(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	if resp != nil {
		//fmt.Printf("响应: %s %s\n", resp.Request.Method, resp.Request.URL.String())
	}
	return resp
}
