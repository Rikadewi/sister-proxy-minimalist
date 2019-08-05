package main

import (
	"net/http"
	"net/http/httputil"
)

type Proxy struct {
	HTTPProxy *httputil.ReverseProxy
}

func main() {
	proxy := &Proxy{
		HTTPProxy: NewForwardingHTTPProxy(),
	}

	server := &http.Server{
		Addr:    ":8000",
		Handler: proxy,
	}
	server.ListenAndServe()
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Scheme == "http" {
		p.handleHTTP(w, r)
	}
}

func (p *Proxy) handleHTTP(w http.ResponseWriter, r *http.Request) {
	p.HTTPProxy.ServeHTTP(w, r)
}

func NewForwardingHTTPProxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		if _, ok := req.Header["User-Agent"]; !ok {
			req.Header.Set("User-Agent", "")
		}
	}
	return &httputil.ReverseProxy{
		Director: director,
	}
}
