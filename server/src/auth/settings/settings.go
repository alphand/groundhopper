package settings

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
)

var envs = map[string]string {
  "dev": "settings/dev.json",
  "prod": "settings/prod.json",
}

type Settings struct {
  PrivateKeyPath        string
  PublicKeyPath         string
  JWTExpirationDelta    int
}

var settings Settings = Settings{}
var env = "dev"

func Init() {
  env = os.Getenv("GO_ENV")
  if env == "" {
    fmt.Println("Warning: Settings DEV env is used due to lack of GO_ENV")
    env = "dev"
  }

  LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
  content, err := ioutil.ReadFile(envs[env])
  if err != nil {
    fmt.Println("Error while reading file", err)
  }

  settings = Settings{}
  jsonErr := json.Unmarshal(content, &settings)
  if jsonErr != nil {
    fmt.Println("Error while parsing config file", jsonErr)
  }
}

func GetEnvironment() string {
  return env
}

func Get() Settings {
  if &settings == nil {
    Init()
  }
  return settings
}
