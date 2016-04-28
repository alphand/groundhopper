package models

import(
  "log"
  "utils/netter"
)

type User struct {
  ID string         `json:"_id"`
  Rev string        `json:"_rev"`
  UUID string       `json:"uuid"`
  Email string      `json:"email"`
  Password string   `json:"password"`
}

func (user *User) DBConn() string {
  return Get().ConnStr + "/lvo-accounts"
}

func (user *User) Save(newUser *User) User {
  api := netter.Api(user.DBConn())
  api.Get()


  log.Println("from usermodel", Get().ConnStr, api.Api.BaseUrl, api.Raw)
  return User{}
}
