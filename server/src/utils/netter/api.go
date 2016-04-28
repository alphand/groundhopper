package netter

import (
  "net/http"
  "net/url"
)

type APIStruct struct {
  BaseUrl     *url.URL
  Client      *http.Client
}

func Api(baseUrl string) *Resource {
  u, err := url.Parse(baseUrl)
  if err != nil {
    panic("Api() - url.Parse(baseUrl) Error: "+ err.Error())
  }

  apiInstance := &APIStruct{BaseUrl: u}

  client := &http.Client{}

  apiInstance.Client = client

  return &Resource{Url:"", Api: apiInstance}
}
