package auth

import (
  "encoding/json"
  "net/http"

  "errors"

  "log"
  "fmt"
  "time"

  "auth/models"

  // "github.com/dgrijalva/jwt-go"
  // "github.com/gorilla/context"
  // "github.com/auth0/go-jwt-middleware"k
)


func getAccount(email string) (*models.EmailCred, error) {
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

  var data models.EmailFinder
  if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
    log.Println(err)
    return nil, err
  }

  log.Println("EF Data", data)

  if len(data.Rows) == 0 {
    return nil, errors.New("No email found!")
  }

  return &models.EmailCred{
    Email: data.Rows[0].Value.Email,
    Password: data.Rows[0].Value.Password,
  }, nil;
}

func authenticate(email string, password string) (*models.Auth, error) {
  acc, err := getAccount(email)
  if err != nil {
    return nil, err
  }

  if (acc != nil && acc.Password != password) {
    return &models.Auth {
      Authenticated: false,
    }, nil
  }

  return &models.Auth{
    Authenticated:false,
    TokenId: acc.Email,
    Expires:time.Now().Add(time.Minute * 3).Unix(),
  }, nil
}

func PostCreateSession(w http.ResponseWriter, r *http.Request) {

  decoder := json.NewDecoder(r.Body)
  var cred models.PostCreds
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
}
