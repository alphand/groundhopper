package models

import(
  "log"
)

type User struct {
  ID string         `json:"_id"`
  Rev string        `json:"_rev"`
  UUID string       `json:"uuid"`
  Email string      `json:"email"`
  Password string   `json:"password"`
}

func (user *User) DBConn() string {
  return Get().ConnStr
}

func (user *User) Save(newUser User) User {
  log.Println("from usermodel", Get().ConnStr)
  return User{}
}
