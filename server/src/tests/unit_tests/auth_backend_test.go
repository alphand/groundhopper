package unit_tests

import (
  // "auth/core/authentication"
  // "auth/core/redis"
  "auth/models"
  "auth/settings"

  "os"
  // "github.com/stretchr/testify/assert"
	. "gopkg.in/check.v1"
  "testing"
)

func Test(t *testing.T) {
  TestingT(t)
}

type AuthBackendTestSuite struct {}

var _ = Suite(&AuthBackendTestSuite{})
// var t *testing.t

func (s *AuthBackendTestSuite) SetUpSuite(c *C)  {
  os.Setenv("GO_ENV", "test")
  settings.Init()
  models.Init()
}

func (suite *AuthBackendTestSuite) TestModelConnection(c *C) {
  var dbConn = "http://192.168.99.100:5984/"

  userMdl := &models.User{
    Email: "email",
    Password: "pass",
  }
  c.Assert(userMdl.DBConn(), Equals, dbConn)
}
