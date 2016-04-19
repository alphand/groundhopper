package main

import (
  "proxy"
  "auth"

  "encoding/json"

  "net/http"

  "fmt"
  "log"
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

func PostCreateSession(w http.ResponseWriter, r *http.Request) {

  type Creds struct {
    Email string `json:"email"`
    Password string `json:"password"`
  }

  decoder := json.NewDecoder(r.Body)
  var cred Creds
  err := decoder.Decode(&cred)

  if err != nil {
    panic("Cannot decode")
  }

  authData, err := auth.Authenticate(cred.Email, cred.Password)
  log.Println(authData)

  log.Printf("decode param: email => %s, pass => %s", cred.Email, cred.Password)
  fmt.Fprintf (w, "Post create session: %s -- %s", cred.Email, cred.Password)
}
