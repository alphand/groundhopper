package routers

import (
  "controllers"
  "net/http"

  "log"
  "time"

  "github.com/gorilla/mux"
)

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
    controllers.IndexHandler,
  },
  Route{
    "AccountsIndex",
    "GET",
    "/api/accounts",
    controllers.AccountsIndex,
  },
  Route{
    "GetProxyIndex",
    "GET",
    `/proxy/{rest:[a-zA-Z0-9=\-\/_:\.]+}`,
    controllers.ProxyHandler,
  },
}

func LoggerMiddleware(inner http.Handler, name string) http.Handler {
  return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
    start := time.Now()
    inner.ServeHTTP(w, r)
    log.Printf(
      "%s\t%s\t%s\t%s",
      r.Method,
      r.RequestURI,
      name,
      time.Since(start),
    )
  })
}

func SetHelloRoutes (router *mux.Router) *mux.Router {

  for _, route := range routes {
    var handler http.Handler
    handler = route.HandlerFunc
    handler = LoggerMiddleware(handler, route.Name)

    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(handler)
  }

  return router
}
