package controllers

import (
  "auth/models"
  "auth/services"

  "encoding/json"

  "net/http"
)

func PostAccountsLogin(w http.ResponseWriter, r *http.Request) {

  cred := new(models.User)
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&cred)

  if err != nil {
    panic("Cannot decode")
  }

  credUser, err := services.GetAccount(cred.Email)

  if err != nil {
    w.Header().Set("Content-Type","application/json;charset=UTF-8")
    w.WriteHeader(422)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  responseStatus, token := services.Login(credUser)
  w.Header().Set("Content-Type","application/json;charset=UTF-8")
  w.WriteHeader(responseStatus)
  if err := json.NewEncoder(w).Encode(token); err != nil {
    panic(err)
  }
}
