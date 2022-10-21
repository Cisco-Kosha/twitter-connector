package app

import (
	"encoding/json"
	"github.com/91pavan/twitter-connector/pkg/models"
	"github.com/dghubble/oauth1"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/gorilla/mux"
)

// ListConnectorSpecification godoc
// @Summary Get details of the connector specification
// @Description Get all environment variables that need to be supplied
// @Tags specification
// @Accept  json
// @Produce  json
// @Success 200 {object} object
// @Router /api/v1/specification/list [get]
func (a *App) ListConnectorSpecification(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	fields := &models.SpecificationV2{}

	apiKey := models.FieldSpecificationV2{
		Field:       "API_KEY",
		Type:        "String",
		MinLength:   3,
		MaxLength:   40,
		Required:    true,
		Description: "Consumer Key",
		Example:     "TmiVYoLyeUyX9MBdKjtqQZjCB",
	}
	fields.Fields = append(fields.Fields, apiKey)
	apiSecret := models.FieldSpecificationV2{
		Field:       "API_KEY_SECRET",
		Type:        "String",
		MinLength:   10,
		MaxLength:   60,
		Required:    true,
		Description: "Consumer Secret",
		Example:     "oA8f7fXLkZ7dKGukms1yFUNDdhlqd9koJNGfDaedpAoeaoepBd",
	}
	fields.Fields = append(fields.Fields, apiSecret)

	accessToken := models.FieldSpecificationV2{
		Field:       "ACCESS_TOKEN",
		Type:        "String",
		MinLength:   10,
		MaxLength:   60,
		Required:    true,
		Description: "Twitter Developer Access Token",
		Example:     "337265965-fPOgoWoSdSQxM8VE8GfvKzGYrWhX9WjliZIY0ENS",
	}
	fields.Fields = append(fields.Fields, accessToken)

	accessTokenSecret := models.FieldSpecificationV2{
		Field:       "ACCESS_TOKEN_SECRET",
		Type:        "String",
		MinLength:   10,
		MaxLength:   60,
		Required:    true,
		Description: "Twitter Developer Access Token Secret",
		Example:     "x4ZocH8PIYRFBF1bJJ7HgetSuGqk8NUNYspafEIX4cdi1",
	}
	fields.Fields = append(fields.Fields, accessTokenSecret)

	respondWithJSON(w, http.StatusOK, fields)
}

// testConnectorSpecification godoc
// @Summary Test if access tokens work against the specification
// @Description Check if user account can be verified
// @Tags specification
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Specification
// @Router /api/v1/specification/test [post]
func (a *App) testConnectorSpecification(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	if (*r).Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	var t models.Specification
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		a.Log.Errorf("Error parsing json payload", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	config := oauth1.NewConfig(t.ApiKey, t.ApiKeySecret)
	token := oauth1.NewToken(t.AccessToken, t.AccessTokenSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "User not verified")
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

// getUsername godoc
// @Summary Get details of the twitter user
// @Description Get profile information of the twitter user
// @Tags users
// @Accept  json
// @Produce  json
// @Param username path string true "Enter username"
// @Success 200 {object} twitter.User
// @Router /api/v1/user/{username} [get]
func (a *App) getUsername(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	username := vars["username"]
	user, _, err := a.Client.Users.Show(&twitter.UserShowParams{
		ScreenName: username,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

// getTimeline godoc
// @Summary Get tweets from the twitter user's timeline
// @Description Get all homepage tweets from twitter's main user timeline
// @Tags timeline
// @Accept  json
// @Produce  json
// @Success 200 {object} []twitter.Tweet
// @Router /api/v1/timeline [get]
func (a *App) getTimeline(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	timeline, _, err := a.Client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, timeline)
}

// getFollowers godoc
// @Summary Get followers from the twitter user's profile
// @Description Get all followers list from twitter's main user profile
// @Tags followers
// @Accept  json
// @Produce  json
// @Success 200 {object} twitter.Followers
// @Router /api/v1/followers [get]
func (a *App) getFollowers(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	followers, _, err := a.Client.Followers.List(&twitter.FollowerListParams{})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, followers)
}

// searchTweets godoc
// @Summary Search all tweets for a given query text
// @Description Search tweets
// @Tags search
// @Accept  json
// @Produce  json
// @Param query query string true "Enter query string"
// @Success 200 {object} twitter.Search
// @Router /api/v1/search [get]
func (a *App) searchTweets(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	query := r.URL.Query().Get("query")

	if query != "" {
		res, _, err := a.Client.Search.Tweets(&twitter.SearchTweetParams{
			Query: query,
		})
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, res)
	}

	respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "query parameter empty"})
}

// postTweet godoc
// @Summary Post a tweet on user's timeline
// @Description Submit a new tweet
// @Tags tweets
// @Accept  json
// @Produce  json
// @Param text body models.Tweet true "Enter tweet string"
// @Success 200 {object} twitter.Tweet
// @Router /api/v1/tweet [post]
func (a *App) postTweet(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	if (*r).Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	var t models.Tweet
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		a.Log.Errorf("Error parsing json payload", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	a.Log.Infof("Posting a new tweet: %+v", t)
	tw, _, err := a.Client.Statuses.Update(t.Text, nil)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Unable to post tweet. Error: "+err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, tw)
}
