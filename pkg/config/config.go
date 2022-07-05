package config

import (
	"flag"
	"os"
)

type Config struct {
	apiKey            string
	apiKeySecret      string
	accessToken       string
	accessTokenSecret string
	BearerToken       string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.apiKey, "apiKey", os.Getenv("API_KEY"), "API Key")
	flag.StringVar(&conf.apiKeySecret, "apiKeySecret", os.Getenv("API_KEY_SECRET"), "API Key Secret")
	flag.StringVar(&conf.accessToken, "accessToken", os.Getenv("ACCESS_TOKEN"), "Access Token")
	flag.StringVar(&conf.accessTokenSecret, "accessTokenSecret", os.Getenv("ACCESS_TOKEN_SECRET"), "Access Token Secret")
	flag.StringVar(&conf.BearerToken, "bearerToken", os.Getenv("BEARER_TOKEN"), "Bearer Token")

	flag.Parse()

	return conf
}

func (c *Config) GetApiKey() string {
	return c.apiKey
}

func (c *Config) GetApiKeySecret() string {
	return c.apiKeySecret
}

func (c *Config) GetAccessToken() string {
	return c.accessToken
}

func (c *Config) GetAccessTokenSecret() string {
	return c.accessTokenSecret
}
