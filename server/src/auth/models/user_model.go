package models

import(
  "errors"
  "encoding/json"
  "log"
  "utils/netter"
)

type EmailFinder struct {
  TotalRows int `json:"total_rows"`
  Offset int `json:"offset"`
  Rows []struct {
    ID string `json:"id"`
    Key string `json:"key"`
    Value User `json:"value"`
  } `json:"rows"`
}

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

func GetUserByEmail(email string) (*User, error) {
  api := netter.Api(Get().ConnStr + "/lvo-accounts")
  resp := &EmailFinder{}
  query := make(map[string]string)
  encEmail, _ := json.Marshal(email)
  query["key"] = string(encEmail)
  api.Res("_design/email_finder/_view/email_finder", resp).Get(query)

  if resp == nil || len(resp.Rows) <= 0 {
    return nil, errors.New("User not found")
  }

  return &resp.Rows[0].Value, nil
}

func (user *User) Save(newUser *User) User {
  api := netter.Api(user.DBConn())
  type dbRes struct {
    DBName string `json:"db_name"`
    DocCount int `json:"doc_count"`
  }
  resp := &dbRes{}
  api.Res("", resp).Get()


  log.Println("from usermodel", Get().ConnStr, api.Api.BaseUrl, resp.DocCount)
  return User{}
}
