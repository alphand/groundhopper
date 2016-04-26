package models

import (
  "auth/settings"
)

type Auth struct {
  Authenticated     bool        `json:"authenticated"`
  TokenId           string      `json:"tokenId"`
  Expires           int64       `json:"expires"`
}

type EmailFinder struct {
  TotalRows int `json:"total_rows"`
  Offset int `json:"offset"`
  Rows []struct {
    ID string `json:"id"`
    Key string `json:"key"`
    Value struct {
      ID string         `json:"_id"`
      Rev string        `json:"_rev"`
      UUID string       `json:"uuid"`
      Email string      `json:"email"`
      Password string   `json:"password"`
    } `json:"value"`
  } `json:"rows"`
}

type PostCreds struct {
  Email string `json:"email"`
  Password string `json:"password"`
}

type TokenAuthentication struct {
  Token string `json:"token" form:"token"`
}

type CouchDB struct {
  ConnStr string
}

type Model interface {
  Save() (interface{}, error)
}

var couchDB CouchDB = CouchDB{}

func Init() {
  couchDB = CouchDB{
    ConnStr: settings.Get().DBConn,
  }
}

func Get() CouchDB {
  if &couchDB == nil {
    Init()
  }

  return couchDB
}
