package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const clientID = "d60d36ba61bdca035b96"
const clientSecret = "e6a12bb211a9342521b2e6f9eba72a0dc3bde35f"

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// we will be using a httpClient to make external HTTP requests later in code
	httpClient := http.Client{}

	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {

		// First, we need to get the value of the code query param
		log.Println("Parsing the form")
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(os.Stdout, "couldn't parse query: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		code := r.FormValue("code")
		log.Println("code: ", code)
		// Next, let's forward the HTTP request to call the git
		// to get out access token

		reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)

		req, err := http.NewRequest(http.MethodPost, reqURL, nil)

		if err != nil {
			fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		// We set this header since we want the ResponseWriter
		// as JSON

		req.Header.Set("accept", "application/json")

		//Send out the HTTP Request
		log.Println("Sending the request back to the github")
		res, err := httpClient.Do(req)

		if err != nil {
			fmt.Fprintf(os.Stdout, "couldn't sent request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		defer res.Body.Close()

		// Parse the request body into thew "OAthAccessResponse" struct
		var t OAuthAccessResponse

		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
			fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		// Finally, send a response to redirect the user to the "welcome" page
		// with the access token
		w.Header().Set("Location", "/welcome.html?access_token="+t.AccessToken)
		w.WriteHeader(http.StatusFound)
	})

	log.Println("starting the server at 8080")
	http.ListenAndServe(":8080", nil)
}

/*
The welcome page is the page we show the user after they have logged in. Now that we have the users access token, we can obtain their account information on their behalf as authorized Github users.
We will be using the /user API to get basic info about the user and say hi to them on our welcome page.
*/
