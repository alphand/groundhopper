package authentication

import (
  "bufio"

  "crypto/rsa"
  "crypto/x509"

  "encoding/pem"

  "os"
  "time"

  "auth/core/redis"
  "auth/settings"
  "auth/models"

  "golang.org/x/crypto/bcrypt"
  jwt "github.com/dgrijalva/jwt-go"

)

type JWTAuthBackend struct {
  privateKey    *rsa.PrivateKey
  PublicKey     *rsa.PublicKey
}

const (
  TOKEN_DURATION = 72
  EXPIRE_OFFSET = 3600
)

var authBackendInstance *JWTAuthBackend = nil
func InitJWTAuthBackend() *JWTAuthBackend {
  if authBackendInstance == nil {
    authBackendInstance = &JWTAuthBackend{
      privateKey: getPrivateKey(),
      PublicKey:  getPublicKey(),
    }
  }
  return authBackendInstance
}

func (backend *JWTAuthBackend) GenerateToken(userUUID string) (string, error) {
  token := jwt.New(jwt.SigningMethodRS512)
  token.Claims["exp"] = time.Now().Add(time.Hour * time.Duration(settings.Get().JWTExpirationDelta)).Unix()
  token.Claims["iat"] = time.Now().Unix()
  token.Claims["sub"] = userUUID
  tokenString, err := token.SignedString(backend.privateKey)
  if err != nil {
    // panic(err)
    return "", err
  }

  return tokenString, nil
}

func (backend *JWTAuthBackend) Authenticate(user *models.User) bool {
  hashedPass, _ := bcrypt.GenerateFromPassword([]byte("testing"), 10)

  testUser := models.User{
    UUID: "asdfasdf",
    Email: "darmawan.niko@gmail.com",
    Password: string(hashedPass),
  }

  return user.Email == testUser.Email &&
    bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil
}

func (backend *JWTAuthBackend) Logout(tokenString string, token *jwt.Token) error {
  redisConn := redis.Connect()
  return redisConn.SetValue(tokenString, tokenString,
    backend.getTokenRemainingValidity(token.Claims["exp"]))
}

func (backend *JWTAuthBackend) IsInBlacklist(token string) bool {
  redisConn := redis.Connect()
  redisToken, _ := redisConn.GetValue(token)

  if redisToken == nil {
    return false
  }

  return true
}


// ========= Private Method ============
func readPEM(keyFilePath string) *pem.Block {
  keyFile, err := os.Open(keyFilePath)
  if err != nil {
    panic(err)
  }

  pemfileinfo, _ := keyFile.Stat()
  var size int64 = pemfileinfo.Size()
  pembytes := make([]byte, size)

  buffer := bufio.NewReader(keyFile)
  _, err = buffer.Read([]byte(pembytes))

  keyFile.Close()

  data, _ := pem.Decode([]byte(pembytes))

  return data
}

func getPrivateKey() *rsa.PrivateKey {

  data := readPEM(settings.Get().PrivateKeyPath)
  privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

  if err != nil {
    panic(err)
  }

  return privateKeyImported
}

func getPublicKey() *rsa.PublicKey {
  data := readPEM(settings.Get().PublicKeyPath)
  publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

  if err != nil {
    panic(err)
  }

  rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

  if !ok {
    panic(err)
  }

  return rsaPub
}

func (backend *JWTAuthBackend) getTokenRemainingValidity(timestamp interface{}) int {
  if validity, ok := timestamp.(float64); ok {
    tm := time.Unix(int64(validity), 0)
    remainer := tm.Sub(time.Now())
    if remainer > 0 {
      return int(remainer.Seconds() + EXPIRE_OFFSET)
    }
  }

  return EXPIRE_OFFSET
}
