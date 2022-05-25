package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type myClaims struct {
	jwt.StandardClaims
	Password string
}

const myKey = "i love thursdays when it rains 8723 inches"

func main() {
	http.HandleFunc("/", showHome)
	http.HandleFunc("/register", registerUser)
	http.ListenAndServe(":8080", nil)
}

var cred = map[string]string{}

func showHome(w http.ResponseWriter, r *http.Request) {

	password, err := r.Cookie("password")

	if _, errorExists := CheckError("fetch cokkie Password", err); errorExists {
		password = &http.Cookie{}
	}

	user, err := r.Cookie("user")
	if _, errorExists := CheckError("Fetching User Cokkie", err); errorExists {
		user = &http.Cookie{}
	}

	signedPass, err := jwt.ParseWithClaims(password.Value, &myClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("somebody is trying to hack")
		}
		return []byte(myKey), nil
	})

	isEqual := err == nil && signedPass.Valid

	userName := "UserNotRegisterd"
	hashedPass := "UserNotRegisterd"
	if isEqual {
		userName = user.Value
		hashedPass = password.Value
	}

	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>HMAC Example</title>
	</head>
	<body>
	<p>Cookie Value: <br> UserName: ` + userName + `PasswordHash: ` + hashedPass + `</p>
		<form action="/submit" method="post">
		  <label>username:	<input type="user" name="user" /> </label><br>
			<label>password: <input type = "password" name="password"/> </label><br>
			<input type="submit" />
		</form>
	</body>
	</html>`
	io.WriteString(w, html)

}

func registerUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errMsg := url.QueryEscape("your method was not post")
		fmt.Println(errMsg)
		return
	}

	userName := r.FormValue("user")

	if userName == "" {
		fmt.Println(fmt.Errorf("UserName is empty"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	pass := r.FormValue("password")
	if pass == "" {
		fmt.Println(fmt.Errorf("Password is empty"))
	}

	encrypted_pass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if _, isErrorExist := CheckError("encrypting pass", err); isErrorExist {
		http.Error(w, fmt.Sprintf(err), http.StatusInternalServerError)
	}

	cred[userName] = encrypted_pass
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func CheckError(funcName string, err error) (interface{}, bool) {
	if err != nil {
		log.Println(fmt.Errorf("Error in %s: %w", funcName, err))
		return nil, true
	}
	return nil, false
}
