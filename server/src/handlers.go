package main

import (
  "net/http"
  "fmt"
  "runtime"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello world, I'm running on %s with and %s CPU.",
    runtime.GOOS, runtime.GOARCH)
}

func AccountsIndex(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "testing library")
}
