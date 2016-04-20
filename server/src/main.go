package main

import (
  "auth/settings"
  "os"
  "log"

  "github.com/codegangsta/negroni"
  "routers"
)

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    log.Fatal("PORT environment variable was not set!")
  }

  settings.Init()

  router := routers.InitRoutes()

  n := negroni.Classic()
  n.UseHandler(router)

  n.Run(":" + port)
}
