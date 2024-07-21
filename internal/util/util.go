package util

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func StringifyAttribute(ctx context.Context, name string, value types.Dynamic) (string, error) {
	v := value.UnderlyingValue()
	switch v := v.(type) {
	case types.Number:
		return fmt.Sprintf("%s=\"%s\"", name, v.String()), nil
	case types.String:
		return fmt.Sprintf("%s=\"%s\"", name, v.ValueString()), nil
	case types.Bool:
		return name, nil
	default:
		return "", fmt.Errorf("attribute must be a string, number, or boolean. got %T", v)
	}
}

type ModelGetter interface {
	Get(ctx context.Context, target any) diag.Diagnostics
}

type ModelSetter interface {
	Set(ctx context.Context, val any) diag.Diagnostics
}

func HandleRequest[T any](ctx context.Context, model T, g ModelGetter, s ModelSetter, diags *diag.Diagnostics, h func(m T) bool) {
	diags.Append(g.Get(ctx, model)...)
	if diags.HasError() {
		return
	}

	if !h(model) {
		return
	}

	diags.Append(s.Set(ctx, model)...)
	if diags.HasError() {
		return
	}
}
