package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)


var currentKID = ""
var keys = map[string]key{}


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

// have a crone job that constantly creates a new key after a fixed period of time
func generateKey() (string, error) {

	newKey := make([]byte, 64)
	// rand.Reader in the crypto module which is the most random that your computer can do
	_, err := io.ReadFull(rand.Reader, newKey)
	if err != nil {
		return "", fmt.Errorf("Error in generating key: %w", err)
	}

	// generating uuid with hyphen
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	currentKID = uuid 
	keys[currentKID] = key{
		key:     newKey,
		created: time.Now(),
	}

	fmt.Printf("\n\nthe key is generated : %v\nnewKey is %v\n", currentKID, newKey)
	return currentKID, nil
}

func CreateToken(c *UserClaims) (signedToken string, err error) {
	currentKID, err = generateKey()
// setting expiry of the token 
	c.StandardClaims.ExpiresAt=  time.Now().Add(time.Minute*15).Unix()

	// alternate way of doing it
	// claim:=jwt.MapClaims{}
	// claim["exp"]= time.Now().Add(time.Minute*15).Unix()
	// claim["session_id"]=&c.SessionID


			// this will create a base token
	// this requires a signing method and claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	t.Header["kid"]=currentKID 
	fmt.Printf("\n\nThe current key : %+v\n\n", keys[currentKID])
	signedToken, err = t.SignedString(keys[currentKID].key)
	if err != nil {
		fmt.Print(fmt.Errorf("\n\nSome error in create token: %w\n\n", err))
		return "", err
	}
	return signedToken, nil
}

func ParseToken(signedToken string) (claims *UserClaims, err error) {

	fmt.Println("in parse with claims")
	// it parses the token without verifying and then passes the token into the function defined inside the function to verify it. If
	// the token is verified, then it returns the token
	t, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		// first we are checking whether thhe signing method of the token passed is equal to the expected one
		// it is a good practice to verifu the algo

		fmt.Println("checking the signing method")

		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("Invalid signing algorithm")
		}

		fmt.Println("checking the header", t.Header)

		kid, ok := t.Header["kid"].(string)

		if !ok {
			return nil, fmt.Errorf("Not Found key ID")
		}

		fmt.Println("checking the keys[kid]")

		k, ok := keys[kid]

		if !ok {
			// we can make this error more specific, but for your login code, you want it to be as confusing as possible to keep the hackers away
			return nil, fmt.Errorf("Invalid key id")
		}

		fmt.Println("returning the key")

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

	fmt.Printf("\n\nissued at : %v",cast.ToTime(claims.IssuedAt))

	return claims, nil
}
