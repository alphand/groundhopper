package services

import (
  "auth/core/authentication"
  "auth/models"

  "errors"
  "encoding/json"
  "fmt"
  "net/http"

  jwt "github.com/dgrijalva/jwt-go"
)

func GetAccount(email string) (*models.User, error) {
  keyemail, _ := json.Marshal(email)
  url := fmt.Sprintf("http://192.168.99.100:5984/lvo-accounts/_design/email_finder/_view/email_finder?key=%s", keyemail)

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, err
  }

  client := &http.Client{}

  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()

  var data models.EmailFinder
  if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
    return nil, err
  }

  if len(data.Rows) == 0 {
    return nil, errors.New("No email found!")
  }

  return &models.User{
    UUID: data.Rows[0].Value.UUID,
    Email: data.Rows[0].Value.Email,
    Password: data.Rows[0].Value.Password,
  }, nil;
}


func Login(requestUser *models.User) (int, []byte){
  authBackend := authentication.InitJWTAuthBackend()

  if authBackend.Authenticate(requestUser) {
    token, err := authBackend.GenerateToken(requestUser.UUID)
    if err != nil {
      return http.StatusInternalServerError, []byte("")
    } else {
      response, _ := json.Marshal(models.TokenAuthentication{token})
      return http.StatusOK, response
    }
  }

  return http.StatusUnauthorized, []byte("")
}

func RefreshToken(requestUser *models.User) []byte {
  authBackend := authentication.InitJWTAuthBackend()
  token, err := authBackend.GenerateToken(requestUser.UUID)
  if err != nil {
    panic(err)
  }

  response, err := json.Marshal(models.TokenAuthentication{token})
  if err != nil {
    panic(err)
  }

  return response
}

func Logout(req *http.Request) error {
  authBackend := authentication.InitJWTAuthBackend()
  tokenRequest, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
    return authBackend.PublicKey, nil
  })

  if err != nil {
    return err
  }

  tokenString := req.Header.Get("Authorization")
  return authBackend.Logout(tokenString, tokenRequest)
}
