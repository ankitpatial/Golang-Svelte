package server

//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"net/http"
	"roofix/pkg/account"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/util/log"
	"roofix/pkg/util/validate"
	"roofix/server/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// type Resolver struct{ client *ent.Client }
type Resolver struct{}

type UserCtxFunc func(r *http.Request) context.Context

// GqlServer is for graph api
func GqlServer(userCtx UserCtxFunc) func(w http.ResponseWriter, r *http.Request) {
	srv := handler.NewDefaultServer(schema())
	// limit query complexity
	srv.Use(extension.FixedComplexityLimit(60))
	// print operation name
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		oc := graphql.GetOperationContext(ctx)
		log.InfoBullet("GQL Operation: %s", oc.OperationName)
		return next(ctx)
	})

	// recover
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		log.PrintError(err.(error))
		return gqlerror.Errorf("Internal server error!")
	})

	// errors
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		var valErr *validate.StructValidationError
		if errors.As(e, &valErr) {
			err.Message = fmt.Sprintf("Validation errors:\n%s\n", valErr.AllErrors())
			err.Extensions = map[string]interface{}{
				"errors": valErr.Errors,
			}
		}

		return err
	})

	return func(w http.ResponseWriter, r *http.Request) {
		// set RealIP
		srv.ServeHTTP(w, r.WithContext(userCtx(r)))
	}
}

// Graphiql the Graph Playground
func Graphiql() func(w http.ResponseWriter, r *http.Request) {
	h := playground.Handler("GraphQL", "/query")
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

// schema creates a graphql executable schema.
func schema() graphql.ExecutableSchema {
	c := generated.Config{
		Resolvers: new(Resolver),
	}

	c.Directives.Authorize = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []*enum.Role) (res interface{}, err error) {
		// get session user
		user := account.CtxUser(ctx)
		if user == nil {
			// session user not found
			return nil, errors.New("401 Unauthorized")
		}

		// has roles count more than 0
		if len(roles) == 0 { // we just want to authorize, no need to check has role
			return next(ctx)
		}

		//user is in desired role
		for _, role := range roles {
			if role != nil && user.Role == *role {
				// match found, just call next
				return next(ctx)
			}
		}

		// failed, user is not authorize to go next
		return nil, msg.AsError(msg.NotAuthorized)
	}

	c.Directives.StringFor = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []*enum.Role) (res interface{}, err error) {
		// get session user
		user := account.CtxUser(ctx)
		if user == nil { // return empty if user info not found
			return "", nil
		}

		// check user has role allowed for current operation
		for _, role := range roles {
			if role != nil && user.Role == *role {
				// on match just call next
				return next(ctx)
			}
		}

		return "", nil
	}

	c.Directives.FloatFor = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []*enum.Role) (res interface{}, err error) {
		// get session user
		user := account.CtxUser(ctx)
		if user == nil { // return empty if user info not found
			return 0.0, nil
		}

		// check user has role allowed for current operation
		for _, role := range roles {
			if role != nil && user.Role == *role {
				// on match just call next
				return next(ctx)
			}
		}

		return 0.0, nil
	}

	return generated.NewExecutableSchema(c)
}
