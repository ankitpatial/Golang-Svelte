// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent"
	"roofix/pkg/apiaccess"
	"strconv"
	"sync"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _ApiAccess_id(ctx context.Context, field graphql.CollectedField, obj *ent.ApiAccess) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ApiAccess_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ApiAccess_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ApiAccess",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ApiAccess_url(ctx context.Context, field graphql.CollectedField, obj *ent.ApiAccess) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ApiAccess_url(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.URL, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ApiAccess_url(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ApiAccess",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ApiAccess_username(ctx context.Context, field graphql.CollectedField, obj *ent.ApiAccess) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ApiAccess_username(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Username, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ApiAccess_username(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ApiAccess",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ApiAccess_password(ctx context.Context, field graphql.CollectedField, obj *ent.ApiAccess) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ApiAccess_password(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Password, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ApiAccess_password(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ApiAccess",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ApiAccess_key(ctx context.Context, field graphql.CollectedField, obj *ent.ApiAccess) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ApiAccess_key(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Key, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ApiAccess_key(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ApiAccess",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ApiAccess_secret(ctx context.Context, field graphql.CollectedField, obj *ent.ApiAccess) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ApiAccess_secret(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Secret, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ApiAccess_secret(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ApiAccess",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ApiAccessConnection_edges(ctx context.Context, field graphql.CollectedField, obj *ent.ApiAccessConnection) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ApiAccessConnection_edges(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Edges, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*ent.ApiAccessEdge)
	fc.Result = res
	return ec.marshalOApiAccessEdge2ᚕᚖroofixᚋentᚐApiAccessEdgeᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ApiAccessConnection_edges(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ApiAccessConnection",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "node":
				return ec.fieldContext_ApiAccessEdge_node(ctx, field)
			case "cursor":
				return ec.fieldContext_ApiAccessEdge_cursor(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type ApiAccessEdge", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _ApiAccessConnection_pageInfo(ctx context.Context, field graphql.CollectedField, obj *ent.ApiAccessConnection) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ApiAccessConnection_pageInfo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PageInfo, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(entgql.PageInfo[string])
	fc.Result = res
	return ec.marshalNPageInfo2entgoᚗioᚋcontribᚋentgqlᚐPageInfo(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ApiAccessConnection_pageInfo(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ApiAccessConnection",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "hasNextPage":
				return ec.fieldContext_PageInfo_hasNextPage(ctx, field)
			case "endCursor":
				return ec.fieldContext_PageInfo_endCursor(ctx, field)
			case "hasPreviousPage":
				return ec.fieldContext_PageInfo_hasPreviousPage(ctx, field)
			case "startCursor":
				return ec.fieldContext_PageInfo_startCursor(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type PageInfo", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _ApiAccessConnection_totalCount(ctx context.Context, field graphql.CollectedField, obj *ent.ApiAccessConnection) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ApiAccessConnection_totalCount(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.TotalCount, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ApiAccessConnection_totalCount(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ApiAccessConnection",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ApiAccessEdge_node(ctx context.Context, field graphql.CollectedField, obj *ent.ApiAccessEdge) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ApiAccessEdge_node(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Node, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*ent.ApiAccess)
	fc.Result = res
	return ec.marshalOApiAccess2ᚖroofixᚋentᚐApiAccess(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ApiAccessEdge_node(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ApiAccessEdge",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_ApiAccess_id(ctx, field)
			case "url":
				return ec.fieldContext_ApiAccess_url(ctx, field)
			case "username":
				return ec.fieldContext_ApiAccess_username(ctx, field)
			case "password":
				return ec.fieldContext_ApiAccess_password(ctx, field)
			case "key":
				return ec.fieldContext_ApiAccess_key(ctx, field)
			case "secret":
				return ec.fieldContext_ApiAccess_secret(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type ApiAccess", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _ApiAccessEdge_cursor(ctx context.Context, field graphql.CollectedField, obj *ent.ApiAccessEdge) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ApiAccessEdge_cursor(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Cursor, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(entgql.Cursor[string])
	fc.Result = res
	return ec.marshalOCursor2entgoᚗioᚋcontribᚋentgqlᚐCursor(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ApiAccessEdge_cursor(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ApiAccessEdge",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Cursor does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputApiAccessInput(ctx context.Context, obj interface{}) (apiaccess.Input, error) {
	var it apiaccess.Input
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "name", "url", "username", "password", "key", "secret"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			data, err := ec.unmarshalNID2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ID = data
		case "name":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Name = data
		case "url":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("url"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.URL = data
		case "username":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("username"))
			data, err := ec.unmarshalOString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Username = data
		case "password":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("password"))
			data, err := ec.unmarshalOString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Password = data
		case "key":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("key"))
			data, err := ec.unmarshalOString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Key = data
		case "secret":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("secret"))
			data, err := ec.unmarshalOString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Secret = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputApiAccessOrder(ctx context.Context, obj interface{}) (ent.ApiAccessOrder, error) {
	var it ent.ApiAccessOrder
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"direction", "field"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "direction":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("direction"))
			data, err := ec.unmarshalNOrderDirection2entgoᚗioᚋcontribᚋentgqlᚐOrderDirection(ctx, v)
			if err != nil {
				return it, err
			}
			it.Direction = data
		case "field":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("field"))
			data, err := ec.unmarshalOApiAccessOrderField2ᚖroofixᚋentᚐApiAccessOrderField(ctx, v)
			if err != nil {
				return it, err
			}
			it.Field = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var apiAccessImplementors = []string{"ApiAccess", "Node"}

func (ec *executionContext) _ApiAccess(ctx context.Context, sel ast.SelectionSet, obj *ent.ApiAccess) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, apiAccessImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ApiAccess")
		case "id":
			out.Values[i] = ec._ApiAccess_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "url":
			out.Values[i] = ec._ApiAccess_url(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "username":
			out.Values[i] = ec._ApiAccess_username(ctx, field, obj)
		case "password":
			out.Values[i] = ec._ApiAccess_password(ctx, field, obj)
		case "key":
			out.Values[i] = ec._ApiAccess_key(ctx, field, obj)
		case "secret":
			out.Values[i] = ec._ApiAccess_secret(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var apiAccessConnectionImplementors = []string{"ApiAccessConnection"}

func (ec *executionContext) _ApiAccessConnection(ctx context.Context, sel ast.SelectionSet, obj *ent.ApiAccessConnection) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, apiAccessConnectionImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ApiAccessConnection")
		case "edges":
			out.Values[i] = ec._ApiAccessConnection_edges(ctx, field, obj)
		case "pageInfo":
			out.Values[i] = ec._ApiAccessConnection_pageInfo(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "totalCount":
			out.Values[i] = ec._ApiAccessConnection_totalCount(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var apiAccessEdgeImplementors = []string{"ApiAccessEdge"}

func (ec *executionContext) _ApiAccessEdge(ctx context.Context, sel ast.SelectionSet, obj *ent.ApiAccessEdge) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, apiAccessEdgeImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ApiAccessEdge")
		case "node":
			out.Values[i] = ec._ApiAccessEdge_node(ctx, field, obj)
		case "cursor":
			out.Values[i] = ec._ApiAccessEdge_cursor(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNApiAccessEdge2ᚖroofixᚋentᚐApiAccessEdge(ctx context.Context, sel ast.SelectionSet, v *ent.ApiAccessEdge) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._ApiAccessEdge(ctx, sel, v)
}

func (ec *executionContext) unmarshalNApiAccessInput2roofixᚋpkgᚋapiaccessᚐInput(ctx context.Context, v interface{}) (apiaccess.Input, error) {
	res, err := ec.unmarshalInputApiAccessInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOApiAccess2ᚖroofixᚋentᚐApiAccess(ctx context.Context, sel ast.SelectionSet, v *ent.ApiAccess) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._ApiAccess(ctx, sel, v)
}

func (ec *executionContext) marshalOApiAccessConnection2ᚖroofixᚋentᚐApiAccessConnection(ctx context.Context, sel ast.SelectionSet, v *ent.ApiAccessConnection) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._ApiAccessConnection(ctx, sel, v)
}

func (ec *executionContext) marshalOApiAccessEdge2ᚕᚖroofixᚋentᚐApiAccessEdgeᚄ(ctx context.Context, sel ast.SelectionSet, v []*ent.ApiAccessEdge) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNApiAccessEdge2ᚖroofixᚋentᚐApiAccessEdge(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) unmarshalOApiAccessOrderField2ᚖroofixᚋentᚐApiAccessOrderField(ctx context.Context, v interface{}) (*ent.ApiAccessOrderField, error) {
	if v == nil {
		return nil, nil
	}
	var res = new(ent.ApiAccessOrderField)
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOApiAccessOrderField2ᚖroofixᚋentᚐApiAccessOrderField(ctx context.Context, sel ast.SelectionSet, v *ent.ApiAccessOrderField) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return v
}

// endregion ***************************** type.gotpl *****************************
