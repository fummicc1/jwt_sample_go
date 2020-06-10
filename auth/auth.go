package auth

import (
	"crypto/rsa"

	"github.com/gin-gonic/gin"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// Login is Post Method to handle login.
func Login(c *gin.Context) {
	// var user *model.User
	// c.BindJSON(user)
}
