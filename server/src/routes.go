package main

import "net/http"

type Route struct {
  Name        string
  Method      string
  Pattern     string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
  Route{
    "Index",
    "GET",
    "/",
    indexHandler,
  },
  Route{
    "AccountsIndex",
    "GET",
    "/api/accounts",
    AccountsIndex,
  },
  Route{
    "GetProxyIndex",
    "GET",
    `/proxy/{rest:[a-zA-Z0-9=\-\/_:\.]+}`,
    ProxyHandler,
  },
}
