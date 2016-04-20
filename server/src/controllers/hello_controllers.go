package controllers


import (
  "proxy"

  "net/http"

  "fmt"
  "runtime"

  "github.com/gorilla/mux"
)

func ProxyHandler (w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  url := "http://" + vars["rest"]
  proxyEn := proxy.New(url)

  proxyEn.Handle(w, r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello world, I'm running on %s with and %s CPU.",
    runtime.GOOS, runtime.GOARCH)
}

func AccountsIndex(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "testing library")
}
