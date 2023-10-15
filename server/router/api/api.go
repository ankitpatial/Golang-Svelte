/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"roofix/config"
	"roofix/pkg/rfx"
	"roofix/server/router/oauth"
	"time"
)

func Register(r *chi.Mux) {
	secret := config.Read().Secret
	d := time.Hour * 24 * 7
	s := oauth.NewBearerServer(secret, d, &oauth.Verify{}, nil)

	r.Post("/token", s.ApiUserCred)

	r.Group(func(r chi.Router) {
		r.Use(oauth.Authorize(secret, nil))
		r.Get("/external/job/{id}", getJob)
		r.Post("/external/job", postJob)

		// RFX Bubble app endpoints
		r.Post("/rfx/request-an-estimate", rfx.EstimateRequestHandler)
		r.Post("/rfx/estimate-approved", rfx.EstimateApprovedHandler)
		r.Post("/rfx/save-company", rfx.SaveCompanyHandler)
		r.Post("/rfx/save-user", rfx.SaveUserHandler)
		// ~~
	})

	r.Group(func(r chi.Router) {
		r.Use(EagleViewIP)
		r.Get("/eagleview/OrderStatusUpdate", eagleViewOrderStatusUpdate)
		r.Get("/eagleview/NeedToId", eagleViewNeedToId)
	})

	//r.Group(func(r chi.Router) { // TODO: secure the endpoints
	//	r.Post("/FileDelivery", fileDelivery)
	//	r.Get("/FileDeliveryConfirmation", fileDeliveryConfirm)
	//})

	r.Group(func(r chi.Router) {
		r.Use(oauth.Authorize(secret, nil))

	})
}

func EagleViewIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.IsDevEnv() { // for dev just by-pass
			next.ServeHTTP(w, r)
			return
		}

		//for production allow certain ips only
		// TODO
		//ip := account.CtxIP(ctx)
		next.ServeHTTP(w, r)
	})
}
