package app

import (
	"context"
	"roofix/config"
	"roofix/pkg/util/parse"
)

type CtxData struct {
	User    *User    `json:"user"`
	APIUser *APIUser `json:"apiUser"`
	IP      string   `json:"ip"`
}

type APIUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

func SetCtx(ctx context.Context, data CtxData) context.Context {
	// with IP
	ctx = context.WithValue(ctx, config.IPCtxKey, data.IP)

	// with user
	if data.User != nil {
		ctx = context.WithValue(ctx, config.UserCtxKey, data.User)
	}

	if data.APIUser != nil {
		ctx = context.WithValue(ctx, config.APIUserCtxKey, map[string]string{
			"apiUID":   data.APIUser.ID,
			"apiUName": data.APIUser.Name,
		})
	}

	return ctx
}

func ReadCtx(ctx context.Context) CtxData {
	return CtxData{
		User:    CtxUser(ctx),
		IP:      ctxIP(ctx),
		APIUser: CtxAPIUser(ctx),
	}
}

func ctxIP(ctx context.Context) string {
	ip, _ := ctx.Value("RealIP").(string)
	return ip
}

func CtxUser(ctx context.Context) *User {
	v := ctx.Value(config.UserCtxKey)
	if v == nil {
		return nil
	}

	m := parse.StructToMap(v)
	u := &User{
		ID:        m["id"].(string),
		FirstName: m["firstName"].(string),
		LastName:  m["lastName"].(string),
		Email:     m["email"].(string),
		Role:      m["role"].(string),
	}
	return u
}

func CtxAPIUser(ctx context.Context) *APIUser {
	v := ctx.Value(config.APIUserCtxKey)
	if v == nil {
		return nil
	}

	m := parse.StructToMap(v)
	u := &APIUser{
		ID: m["apiUID"].(string),
	}

	if name, ok := m["apiUName"]; ok {
		u.Name = name.(string)
	}

	return u
}
