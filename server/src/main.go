package main

import (
  "os"
  "log"
  "net/http"
)

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    log.Fatal("PORT environment variable was not set!")
  }

  router := NewRouter()
  server := http.Server{
      Addr: ":"+port,
      Handler: router,
  }

  err := server.ListenAndServe()
  if err != nil {
    log.Fatal("Could not listen: ", err)
  } else {
    log.Println("Ready to Serve on :%s", port)
  }
}
