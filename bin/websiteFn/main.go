/*
 * Copyright (c) 2022. SimSaw Software Pvt. Ltd. All Right Reserved.
 * Author: Ankit Patial
 */

package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	adapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"net/http"
	"roofix/config"
	"roofix/pkg/util/storage"
	"roofix/server/router"
	"roofix/server/router/response"
)

var proxy *adapter.ChiLambda

func init() {
	var r = router.NewRouter()

	r.Get("/favicon.*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		w.Write(getResource(r.Context(), "favicon.ico"))
	})

	r.Get("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		res := getResource(r.Context(), "robots.txt")
		response.String(w, r, string(res))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res := getResource(r.Context(), "index.html")
		response.Html(w, r, string(res))
	})

	r.Get("/terms-of-use", func(w http.ResponseWriter, r *http.Request) {
		res := getResource(r.Context(), "terms-of-use.html")
		response.Html(w, r, string(res))
	})

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		res := getResource(r.Context(), "fallback.html")
		response.Html(w, r, string(res))
	})

	proxy = adapter.New(r)
}

func main() {
	config.PrintBuildInfo()
	lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return proxy.ProxyWithContext(ctx, req)
	})
}

func getResource(ctx context.Context, key string) []byte {
	d, _ := storage.GetObject(ctx, config.Read().Infra.AssetsBucket, key)
	return d
}
