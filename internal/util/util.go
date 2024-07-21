package util

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

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
