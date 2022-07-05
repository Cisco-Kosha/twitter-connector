package app

import (
	"github.com/dghubble/oauth1"
	"log"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/gorilla/mux"

	"github.com/91pavan/twitter-connector/pkg/config"
	"github.com/91pavan/twitter-connector/pkg/logger"
)

type App struct {
	Router *mux.Router
	Client *twitter.Client
	Log    logger.Logger
	Cfg    *config.Config
}

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}

// Initialize creates the necessary scaffolding of the app
func (a *App) Initialize(log logger.Logger) {

	cfg := config.Get()

	a.Cfg = cfg
	a.Log = log

	a.Router = router()

	oauth1Config := oauth1.NewConfig(cfg.GetApiKey(), cfg.GetApiKeySecret())
	oauth1Token := oauth1.NewToken(cfg.GetAccessToken(), cfg.GetAccessTokenSecret())
	httpClient := oauth1Config.Client(oauth1.NoContext, oauth1Token)

	// Twitter client
	a.Client = twitter.NewClient(httpClient)

	a.initializeRoutes()

}

// Run starts the app and serves on the specified addr
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
