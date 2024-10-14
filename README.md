# Twitter API in Go

## Introduction

This Go-based project demonstrates how to interact with the **Twitter API** using OAuth1. In Assignment we have done the following:
- Interacting with the Twitter API to **post** and **delete tweets**.
- Building a simple HTTP API in Go that can be tested using tools like **Postman**.

---

## Setup Instructions

### 1. Set Up a Twitter Developer Account

Before running this project, you need to have a **Twitter Developer Account** and access to Twitter's API keys. Follow these steps:

1. Go to the [Twitter Developer Portal](https://developer.twitter.com/).
2. Create a **Developer Account** if you don't have one already.
3. Create a new project by clicking on **Create New App** in the **Projects & Apps** section.
4. Fill in the required details (App Name, etc.), and once the app is created, you'll be provided with the **API Key**, **API Key Secret**, **Access Token**, and **Access Token Secret**.

### 2. Generate API Keys

After setting up your Twitter Developer Account:

1. Go to your app’s **Keys and Tokens** section.
2. Generate your **API Key** and **API Key Secret**.
3. Generate your **Access Token** and **Access Token Secret**.

### 3. Set Up Environment Variables

Set the following environment variables with the Twitter credentials you've obtained:

For **Linux/macOS**, use:
```bash
export TWITTER_API_KEY="your-api-key"
export TWITTER_API_KEY_SECRET="your-api-key-secret"
export TWITTER_ACCESS_TOKEN="your-access-token"
export TWITTER_ACCESS_TOKEN_SECRET="your-access-token-secret"
```
For **Windows (CMD)**, use:
```cmd
set TWITTER_API_KEY=your-api-key
set TWITTER_API_KEY_SECRET=your-api-key-secret
set TWITTER_ACCESS_TOKEN=your-access-token
set TWITTER_ACCESS_TOKEN_SECRET=your-access-token-secret
```

### 4. Run the Program

1. Clone this repository or download the project files.
2. Install the required Go packages:
   ```bash
   go get github.com/dghubble/oauth1
   ```
3. Run the Go program:
   ```bash
   go run main.go
   ```
4. The server will start on `localhost:8080`.

---

## Program Details

### 1. Posting a Tweet

To post a tweet, the `/tweet` endpoint is used. This endpoint accepts a `POST` request containing the tweet text in JSON format. The program will authenticate with Twitter and make a `POST` request to the Twitter API to post the tweet.

#### Example API Request:

- **Endpoint**: `POST http://localhost:8080/tweet`
- **Request Body**:
  ```json
  {
      "text": "I am excited to share that I have developed a Twitter API in Go!"
  }
  ```

#### Example Response:

- **Response Body**:
  ```json
  {
      "text": "I am excited to share that I have developed a Twitter API in Go!"
  }
  ```

### 2. Deleting a Tweet

To delete a tweet, the `/delete/{tweet_id}` endpoint is used. This endpoint accepts a `DELETE` request, where `{tweet_id}` is the unique ID of the tweet to be deleted. This can be found on the URL of the tweet E.g: https://x.com/FazalUrRehaman7/status/1844875797158412761. The program makes a `DELETE` request to the Twitter API to remove the specified tweet.

#### Example API Request:

- **Endpoint**: `DELETE http://localhost:8080/delete/1844875797158412761`

#### Example Response:

- **Response Body**:
  ```text
  Tweet deleted: 1844875797158412761
  ```

---

## Error Handling

### 1. API Errors

The program checks the status code of the response from Twitter's API to determine whether the request was successful:

- If the tweet is successfully posted, a `201 Created` status is expected.
- If the tweet is successfully deleted, a `200 OK` status is expected.

If the response status is not as expected, the program logs an error and exits:
```go
if resp.StatusCode != http.StatusCreated {
    log.Fatalf("Error response from Twitter API: %s", resp.Status)
}
```

### 2. Invalid Input Handling

The program handles invalid input (e.g., malformed JSON) by returning an HTTP `400 Bad Request` response to the client:
```go
if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
    http.Error(w, "Invalid request", http.StatusBadRequest)
    return
}
```

### 3. Tweet Not Found

If the `DELETE` request for a tweet fails (e.g., the tweet doesn't exist), the program returns a `404 Not Found` error to the client:
```go
if resp.StatusCode != http.StatusOK {
    log.Fatalf("Error response from Twitter API: %s", resp.Status)
    http.Error(w, "Tweet not found", http.StatusNotFound)
}
```

With this error handling approach, the program ensures graceful failure and communicates errors clearly to the client.

---
## Twitter Account:

**[@FazalUrRehaman7](https://twitter.com/FazalUrRehaman7)**
