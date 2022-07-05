package main

import (
	"fmt"

	"github.com/91pavan/twitter-connector/pkg/app"
	"github.com/91pavan/twitter-connector/pkg/logger"

	_ "github.com/91pavan/twitter-connector/docs"
)

var (
	log  = logger.New("app", "twitter-connector")
	port = 8005
)

// @title Twitter Connector API
// @version 2.0
// @description This is a Kosha REST serice for exposing many twitter features as REST APIs with better consistency, observability
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email eti@cisco.io
// @host localhost:8005
// @BasePath /
func main() {

	a := app.App{}
	a.Initialize(log)
	log.Infof("Running twitter-connector on port %d", port)
	a.Run(fmt.Sprintf(":%d", port))

}
