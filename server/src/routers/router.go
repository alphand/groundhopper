package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  router = SetAuthenticationRoutes(router)
  router = SetHelloRoutes(router)
  return router;
}
