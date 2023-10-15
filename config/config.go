package config

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"log"
	"os"
	"path/filepath"
	"roofix/pkg/util/ptr"
)

var (
	//go:embed *.json
	files embed.FS
	conf  *Config
	// BuildUser will set on build time to track binary built by user
	BuildUser = "sys"
	// BuildTime is the timestamp the build was created
	BuildTime = ""
	// UseSecretMgr will be set on build time to allow reading stored secrets
	UseSecretMgr = "false"
	// AppEnv is runtime config env flag. Will be set on build time also
	AppEnv = "development"
)

const (
	// AppName is the module name
	AppName       = "roofix"
	AppEnvDev     = "development"
	AppEnvProd    = "production"
	AppEnvStage   = "stage"
	UserCtxKey    = "auth-user"
	APIUserCtxKey = "oauth.claims"
	IPCtxKey      = "RealIP"
	MoneySymbol   = "$"
)

func PrintBuildInfo() {
	env := ""
	if conf != nil {
		env = conf.Env
	}

	fmt.Printf(" == Env: %s  Build: %s by %s ==\n", env, BuildTime, BuildUser)
}

// Read config file as per the app env. config will be cached to avoid over read.
func Read() *Config {
	return ReadByEnv(AppEnv)
}

func ReadByEnv(env string) *Config {
	if conf != nil { // already in memory
		return conf
	}

	data, err := files.ReadFile(fmt.Sprintf("%s.json", env))
	if err != nil {
		log.Fatalf("failed to read config file %s", err)
	}

	conf = &Config{
		Env: env,
	}

	if err := json.Unmarshal(data, conf); err != nil {
		fmt.Printf("failed to unmarshal %v\n", err)
		return conf
	}

	// dev env check config.local.json for local db user & password
	if IsDevEnv() {
		data, _ = files.ReadFile("local.json")
		if data != nil {
			fmt.Println("found local.json")
			var local map[string]string
			if err := json.Unmarshal(data, &local); err == nil { // no error
				conf.DB.User = local["dbUser"]
				conf.DB.Password = local["dbPassword"]
			}
		}

		if u := os.Getenv("DB_USER"); u != "" {
			conf.DB.User = u
		}

		if p := os.Getenv("DB_PASSWORD"); p != "" {
			conf.DB.Password = p
		}
	}

	if UseSecretMgr == "true" { // ignore intellisense warning, UseSecretMgr will set on build time
		// read from secret manager
		secret := fromSecretMgr()
		if secret == nil {
			return conf
		}

		proxy := os.Getenv("DB_PROXY")
		if proxy != "" {
			conf.DB.EnableTLS = true
			conf.DB.Host = proxy
		} else {
			conf.DB.Host = secret.DbHost
		}

		conf.DB.Port = secret.DbPort
		conf.DB.Database = secret.DbName
		conf.DB.User = secret.DbUser
		conf.DB.Password = secret.DbPwd
	}

	//fmt.Printf("%v", conf.DB)
	return conf
}

// readDbSecret will read aws secret manager to get deployed db settings
func fromSecretMgr() *SecretValue {
	cl := secretsmanager.NewFromConfig(AwsConfig())
	res, err := cl.GetSecretValue(context.Background(), &secretsmanager.GetSecretValueInput{
		SecretId: ptr.Str(DbSecretName()),
	})

	if err != nil {
		fmt.Printf("failed to GetSecretValue %v\n", err)
		return nil
	}

	if res.SecretString == nil {
		return nil
	}

	var secret SecretValue
	if err := json.Unmarshal([]byte(*res.SecretString), &secret); err != nil {
		fmt.Printf("failed to Unmarshal %v\n", err)
		return nil
	}

	return &secret
}

func IsDevEnv() bool {
	return AppEnv == AppEnvDev // AppEnv is changed on build as per the build env
}

func DbSecretName() string {
	return Read().Infra.DB.SecretName
}

func WsDomain() string {
	return Read().Infra.WsDomain
}

func DataBucket() string {
	return Read().Infra.DataBucket
}

func LogBucket() string {
	return Read().Infra.LogBucket
}

func MailQueueName() string {
	return Read().Infra.MailQueue
}

func TaskQueueName() string {
	return Read().Infra.TaskQueue
}

func NotificationQueueName() string {
	return Read().Infra.NotificationQueue
}

// AppHome is path to app home directory
func AppHome() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dirPath := filepath.Join(home, "."+AppName)
	if _, err = os.Stat(dirPath); os.IsNotExist(err) { // create dir
		err = os.Mkdir(dirPath, 0750)
	}
	return dirPath, nil
}

func AwsConfig() aws.Config {
	var cfg aws.Config
	var err error

	conf := Read()
	profile := conf.AWS.ProfileName
	if profile != "" {
		cfg, err = config.LoadDefaultConfig(
			context.Background(),
			config.WithRegion(conf.AWS.Region),
			config.WithSharedConfigProfile(conf.AWS.ProfileName))
	} else {
		cfg, err = config.LoadDefaultConfig(
			context.Background(),
			config.WithRegion(conf.AWS.Region),
			config.WithSharedConfigProfile(""))
	}

	if err != nil {
		panic(err)
	}

	return cfg
}
