package main

import (
	"encoding/json"
	"github.com/Cisco-Kosha/twitter-connector/pkg/app"
	"github.com/Cisco-Kosha/twitter-connector/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequest(t *testing.T) {
	t.Parallel()

	r, _ := http.NewRequest("GET", "/api/v1/specification/list", nil)
	w := httptest.NewRecorder()

	a := app.App{}
	a.ListConnectorSpecification(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	fields := &models.SpecificationV2{}

	apiKey := models.FieldSpecificationV2{
		Field:       "API_KEY",
		DisplayName: "Consumer Key or API Key",
		Type:        "String",
		MinLength:   3,
		MaxLength:   40,
		Required:    true,
		Description: "Consumer Key",
		Example:     "TmiVYoLyeUyX9MBdKjtqQZjCB",
		HelpSection: "Navigate to the developer portal\nExpand the 'Projects and Apps' dropdown in the sidenav\nOpen the App which is associated with the API Key and Secret that you would like to find or regenerate\nNavigate to the Keys and tokens tab\nFrom there, you will find all of the credentials associated with your App. ",
	}
	fields.Fields = append(fields.Fields, apiKey)
	apiSecret := models.FieldSpecificationV2{
		Field:       "API_KEY_SECRET",
		DisplayName: "Consumer Key Secret or Consumer Key Secret",
		Type:        "String",
		MinLength:   10,
		MaxLength:   60,
		Required:    true,
		Description: "Consumer Secret",
		Example:     "oA8f7fXLkZ7dKGukms1yFUNDdhlqd9koJNGfDaedpAoeaoepBd",
		HelpSection: "Navigate to the developer portal\nExpand the 'Projects and Apps' dropdown in the sidenav\nOpen the App which is associated with the API Key and Secret that you would like to find or regenerate\nNavigate to the Keys and tokens tab",
	}
	fields.Fields = append(fields.Fields, apiSecret)

	accessToken := models.FieldSpecificationV2{
		Field:       "ACCESS_TOKEN",
		DisplayName: "Access Token",
		Type:        "String",
		MinLength:   10,
		MaxLength:   60,
		Required:    true,
		Description: "Twitter Developer Access Token",
		Example:     "337265965-fPOgoWoSdSQxM8VE8GfvKzGYrWhX9WjliZIY0ENS",
		HelpSection: "Login to your Twitter account on developer.twitter.com.\nNavigate to the Twitter app dashboard and open the Twitter app for which you would like to generate access tokens.\nNavigate to the \"Keys and Tokens\" page.\nSelect 'Create' under the \"Access token & access token secret\" section.",
	}
	fields.Fields = append(fields.Fields, accessToken)

	accessTokenSecret := models.FieldSpecificationV2{
		Field:       "ACCESS_TOKEN_SECRET",
		DisplayName: "Access Token Secret",
		Type:        "String",
		MinLength:   10,
		MaxLength:   60,
		Required:    true,
		Description: "Twitter Developer Access Token Secret",
		Example:     "x4ZocH8PIYRFBF1bJJ7HgetSuGqk8NUNYspafEIX4cdi1",
		HelpSection: "Login to your Twitter account on developer.twitter.com.\nNavigate to the Twitter app dashboard and open the Twitter app for which you would like to generate access tokens.\nNavigate to the \"Keys and Tokens\" page.\nSelect 'Create' under the \"Access token & access token secret\" section.",
	}
	fields.Fields = append(fields.Fields, accessTokenSecret)
	res, _ := json.Marshal(fields)
	assert.Equal(t, res, w.Body.Bytes())
}
