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
	res, _ := json.Marshal(fields)
	assert.Equal(t, res, w.Body.Bytes())
}
