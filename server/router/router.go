/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package router

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"io"
	"net/http"
	"net/url"
	"roofix/config"
	"roofix/pkg/account"
	"roofix/pkg/audit"
	"roofix/pkg/document"
	"roofix/pkg/enum"
	"roofix/pkg/util/log"
	"roofix/pkg/util/storage"
	ws "roofix/pkg/ws/dev"
	"roofix/server"
	"roofix/server/router/api"
	"roofix/server/router/response"
	"strings"
	"time"
)

var conf *config.Config

func init() {
	conf = config.Read()
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	// router middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{conf.Website.Url},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Current", "lat-CSRF-Token", "lat-Real-IP"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.RealIP)
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "RealIP", r.RemoteAddr)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(secure)
	r.Use(middleware.Compress(5, "application/json", "text/html", "text/plain"))

	// home route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response.Ok(w, "Hello World!")
	})

	r.Post("/set-password", setPassword)
	r.Post("/login", login)
	r.Delete("/logout", logout)

	// graph api
	r.Post("/query", server.GqlServer(withUserCtx))

	//if !conf.IsProdEnv() { // show in DEV & Staging env
	r.Get("/graphiql", server.Graphiql())
	//}

	if conf.IsDevEnv() { // DEV env only
		hub := ws.NewHub()
		go hub.Run()
		r.Get("/ws", ws.ServeWs(hub))
		r.Post("/ws-msg", ws.HandleWsMsg(hub))
		r.Get("/file-download", devFileDownload)
		r.Post("/file-upload", devFileUpload)
		r.Get("/PublicData/{id}/*", publicData)
	}

	// oAuth api to post job
	api.Register(r)
	return r
}

func login(w http.ResponseWriter, r *http.Request) {
	// bind input
	var inp account.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		response.BadRequest(w, err)
		return
	}

	ctx := r.Context()

	// authenticate
	u, err := account.Authenticate(ctx, &inp)
	if err != nil {
		audit.OpFailed(ctx, audit.AccLogin, err)
		response.BadRequest(w, err)
		return
	}

	newUserSession(w, r, u.ID, u.Role)
}

func setPassword(w http.ResponseWriter, r *http.Request) {
	var inp account.SetPasswordInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		response.BadRequest(w, err)
		return
	}
	ctx := r.Context()
	u, err := account.SetPassword(ctx, &inp)
	if err != nil {
		audit.OpFailed(ctx, audit.AccChangePassword, err)
		response.BadRequest(w, err)
	}
	audit.OpSuccess(ctx, audit.AccChangePassword, fmt.Sprintf("password changed for user: %s", u.ID))

	newUserSession(w, r, u.ID, u.Role)
}

func newUserSession(w http.ResponseWriter, r *http.Request, uID string, role enum.Role) {
	// create session detail in DB
	ctx := r.Context()
	expiry := time.Now().Add(time.Hour * 24 * 7) // 7 day
	sid, err := account.NewSession(ctx, uID, role, expiry, r.RemoteAddr)
	if err != nil {
		audit.UserOpFailed(ctx, uID, audit.AccLogin, err)
		response.ServerError(w, err)
		return
	}

	// session cookie
	http.SetCookie(w, setSessionCookie(sid, expiry))
	audit.UserOpSuccess(ctx, uID, audit.AccLogin, "logged in")

	u := account.SessionUser(ctx, sid)
	response.Ok(w, u)
}

func logout(w http.ResponseWriter, r *http.Request) {
	ctx := withUserCtx(r)
	clearSessionCookie(w, r)
	audit.Operation(ctx, audit.AccLogout, "logout", nil)
	response.Ok(w, nil)
}

func withUserCtx(r *http.Request) context.Context {
	var sid string

	ctx := r.Context()
	bt := r.Header.Get("Authorization")
	if bt == "" { // bearer is missing, read COOKIE
		sid = readSessionCookie(r)
	} else { // BEARER token
		t, err := stripBearer(bt)
		if err != nil {
			log.Error(err)
			return ctx
		}

		// parse token
		sid, err = account.ParseAuthToken(t)
		if err != nil {
			log.Error(err)
			return ctx
		}
	}

	if sid == "" { // no session ID found
		return ctx
	}

	// pull session user data
	user := account.SessionUser(r.Context(), sid) // session cookie found
	if user == nil {                              // may come empty if user session is null or expired
		return r.Context()
	}

	user.SessionID = sid
	// and call the next with our new context
	return account.CtxWithUser(r.Context(), user)
}

func stripBearer(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}

// devFileDownload to get file from storage on dev machine
func devFileDownload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	q, _ := url.ParseQuery(r.URL.RawQuery)
	bucket := q.Get("bucket")
	key := q.Get("key")

	b, err := storage.GetObject(ctx, bucket, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Content-Disposition: inline
	// Content-Disposition: attachment; filename="filename.jpg"
	// Content-Length
	w.Header().Set("Content-Disposition", "inline")
	//w.Header().Set("Content-Current", )
	w.WriteHeader(http.StatusOK)
	io.Copy(w, bytes.NewReader(b))
}

// devFileUpload to upload file to storage on dev machine
func devFileUpload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	q, _ := url.ParseQuery(r.URL.RawQuery)
	bucket := q.Get("bucket")
	key := q.Get("key")

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := storage.PutObject(ctx, bucket, key, buf.Bytes(), nil); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = document.SetReady(ctx, bucket, key)

	// simulate delay
	time.Sleep(time.Second * 2)
}

func publicData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	bucket := config.DataBucket()
	key := r.URL.Path[1:]

	b, err := storage.GetObject(ctx, bucket, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Content-Disposition: inline
	// Content-Disposition: attachment; filename="filename.jpg"
	// Content-Length
	w.Header().Set("Content-Disposition", "inline")
	//w.Header().Set("Content-Current", )
	w.WriteHeader(http.StatusOK)
	io.Copy(w, bytes.NewReader(b))
}
