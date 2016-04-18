package main

import (
  "net/http"
  "net/http/httputil"
  "net/url"

  "fmt"
  "runtime"

  "github.com/gorilla/mux"
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

func (p *Prox) handle (w http.ResponseWriter, r *http.Request) {
  w.Header().Set("X-GoProxy", "GoProxy")
  p.proxy.ServeHTTP(w, r)
}

func ProxyHandler (w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  url := "http://" + vars["rest"]
  proxy := New(url)

  proxy.handle(w, r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello world, I'm running on %s with and %s CPU.",
    runtime.GOOS, runtime.GOARCH)
}

func AccountsIndex(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "testing library")
}
