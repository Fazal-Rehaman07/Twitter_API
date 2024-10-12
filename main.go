package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dghubble/oauth1"
)

var tweet struct {
	Text string `json:"text"`
}

var (
	consumerKey       = os.Getenv("TWITTER_API_KEY")
	consumerSecret    = os.Getenv("TWITTER_API_KEY_SECRET")
	accessToken       = os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
)

// Twitter API client setup
func getClient() *http.Client {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	return httpClient
}

// Function to add a tweet
func addPost(tweet string) {
	url := "https://api.twitter.com/2/tweets"
	body := map[string]interface{}{
		"text": tweet,
	}
	postBody, _ := json.Marshal(body)

	client := getClient()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Error response from Twitter API: %s", resp.Status)
	}

	fmt.Println("Tweet successfully posted!")
}

// Function to delete a tweet
func deletePost(tweetID string) bool {
	url := fmt.Sprintf("https://api.twitter.com/2/tweets/%s", tweetID)

	client := getClient()
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error response from Twitter API: %s", resp.Status)
	}

	fmt.Println("Tweet successfully deleted!")
	return true
}

// HTTP handler to add a tweet
func postTweetHandler(w http.ResponseWriter, r *http.Request) {

	if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	addPost(tweet.Text)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweet)
	// fmt.Fprintf(w, "Tweet ID: %s", tweet.Tweet_id)
	// fmt.Fprintf(w, "Tweet posted: %s", tweet.Text)

}

// HTTP handler to delete a tweet
func deleteTweetHandler(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/delete/")

	if deletePost(idStr) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Tweet deleted: %s", idStr)
	} else {
		http.Error(w, "Tweet not found", http.StatusNotFound)
	}

}

func main() {
	// Define routes to test in Postman
	http.HandleFunc("/tweet", postTweetHandler)
	http.HandleFunc("/delete/", deleteTweetHandler)

	// Start server
	fmt.Println("Server is running on port 8080!")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
