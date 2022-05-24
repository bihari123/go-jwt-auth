package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwt requires that you use a signing method
// by default, three signing methods are provided: 1. ECDSA 2. RSA 3.HMAC
// ECDSA and RSA use two keys for signing and checking (private key to sign but anyone with a public key can see if the signature is valid) a signature whereas HMAC uses a single key
// you also meed claims and golang provides a generic map claims
// or you can make your own claims using the standard claims as the base.

type UserClaims struct {
	jwt.StandardClaims // using standard claims as a base
	SessionID          int64
}



// Structured version of Claims Section, as referenced at
// https://tools.ietf.org/html/rfc7519#section-4.1
// See examples for how to use this with your own claim types
// type StandardClaims struct {
//	Audience  string `json:"aud,omitempty"`
//	ExpiresAt int64  `json:"exp,omitempty"`
//	Id        string `json:"jti,omitempty"`
//	IssuedAt  int64  `json:"iat,omitempty"`
//	Issuer    string `json:"iss,omitempty"`
//	NotBefore int64  `json:"nbf,omitempty"`
//	Subject   string `json:"sub,omitempty"`
// }



// Validating the claims
// the reason we are using a pointer to UserClaims is bcoz the jwt.StandardClaims uses a pointer 
func (u *UserClaims) Valid() error { // checks if the token is expired or not 
	// Compares the exp claim against cmp.
	// If required is false, this method will return true if the value matches or is unset
	//func (c *StandardClaims) VerifyExpiresAt(cmp int64, req bool) bool {
	//  return verifyExp(c.ExpiresAt, cmp, req)
	//}

	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Token has expired")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("invalid session id")
	}
	return nil 
}


func createToken(c * UserClaims)(token string, err error) {
  jwt.NewWithClaims()
}
