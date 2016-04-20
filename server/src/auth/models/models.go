package models

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

type User struct {
  ID string         `json:"_id"`
  Rev string        `json:"_rev"`
  UUID string       `json:"uuid"`
  Email string      `json:"email"`
  Password string   `json:"password"`
}

type PostCreds struct {
  Email string `json:"email"`
  Password string `json:"password"`
}

type TokenAuthentication struct {
  Token string `json:"token" form:"token"` 
}
