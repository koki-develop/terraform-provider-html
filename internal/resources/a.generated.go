package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-html/internal/util"
)

var (
	_ resource.Resource = &resource_a{}
)

func newResource_a() resource.Resource {
	return &resource_a{}
}

type resource_a struct{}

func (r *resource_a) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_a"
}

func (r *resource_a) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The **`<a>`** [HTML](https://developer.mozilla.org/en-US/docs/Web/HTML) element (or _anchor_ element), with [its `href` attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a#href), creates a hyperlink to web pages, files, email addresses, locations in the same page, or anything else a URL can address.\n\nFor more information, see the [documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a).",
		Attributes: map[string]schema.Attribute{
			"children": schema.ListAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "The children of the element.",
				Optional:            true,
			},
			"attributionsrc": schema.DynamicAttribute{
				MarkdownDescription: "Specifies that you want the browser to send an [`Attribution-Reporting-Eligible`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Attribution-Reporting-Eligible) header.",
				Optional:            true,
			},
			"download": schema.DynamicAttribute{
				MarkdownDescription: "Causes the browser to treat the linked URL as a download.",
				Optional:            true,
			},
			"href": schema.DynamicAttribute{
				MarkdownDescription: "The URL that the hyperlink points to.",
				Optional:            true,
			},
			"hreflang": schema.DynamicAttribute{
				MarkdownDescription: "Hints at the human language of the linked URL.",
				Optional:            true,
			},
			"ping": schema.DynamicAttribute{
				MarkdownDescription: "A space-separated list of URLs.",
				Optional:            true,
			},
			"referrerpolicy": schema.DynamicAttribute{
				MarkdownDescription: "How much of the [referrer](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referer) to send when following the link.",
				Optional:            true,
			},
			"rel": schema.DynamicAttribute{
				MarkdownDescription: "The relationship of the linked URL as space-separated link types.",
				Optional:            true,
			},
			"target": schema.DynamicAttribute{
				MarkdownDescription: "Where to display the linked URL, as the name for a _browsing context_ (a tab, window, or [`<iframe>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe)).",
				Optional:            true,
			},
			"type": schema.DynamicAttribute{
				MarkdownDescription: "Hints at the linked URL's format with a [MIME type](https://developer.mozilla.org/en-US/docs/Glossary/MIME_type).",
				Optional:            true,
			},
			"accesskey": schema.DynamicAttribute{
				MarkdownDescription: "The **`accesskey`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) provides a hint for generating a keyboard shortcut for the current element.",
				Optional:            true,
			},

			"html": schema.StringAttribute{
				MarkdownDescription: "The html of the element.",
				Computed:            true,
			},
		},
	}
}

type resource_aModel struct {
	Children       types.List    `tfsdk:"children"`
	Attributionsrc types.Dynamic `tfsdk:"attributionsrc"`
	Download       types.Dynamic `tfsdk:"download"`
	Href           types.Dynamic `tfsdk:"href"`
	Hreflang       types.Dynamic `tfsdk:"hreflang"`
	Ping           types.Dynamic `tfsdk:"ping"`
	Referrerpolicy types.Dynamic `tfsdk:"referrerpolicy"`
	Rel            types.Dynamic `tfsdk:"rel"`
	Target         types.Dynamic `tfsdk:"target"`
	Type           types.Dynamic `tfsdk:"type"`
	Accesskey      types.Dynamic `tfsdk:"accesskey"`

	HTML types.String `tfsdk:"html"`
}

func (r *resource_a) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resource_a) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resource_a) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resource_a) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resource_a) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resource_aModel{},
		g,
		s,
		diags,
		func(m *resource_aModel) bool {
			html := new(strings.Builder)
			html.WriteString("<a")

			attrs := []string{}
			if !m.Attributionsrc.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "attributionsrc", m.Attributionsrc)
				if err != nil {
					diags.AddError("invalid attributionsrc attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Download.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "download", m.Download)
				if err != nil {
					diags.AddError("invalid download attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Href.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "href", m.Href)
				if err != nil {
					diags.AddError("invalid href attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Hreflang.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "hreflang", m.Hreflang)
				if err != nil {
					diags.AddError("invalid hreflang attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Ping.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "ping", m.Ping)
				if err != nil {
					diags.AddError("invalid ping attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Referrerpolicy.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "referrerpolicy", m.Referrerpolicy)
				if err != nil {
					diags.AddError("invalid referrerpolicy attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Rel.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "rel", m.Rel)
				if err != nil {
					diags.AddError("invalid rel attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Target.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "target", m.Target)
				if err != nil {
					diags.AddError("invalid target attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Type.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "type", m.Type)
				if err != nil {
					diags.AddError("invalid type attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}

			if len(attrs) > 0 {
				html.WriteString(" ")
				html.WriteString(strings.Join(attrs, " "))
			}

			if m.Children.IsNull() {
				if len(attrs) > 0 {
					html.WriteString(" ")
				}
				html.WriteString("/>")
			} else {
				html.WriteString(">")

				var children []string
				d := m.Children.ElementsAs(ctx, &children, true)
				diags.Append(d...)
				if diags.HasError() {
					return false
				}
				for _, child := range children {
					html.WriteString(child)
				}
				html.WriteString("</a>")
			}

			m.HTML = types.StringValue(html.String())
			return true
		},
	)
}

func init() {
	register(newResource_a)
}
