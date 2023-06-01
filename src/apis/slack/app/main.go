package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"log"
	"net/http"
	"os"
)

var (
	slackClientID     = os.Getenv("SLACK_CLIENT_ID")
	slackClientSecret = os.Getenv("SLACK_CLIENT_SECRET")

	listenAddress = os.Getenv("LISTEN_ADDRESS")
	certFile      = os.Getenv("CERT_FILE")
	certKeyFile   = os.Getenv("CERT_KEY_FILE")
)

func init() {
	if slackClientID == "" {
		panic("SLACK_CLIENT_ID is not set")
	}

	if slackClientSecret == "" {
		panic("SLACK_CLIENT_SECRET is not set")
	}

	if listenAddress == "" {
		listenAddress = ":8433"
	}

	if certFile == "" {
		panic("CERT_FILE is not set")
	}

	if certKeyFile == "" {
		panic("CERT_KEY_FILE is not set")
	}
}

func Handler(rw http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	oAuthV2Response, err := slack.GetOAuthV2Response(http.DefaultClient, slackClientID, slackClientSecret, code, "")
	if err != nil {
		log.Printf("failed to get oauth response: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
	log.Printf("Access token: %s", oAuthV2Response.AccessToken)
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(
		http.ListenAndServeTLS(
			fmt.Sprintf(":%s", listenAddress),
			certFile,
			certKeyFile,
			nil,
		),
	)
}
