/* Copyright (c) 2022. Ankit Patial.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 */

package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"roofix/config"
	"roofix/pkg/ws"
)

const (
	Connect    = "$connect"
	Disconnect = "$disconnect"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, e events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	config.PrintBuildInfo()
	connId := e.RequestContext.ConnectionID
	route := e.RequestContext.RouteKey

	params := e.QueryStringParameters

	response := func(code int, body string) (events.APIGatewayProxyResponse, error) {
		return events.APIGatewayProxyResponse{
			StatusCode: code,
			Body:       body,
		}, nil
	}

	switch route {
	case Connect:
		token := params["token"]
		ws.Register(ctx, connId, token)
	case Disconnect:
		ws.UnRegister(ctx, connId)
	default:
		err := ws.MsgFromClient(ctx, connId, e.Body)
		if err != nil {
			return response(http.StatusInternalServerError, err.Error())
		}
	}

	return response(http.StatusOK, "OK")
}
