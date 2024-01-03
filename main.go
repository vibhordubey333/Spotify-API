package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	// Set up OAuth2 configuration with your client ID and secret
	config := &clientcredentials.Config{
		ClientID:     "",
		ClientSecret: "",
		TokenURL:     spotify.TokenURL,
	}

	// Get an access token using the client credentials flow
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Use the access token to create a Spotify client
	client := spotify.Authenticator{}.NewClient(token)

	// Example: Use the Spotify client to get a track
	track, err := client.GetTrack("3PzsbWSQdLCKDLxn7YZfkM")
	if err != nil {
		log.Fatal(err)
	}

	// Print the track details
	fmt.Println("Track Name:", track.Name)
	fmt.Println("Artist:", track.Artists[0].Name)
}

//package main
//
//import (
//	"fmt"
//	"net/http"
//
//	"github.com/zmb3/spotify"
//	_ "golang.org/x/oauth2"
//)
//
//func main() {
//	// Set up the OAuth2 configuration with your client ID and secret
//	auth := spotify.NewAuthenticator(
//		"http://localhost:8080/callback", // Redirect URL after authentication
//		spotify.ScopeUserReadPrivate,     // Scopes needed for access
//	)
//	auth.SetAuthInfo("your_client_id", "your_client_secret")
//
//	// Start a local server to handle the callback after authorization
//	http.HandleFunc("/callback", completeAuth)
//	go http.ListenAndServe(":8080", nil)
//
//	// Obtain the URL for the user to authorize the application
//	url := auth.AuthURL("state")
//	fmt.Println("Please log in to Spotify by visiting the following page:", url)
//
//	// Wait for the authorization to complete
//	select {}
//}
//
//func completeAuth(w http.ResponseWriter, r *http.Request) {
//	// Get the authorization code from the callback URL
//	token, err := spotify.NewAuthenticator("http://localhost:8080/callback").Token("state", r)
//	if err != nil {
//		http.Error(w, "Couldn't get token", http.StatusForbidden)
//		return
//	}
//
//	// Use the token to access the Spotify API
//	client := auth.NewClient(token)
//	user, err := client.CurrentUser()
//	if err != nil {
//		http.Error(w, "Couldn't get user", http.StatusForbidden)
//		return
//	}
//	fmt.Fprintf(w, "Logged in as: %s", user.ID)
//}
