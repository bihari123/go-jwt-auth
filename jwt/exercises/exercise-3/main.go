package main

import (
	"crypto/hmac"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/icza/gog"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	http.HandleFunc("/", showHome)
	http.HandleFunc("/register", registerUser)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}

var cred = map[string][]byte{}

func showHome(w http.ResponseWriter, r *http.Request) {
	isEqual := gog.If(len(cred) > 0, true, false)

	userName := "UserNotRegisterd"
	hashedPass := []byte("UserNotRegisterd")
	if isEqual {
		// userName = user.Value
		// hashedPass = password.Value
		for k, v := range cred {
			userName, hashedPass = k, v
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
	<h1>LOG IN</h1>
        <form action="/login" method="POST">
            <input type="text" name="user">
			<input type="password" name="password">
			<input type="submit">
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

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errMsg := url.QueryEscape("someone is trying to hack")
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
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if _, ok := cred[userName]; ok {

		// CompareHashAndPassword compares a bcrypt hashed password with its possible
		// plaintext equivalent. Returns nil on success, or an error on failure.
		//func CompareHashAndPassword(hashedPassword, password []byte) error {

		err := bcrypt.CompareHashAndPassword(cred[userName], []byte(pass))
		if _, isErrorExist := CheckError("compareHashAndPass", err); isErrorExist {
			fmt.Println("Password and username don't match")
			return
		}
	}
	msg := url.QueryEscape(userName + " is LoggedIn")
	http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)

}

func CheckError(funcName string, err error) (interface{}, bool) {
	if err != nil {
		log.Println(fmt.Errorf("Error in %s: %w", funcName, err))
		return nil, true
	}
	return nil, false
}

func createToken(sid string )(signed string){
  key:=[]byte("this is a key")
  mac:= hmac.New(sha256,key)
  mac.Write([]byte(sid))
  return
}

