package auth

import (
  "encoding/json"
  "net/http"

  "errors"

  "log"
  "fmt"
  "time"

  // "github.com/dgrijalva/jwt-go"
  // "github.com/gorilla/context"
  // "github.com/auth0/go-jwt-middleware"k
)

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

func getAccount(email string) (*EmailCred, error) {
  keyemail, _ := json.Marshal(email)
  url := fmt.Sprintf("http://192.168.99.100:5984/lvo-accounts/_design/email_finder/_view/email_finder?key=%s", keyemail)

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal("Failed NewRequest: ", err)
    return nil, err
  }

  client := &http.Client{}

  resp, err := client.Do(req)
  if err != nil {
    log.Fatal("Failed DO:", err)
    return nil, err
  }

  defer resp.Body.Close()

  var data EmailFinder
  if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
    log.Println(err)
    return nil, err
  }

  log.Println("EF Data", data)

  if len(data.Rows) == 0 {
    return nil, errors.New("No email found!")
  }

  return &EmailCred{
    Email: data.Rows[0].Value.Email,
    Password: data.Rows[0].Value.Password,
  }, nil;
}

func authenticate(email string, password string) (*Auth, error) {
  acc, err := getAccount(email)
  if err != nil {
    return nil, err
  }

  return &Auth{
    Authenticated:false,
    TokenId: acc.Email,
    Expires:time.Now().Add(time.Minute * 3).Unix(),
  }, nil
}

func PostCreateSession(w http.ResponseWriter, r *http.Request) {

  type Creds struct {
    Email string `json:"email"`
    Password string `json:"password"`
  }

  decoder := json.NewDecoder(r.Body)
  var cred Creds
  err := decoder.Decode(&cred)

  if err != nil {
    panic("Cannot decode")
  }

  authData, err := authenticate(cred.Email, cred.Password)

  if err != nil {
    w.Header().Set("Content-Type","application/json;charset=UTF-8")
    w.WriteHeader(422)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  log.Println(authData)

  w.Header().Set("Content-Type","application/json;charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  if err := json.NewEncoder(w).Encode(authData); err != nil {
    panic(err)
  }

  // log.Printf("decode param: email => %s, pass => %s", cred.Email, cred.Password)
  // fmt.Fprintf (w, "Post create session: %s -- %s", cred.Email, cred.Password)
}
