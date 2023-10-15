// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"roofix/pkg/model"
	model1 "roofix/server/model"
	"strconv"
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

func (ec *executionContext) _MobileAppSettings_logoURL(ctx context.Context, field graphql.CollectedField, obj *model.MobileAppSettings) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_MobileAppSettings_logoURL(ctx, field)
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
		return obj.LogoURL, nil
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

func (ec *executionContext) fieldContext_MobileAppSettings_logoURL(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "MobileAppSettings",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _MobileAppSettings_primaryColor(ctx context.Context, field graphql.CollectedField, obj *model.MobileAppSettings) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_MobileAppSettings_primaryColor(ctx, field)
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
		return obj.PrimaryColor, nil
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

func (ec *executionContext) fieldContext_MobileAppSettings_primaryColor(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "MobileAppSettings",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _MobileAppSettings_hideTabs(ctx context.Context, field graphql.CollectedField, obj *model.MobileAppSettings) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_MobileAppSettings_hideTabs(ctx, field)
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
		return obj.HideTabs, nil
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
	res := resTmp.([]string)
	fc.Result = res
	return ec.marshalNString2ᚕstringᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_MobileAppSettings_hideTabs(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "MobileAppSettings",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputInputMobileAppSettings(ctx context.Context, obj interface{}) (model1.InputMobileAppSettings, error) {
	var it model1.InputMobileAppSettings
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"logoURL", "primaryColor", "hideTabs"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "logoURL":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("logoURL"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.LogoURL = data
		case "primaryColor":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("primaryColor"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.PrimaryColor = data
		case "hideTabs":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("hideTabs"))
			data, err := ec.unmarshalOString2ᚕstringᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.HideTabs = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var mobileAppSettingsImplementors = []string{"MobileAppSettings"}

func (ec *executionContext) _MobileAppSettings(ctx context.Context, sel ast.SelectionSet, obj *model.MobileAppSettings) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mobileAppSettingsImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("MobileAppSettings")
		case "logoURL":
			out.Values[i] = ec._MobileAppSettings_logoURL(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "primaryColor":
			out.Values[i] = ec._MobileAppSettings_primaryColor(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "hideTabs":
			out.Values[i] = ec._MobileAppSettings_hideTabs(ctx, field, obj)
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

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNInputMobileAppSettings2roofixᚋserverᚋmodelᚐInputMobileAppSettings(ctx context.Context, v interface{}) (model1.InputMobileAppSettings, error) {
	res, err := ec.unmarshalInputInputMobileAppSettings(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNMobileAppSettings2roofixᚋpkgᚋmodelᚐMobileAppSettings(ctx context.Context, sel ast.SelectionSet, v model.MobileAppSettings) graphql.Marshaler {
	return ec._MobileAppSettings(ctx, sel, &v)
}

func (ec *executionContext) marshalNMobileAppSettings2ᚖroofixᚋpkgᚋmodelᚐMobileAppSettings(ctx context.Context, sel ast.SelectionSet, v *model.MobileAppSettings) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._MobileAppSettings(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************