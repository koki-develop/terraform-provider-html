package util

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
)

var (
	patternJS   = regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$")
	patternJSON = regexp.MustCompile("[/+]json$")
	patternXML  = regexp.MustCompile("[/+]xml$")
)

func MinifyHTML(s string) (string, error) {
	m := minify.New()

	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFuncRegexp(patternJS, js.Minify)
	m.AddFuncRegexp(patternJSON, json.Minify)
	m.AddFuncRegexp(patternXML, xml.Minify)

	return m.String("text/html", s)
}

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
