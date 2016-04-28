package netter

import (
  "encoding/json"
  "errors"
  "io"
  "net/http"
  "net/url"
)

type Resource struct {
  Api           *APIStruct
  Url           string
  QueryValues   url.Values
  Payload       io.Reader
  Headers       http.Header
  Response      interface{}
  Raw           *http.Response
}

func (r *Resource) Res(options ...interface{}) *Resource {
  if len(options) > 0 {
    var url string
    if len(r.Url) > 0 {
      url = r.Url + "/" + options[0].(string)
    } else {
      url = options[0].(string)
    }

    headers := r.Headers
    if headers == nil {
      headers = http.Header{}
    }

    newR := &Resource{Url:url, Api:r.Api, Headers:headers}

    if len(options) > 1 {
      newR.Response = options[1]
    }

    return newR
  }
  return r
}

func (r *Resource) setQuery(querystring map[string]string) *Resource {
  r.QueryValues = make(url.Values)
  for k,v := range querystring {
    r.QueryValues.Set(k, v)
  }

  return r
}

func (r *Resource) Get(options ...interface{}) (*Resource, error) {
  if len(options) > 0 {
    if qry, ok := options[0].(map[string]string); ok {
      r.setQuery(qry)
    } else {
      return nil, errors.New("can't use options[0] as query")
    }
  }

  return r.do("GET")
}

func (r *Resource) do(method string) (*Resource, error) {
  url := *r.Api.BaseUrl
  if len(url.Path) > 0 {
    url.Path += "/" + r.Url
  } else {
    url.Path = r.Url
  }

  url.RawQuery = r.QueryValues.Encode()
  req, err := http.NewRequest(method, url.String(), r.Payload)
  if err != nil {
    return r, err
  }

  if r.Headers != nil {
    for k, _ := range r.Headers {
      req.Header.Set(k, r.Headers.Get(k))
    }
  }

  resp, err := r.Api.Client.Do(req)
  if err != nil {
    return r, err
  }

  r.Raw = resp

  if resp.StatusCode >= 400 {
    return r, nil
  }

  defer resp.Body.Close()

  err = json.NewDecoder(resp.Body).Decode(r.Response)
  if err != nil {
    return r, err
  }

  return r, nil
}
