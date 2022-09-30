package main

import (
	"encoding/json"
	"github.com/91pavan/twitter-connector/pkg/app"
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
	resp := map[string]string{"API_KEY": "Access key",
		"API_KEY_SECRET":      "Access key Secret",
		"ACCESS_TOKEN":        "Access Token",
		"ACCESS_TOKEN_SECRET": "Access Token Secret",
	}
	res, _ := json.Marshal(resp)
	assert.Equal(t, res, w.Body.Bytes())
}
