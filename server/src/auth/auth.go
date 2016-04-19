package auth

import "time"

type Auth struct {
  Authenticated     bool        `json:"authenticated"`
  TokenId           string      `json:"tokenId"`
  Expires           time.Time   `json:"expires"`
}
