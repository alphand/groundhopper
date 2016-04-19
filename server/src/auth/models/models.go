package models

type Auth struct {
  Authenticated     bool        `json:"authenticated"`
  TokenId           string      `json:"tokenId"`
  Expires           int64  `json:"expires"`
}

type EmailFinder struct {
  TotalRows int `json:"total_rows"`
  Offset int `json:"offset"`
  Rows []struct {
    ID string `json:"id"`
    Key string `json:"key"`
    Value struct {
      ID string `json:"_id"`
      Rev string `json:"_rev"`
      Email string `json:"email"`
      Password string `json:"password"`
    } `json:"value"`
  } `json:"rows"`
}

type EmailCred struct {
  Email string
  Password string
}

type PostCreds struct {
  Email string `json:"email"`
  Password string `json:"password"`
}
