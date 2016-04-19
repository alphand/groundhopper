package main

import (
  "os"
  "log"

  "github.com/codegangsta/negroni"
)

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    log.Fatal("PORT environment variable was not set!")
  }

  router := NewRouter()

  n := negroni.Classic()
  n.UseHandler(router)

  n.Run(":" + port)
}
