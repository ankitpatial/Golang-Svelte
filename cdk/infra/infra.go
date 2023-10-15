package infra

import (
	"roofix/config"
	"strings"
)

func SuffixIt(appEnv, name string) string {
	switch appEnv {
	case config.AppEnvProd:
		return name + "-PROD"
	case config.AppEnvStage:
		return name + "-stage"
	default:
		return name + "-" + strings.ToUpper(appEnv)
	}
}
