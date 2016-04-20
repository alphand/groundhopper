package routers

import(
  "controllers"
  // "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
  router.HandleFunc("/accounts/login", controllers.PostAccountsLogin).Methods("POST")
  return router;
}
