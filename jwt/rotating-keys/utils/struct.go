package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwt requires that you use a signing method
// by default, three signing methods are provided: 1. ECDSA 2. RSA 3.HMAC
// ECDSA and RSA use two keys for signing and checking (private key to sign but anyone with a public key can see if the signature is valid) a signature whereas HMAC uses a single key
// you also meed claims and golang provides a generic map claims
// or you can make your own claims using the standard claims as the base.

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

type UserClaims struct {
	jwt.StandardClaims // using standard claims as a base
	SessionID          int64
}


type key struct {
	key     []byte
	created time.Time // we define the created time bcoz then we can decide to delete the key that was created over a week ago
}


