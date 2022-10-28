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

type SpecificationV2 struct {
	Fields []FieldSpecificationV2 `json:"fields,omitempty"`
}

type FieldSpecificationV2 struct {
	Field       string `json:"field,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Type        string `json:"type,omitempty"`
	MinLength   int    `json:"minLength,omitempty"`
	MaxLength   int    `json:"maxLength,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Description string `json:"description,omitempty"`
	Example     string `json:"example,omitempty"`
	HelpSection string `json:"helpSection,omitempty"`
}
