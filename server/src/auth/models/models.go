package models

import (
  "auth/settings"
)

type Auth struct {
  Authenticated     bool        `json:"authenticated"`
  TokenId           string      `json:"tokenId"`
  Expires           int64       `json:"expires"`
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
  Get()  (interface{}, error)
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
