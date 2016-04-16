package main

import (
  "os"
  "fmt"
  "log"
  "net/http"
  "runtime"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello world, I'm running on %s with and %s CPU.",
    runtime.GOOS, runtime.GOARCH)
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    log.Fatal("PORT environment variable was not set!")
  }

  server := http.Server{
      Addr: ":"+port,
      Handler: &myHandler{},
  }

  mux = make(map[string]func(http.ResponseWriter, *http.Request))
  mux["/"] = indexHandler

  err := server.ListenAndServe()
  if err != nil {
    log.Fatal("Could not listen: ", err)
  } else {
    log.Println("Ready to Serve on :%s", port)
  }
}

type myHandler struct {}

func (*myHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
  if h, ok := mux[r.URL.String()]; ok {
    h(w, r)
    return
  }

  fmt.Fprint(w, "My server url: "+r.URL.String())
}
