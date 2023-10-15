package config

import (
	"fmt"
	"net/url"
)

// Config has all the app level configurations
type Config struct {
	Env             string
	SessionCookieId string     `json:"sid"`
	Secret          string     `json:"secret"`
	GoogleApiKey    string     `json:"googleApiKey"`
	Website         Server     `json:"website"`
	Cookie          *Cookie    `json:"cookie"`
	DB              DbSettings `json:"db"`
	AWS             struct {
		Region      string `json:"region"`
		ProfileName string `json:"profileName"`
	} `json:"aws"`
	Mail struct {
		Sender string `json:"sender"`
	} `json:"mail"`
	Infra *Infra `json:"infra"`
}

func (c *Config) IsDevEnv() bool {
	return c.Env == AppEnvDev
}

func (c *Config) IsProdEnv() bool {
	return c.Env == AppEnvProd
}

// Server related info
type Server struct {
	Url      string `json:"url"`
	AssetUrl string `json:"assetUrl"`
	Port     int16  `json:"port"`
}

func (c *Server) Domain() string {
	u, _ := url.Parse(c.Url)

	return u.Hostname()
}

// DbSettings of app
type DbSettings struct {
	Engine    string `json:"engine"`
	Host      string `json:"host"`
	Port      uint16 `json:"port"`
	Database  string `json:"database"`
	User      string `json:"user"`
	Password  string `json:"password"`
	EnableTLS bool
}

func (o *DbSettings) ConnStr() string {
	var ssl string
	if o.EnableTLS {
		ssl = "&tls=true"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&multiStatements=true%s", o.User, o.Password, o.Host, o.Port, o.Database, ssl)
}

func (o *DbSettings) ConnStrFull() string {
	var ssl string
	if o.EnableTLS {
		ssl = "&tls=true"
	}
	return fmt.Sprintf("%s://%s:%s@tcp(%s:%d)/%s?parseTime=true&multiStatements=true%s", o.Engine, o.User, o.Password, o.Host, o.Port, o.Database, ssl)
}

// SecretValue on aws deployed DB
type SecretValue struct {
	AppSecret           string `json:"appSecret"`
	StateSecret         string `json:"stateSecret"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
	DbEngin             string `json:"engin"`
	DbName              string `json:"dbname"`
	DbHost              string `json:"host"`
	DbPort              uint16 `json:"port"`
	DbUser              string `json:"username"`
	DbPwd               string `json:"password"`
}

type Cookie struct {
	Previous *CookieSecret `json:"previous"`
	Current  *CookieSecret `json:"current"`
}

type CookieSecret struct {
	HashKey  string `json:"hashKey"`
	BlockKey string `json:"blockKey"`
}

type Infra struct {
	Stack struct {
		Name    string `json:"name"`
		Account string `json:"account"`
		Region  string `json:"region"`
	} `json:"stack"`
	Domain            string `json:"domain"`
	WebsiteDomain     string `json:"websiteDomain"`
	WsDomain          string `json:"wsDomain"`
	AssetsDomain      string `json:"assetsDomain"`
	AssetsBucket      string `json:"assetsBucket"`
	DataBucket        string `json:"dataBucket"`
	LogBucket         string `json:"logBucket"`
	MailQueue         string `json:"mailQueue"`
	TaskQueue         string `json:"taskQueue"`
	NotificationQueue string `json:"notificationQueue"`
	DataSyncQueue     string `json:"dataSyncQueue"`
	ErrorQueue        string `json:"errorQueue"`
	HostedZoneID      string `json:"hostedZoneID"`
	HostedZone        string `json:"hostedZone"`
	DB                struct {
		SecretName string `json:"secretName"`
		Name       string `json:"name"`
		Username   string `json:"username"`
	} `json:"db"`
}
