package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"github.com/icza/gog"
)


func main() {
	http.HandleFunc("/", showHome)
	http.HandleFunc("/register", registerUser)
	http.ListenAndServe(":8080", nil)
}

var cred = map[string][]byte{}

func showHome(w http.ResponseWriter, r *http.Request) {


	isEqual:=	gog.If(len(cred)>0,true,false)	
	
	userName := "UserNotRegisterd"
	hashedPass := []byte("UserNotRegisterd")
	if isEqual {
		// userName = user.Value
		// hashedPass = password.Value
		for k,v:=range cred{
			userName,hashedPass=k,v 
		}
		
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
	<p>Cookie Value: <br> UserName: ` + userName + `<br> PasswordHash: ` + string(hashedPass) + `</p>
		<form action="/register" method="post">
		  <label>username:	<input type="text" name="user" /> </label><br>
			<label>password: <input type = "text" name="password"/> </label><br>
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
		http.Error(w, fmt.Sprintf(err.Error()), http.StatusInternalServerError)
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
