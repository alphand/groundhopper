package main

import (
  "proxy"

  "net/http"

  "fmt"
  "runtime"

  "github.com/gorilla/mux"

  // "github.com/dgrijalva/jwt-go"
  // "github.com/gorilla/context"
  // "github.com/auth0/go-jwt-middleware"

)

func ProxyHandler (w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  url := "http://" + vars["rest"]
  proxyEn := proxy.New(url)

  proxyEn.Handle(w, r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello world, I'm running on %s with and %s CPU.",
    runtime.GOOS, runtime.GOARCH)
}

func AccountsIndex(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "testing library")
}
