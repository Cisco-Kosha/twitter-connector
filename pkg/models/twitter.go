package models

type Tweet struct {
	Text string `json:"text,omitempty"`
}

type Specification struct {
	ApiKey            string `json:"api_key,omitempty"`
	ApiKeySecret      string `json:"api_key_secret,omitempty"`
	AccessToken       string `json:"access_token,omitempty"`
	AccessTokenSecret string `json:"access_token_secret,omitempty"`
}
