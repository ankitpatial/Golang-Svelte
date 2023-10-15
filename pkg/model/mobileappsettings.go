package model

type MobileAppSettings struct {
	LogoURL      string   `json:"logoURL"`
	PrimaryColor string   `json:"primaryColor"`
	HideTabs     []string `json:"hideTabs"`
}
