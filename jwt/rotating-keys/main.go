package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pborman/uuid"
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


type key struct{
	key []byte 
	created time.Time // we define the created time bcoz then we can decide to delete the key that was created over a week ago  
}
var currentKID=""
var keys = map[string]key{}

func generateKey()error{
	
	newKey:= make([]byte,64)
	// rand.Reader in the crypto module which is the most random that your computer can do
	_,err:=io.ReadFull(rand.Reader,newKey)
	if err!=nil{
		return fmt.Errorf("Error in generating key: %w",err)
	}
  
  uid,err:=uuid.NewV4()

	return  nil 
}

func createToken(c *UserClaims) (token string, err error) {
	// this will create a base token
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c) // this requires a signing method and claims
	signedToken, err := t.SignedString(keys[currentKID])

	return signedToken, nil
}

func parseToken(signedToken string) (claims *UserClaims, err error) {

	// it parses the token without verifying and then passes the token into the function defined inside the function to verify it. If
	// the token is verified, then it returns the token
	t, err := jwt.ParseWithClaims(signedToken, claims, func(t *jwt.Token) (interface{}, error) {
		// first we are checking whether thhe signing method of the token passed is equal to the expected one
		// it is a good practice to verifu the algo
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("Invalid signing algorithm")
		}

		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("Invalid key ID")
		}

		k,ok:=keys[kid]
		if !ok{
			// we can make this error more specific, but for your login code, you want it to be as confusing as possible to keep the hackers away
			return nil,fmt.Errorf("Invalid key id")
		}


		return k.key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Error in parse token: %w", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("token not valid: %w", err)
	}

	// t.Claims is an interface so we have to assert the type to *UserClaims
	claims = t.Claims.(*UserClaims)
	return claims, nil
}
