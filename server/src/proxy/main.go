package proxy

import (
  "net/http"
  "net/http/httputil"
  "net/url"
)

type Prox struct {
  target        *url.URL
  proxy         *httputil.ReverseProxy
}

func NewHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
  director := func(req *http.Request) {
          req.URL.Scheme = target.Scheme
          req.URL.Host = target.Host
          req.URL.Path = target.Path
  }
  return &httputil.ReverseProxy{Director: director}
}

func New(target string) *Prox {
  locurl, _ := url.Parse(target)
  proxy := NewHostReverseProxy(locurl)
  return &Prox{target: locurl, proxy: proxy}
}

func (p *Prox) Handle (w http.ResponseWriter, r *http.Request) {
  w.Header().Set("X-GoProxy", "GoProxy")
  p.proxy.ServeHTTP(w, r)
}
