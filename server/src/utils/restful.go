package utils

import (
  "net/http"
  "url"
)

type Client struct {
  Client      *http.Client
  BaseUrl     *url.URL
  UserAgent   string


  onRequestCompleted RequestCompletionCallback
}

type RequestCompletionCallback func(*http.Request, *http.Response)
