package app

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) initializeRoutes() {
	var apiV1 = "/api/v1"

	// specification routes
	a.Router.HandleFunc(apiV1+"/specification/list", a.listConnectorSpecification).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/specification/test", a.testConnectorSpecification).Methods("POST", "OPTIONS")

	// core routes
	a.Router.HandleFunc(apiV1+"/user/{username}", a.getUsername).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/tweet", a.postTweet).Methods("POST", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/search", a.searchTweets).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/timeline", a.getTimeline).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/followers", a.getFollowers).Methods("GET", "OPTIONS")

	// Swagger
	a.Router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}
