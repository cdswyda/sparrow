package datatypes

type Session struct {
	User         SessionUser `json:"user"`
	Expires      string      `json:"expires"`
	AccessToken  string      `json:"accessToken"`
	AuthProvider string      `json:"authProvider"`
}

type SessionUser struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Image        string   `json:"image"`
	Picture      string   `json:"picture"`
	Mfa          bool     `json:"mfa"`
	Groups       []string `json:"groups"`
	IntercomHash string   `json:"intercom_hash"`
}
