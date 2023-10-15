// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"roofix/pkg/enum"
	"roofix/server/model"
	"strconv"
	"sync"
	"sync/atomic"

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

func (ec *executionContext) _Options_type(ctx context.Context, field graphql.CollectedField, obj *model.Options) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Options_type(ctx, field)
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
		return obj.Type, nil
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
	res := resTmp.(enum.OptionList)
	fc.Result = res
	return ec.marshalNOptionListType2roofixᚋpkgᚋenumᚐOptionList(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Options_type(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Options",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type OptionListType does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Options_options(ctx context.Context, field graphql.CollectedField, obj *model.Options) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Options_options(ctx, field)
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
		return obj.Options, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*model.Entity)
	fc.Result = res
	return ec.marshalOEntity2ᚕᚖroofixᚋserverᚋmodelᚐEntityᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Options_options(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Options",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Entity_id(ctx, field)
			case "name":
				return ec.fieldContext_Entity_name(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Entity", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var optionsImplementors = []string{"Options"}

func (ec *executionContext) _Options(ctx context.Context, sel ast.SelectionSet, obj *model.Options) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, optionsImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Options")
		case "type":
			out.Values[i] = ec._Options_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "options":
			out.Values[i] = ec._Options_options(ctx, field, obj)
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

func (ec *executionContext) unmarshalNOptionListType2roofixᚋpkgᚋenumᚐOptionList(ctx context.Context, v interface{}) (enum.OptionList, error) {
	var res enum.OptionList
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNOptionListType2roofixᚋpkgᚋenumᚐOptionList(ctx context.Context, sel ast.SelectionSet, v enum.OptionList) graphql.Marshaler {
	return v
}

func (ec *executionContext) marshalNOptions2ᚖroofixᚋserverᚋmodelᚐOptions(ctx context.Context, sel ast.SelectionSet, v *model.Options) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Options(ctx, sel, v)
}

func (ec *executionContext) unmarshalOOptionListType2ᚕroofixᚋpkgᚋenumᚐOptionListᚄ(ctx context.Context, v interface{}) ([]enum.OptionList, error) {
	if v == nil {
		return nil, nil
	}
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]enum.OptionList, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNOptionListType2roofixᚋpkgᚋenumᚐOptionList(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalOOptionListType2ᚕroofixᚋpkgᚋenumᚐOptionListᚄ(ctx context.Context, sel ast.SelectionSet, v []enum.OptionList) graphql.Marshaler {
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
			ret[i] = ec.marshalNOptionListType2roofixᚋpkgᚋenumᚐOptionList(ctx, sel, v[i])
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

func (ec *executionContext) unmarshalOOptionListType2ᚖroofixᚋpkgᚋenumᚐOptionList(ctx context.Context, v interface{}) (*enum.OptionList, error) {
	if v == nil {
		return nil, nil
	}
	var res = new(enum.OptionList)
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOOptionListType2ᚖroofixᚋpkgᚋenumᚐOptionList(ctx context.Context, sel ast.SelectionSet, v *enum.OptionList) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return v
}

func (ec *executionContext) marshalOOptions2ᚕᚖroofixᚋserverᚋmodelᚐOptions(ctx context.Context, sel ast.SelectionSet, v []*model.Options) graphql.Marshaler {
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
			ret[i] = ec.marshalOOptions2ᚖroofixᚋserverᚋmodelᚐOptions(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	return ret
}

func (ec *executionContext) marshalOOptions2ᚕᚖroofixᚋserverᚋmodelᚐOptionsᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.Options) graphql.Marshaler {
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
			ret[i] = ec.marshalNOptions2ᚖroofixᚋserverᚋmodelᚐOptions(ctx, sel, v[i])
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

func (ec *executionContext) marshalOOptions2ᚖroofixᚋserverᚋmodelᚐOptions(ctx context.Context, sel ast.SelectionSet, v *model.Options) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Options(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************