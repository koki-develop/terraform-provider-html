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
				MarkdownDescription: "The **`accesskey`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) provides a hint for generating a keyboard shortcut for the current element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/accesskey))",
				Optional:            true,
			},
			"autocapitalize": schema.DynamicAttribute{
				MarkdownDescription: "The **`autocapitalize`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute that controls whether inputted text is automatically capitalized and, if so, in what manner. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/autocapitalize))",
				Optional:            true,
			},
			"autofocus": schema.DynamicAttribute{
				MarkdownDescription: "The **`autofocus`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is a Boolean attribute indicating that an element should be focused on page load, or when the [`<dialog>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog) that it is part of is displayed. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/autofocus))",
				Optional:            true,
			},
			"class": schema.DynamicAttribute{
				MarkdownDescription: "The **`class`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is a space-separated list of the case-sensitive classes of the element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/class))",
				Optional:            true,
			},
			"contenteditable": schema.DynamicAttribute{
				MarkdownDescription: "The **`contenteditable`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an enumerated attribute indicating if the element should be editable by the user. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/contenteditable))",
				Optional:            true,
			},
			"dir": schema.DynamicAttribute{
				MarkdownDescription: "The **`dir`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute that indicates the directionality of the element's text. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/dir))",
				Optional:            true,
			},
			"draggable": schema.DynamicAttribute{
				MarkdownDescription: "The **`draggable`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute that indicates whether the element can be dragged, either with native browser behavior or the [HTML Drag and Drop API](https://developer.mozilla.org/en-US/docs/Web/API/HTML_Drag_and_Drop_API). ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/draggable))",
				Optional:            true,
			},
			"enterkeyhint": schema.DynamicAttribute{
				MarkdownDescription: "The **`enterkeyhint`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute defining what action label (or icon) to present for the enter key on virtual keyboards. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/enterkeyhint))",
				Optional:            true,
			},
			"exportparts": schema.DynamicAttribute{
				MarkdownDescription: "The **`exportparts`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) allows you to select and style elements existing in nested [shadow trees](https://developer.mozilla.org/en-US/docs/Glossary/Shadow_tree), by exporting their `part` names. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/exportparts))",
				Optional:            true,
			},
			"hidden": schema.DynamicAttribute{
				MarkdownDescription: "The **`hidden`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute indicating that the browser should not render the contents of the element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/hidden))",
				Optional:            true,
			},
			"id": schema.DynamicAttribute{
				MarkdownDescription: "The **`id`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) defines an identifier (ID) which must be unique in the whole document. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/id))",
				Optional:            true,
			},
			"inert": schema.DynamicAttribute{
				MarkdownDescription: "The **`inert`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is a Boolean attribute indicating that the browser will ignore the element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/inert))",
				Optional:            true,
			},
			"inputmode": schema.DynamicAttribute{
				MarkdownDescription: "The **`inputmode`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute that hints at the type of data that might be entered by the user while editing the element or its contents. This allows a browser to display an appropriate virtual keyboard. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/inputmode))",
				Optional:            true,
			},
			"is": schema.DynamicAttribute{
				MarkdownDescription: "The **`is`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) allows you to specify that a standard HTML element should behave like a defined custom built-in element (see [Using custom elements](https://developer.mozilla.org/en-US/docs/Web/API/Web_components/Using_custom_elements) for more details). ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/is))",
				Optional:            true,
			},
			"itemid": schema.DynamicAttribute{
				MarkdownDescription: "The **`itemid`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) provides microdata in the form of a unique, global identifier of an item. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemid))",
				Optional:            true,
			},
			"itemprop": schema.DynamicAttribute{
				MarkdownDescription: "The **`itemprop`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is used to add properties to an item. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemprop))",
				Optional:            true,
			},
			"itemref": schema.DynamicAttribute{
				MarkdownDescription: "Properties that are not descendants of an element with the [`itemscope`](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemscope) attribute can be associated with an item using the [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) **`itemref`**. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemref))",
				Optional:            true,
			},
			"itemscope": schema.DynamicAttribute{
				MarkdownDescription: "**`itemscope`** is a boolean [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) that defines the scope of associated metadata. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemscope))",
				Optional:            true,
			},
			"itemtype": schema.DynamicAttribute{
				MarkdownDescription: "The [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) **`itemtype`** specifies the URL of the vocabulary that will be used to define `itemprop`'s (item properties) in the data structure. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemtype))",
				Optional:            true,
			},
			"lang": schema.DynamicAttribute{
				MarkdownDescription: "The **`lang`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) helps define the language of an element: the language that non-editable elements are written in, or the language that the editable elements should be written in by the user. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/lang))",
				Optional:            true,
			},
			"nonce": schema.DynamicAttribute{
				MarkdownDescription: "The **`nonce`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is a content attribute defining a cryptographic nonce (\"number used once\") which can be used by [Content Security Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP) to determine whether or not a given fetch will be allowed to proceed for a given element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/nonce))",
				Optional:            true,
			},
			"part": schema.DynamicAttribute{
				MarkdownDescription: "The **`part`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) contains a space-separated list of the part names of the element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/part))",
				Optional:            true,
			},
			"popover": schema.DynamicAttribute{
				MarkdownDescription: "The **`popover`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is used to designate an element as a popover element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/popover))",
				Optional:            true,
			},
			"slot": schema.DynamicAttribute{
				MarkdownDescription: "The **`slot`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) assigns a slot in a [shadow DOM](https://developer.mozilla.org/en-US/docs/Web/API/Web_components/Using_shadow_DOM) shadow tree to an element: An element with a `slot` attribute is assigned to the slot created by the [`<slot>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/slot) element whose [`name`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/slot#name) attribute's value matches that `slot` attribute's value. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/slot))",
				Optional:            true,
			},
			"spellcheck": schema.DynamicAttribute{
				MarkdownDescription: "The **`spellcheck`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute that defines whether the element may be checked for spelling errors. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/spellcheck))",
				Optional:            true,
			},
			"style": schema.DynamicAttribute{
				MarkdownDescription: "The **`style`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) contains [CSS](https://developer.mozilla.org/en-US/docs/Web/CSS) styling declarations to be applied to the element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/style))",
				Optional:            true,
			},
			"tabindex": schema.DynamicAttribute{
				MarkdownDescription: "The **`tabindex`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) allows developers to make HTML elements focusable, allow or prevent them from being sequentially focusable (usually with the <kbd>Tab</kbd> key, hence the name) and determine their relative ordering for sequential focus navigation. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/tabindex))",
				Optional:            true,
			},
			"title": schema.DynamicAttribute{
				MarkdownDescription: "The **`title`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) contains text representing advisory information related to the element it belongs to. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/title))",
				Optional:            true,
			},
			"translate": schema.DynamicAttribute{
				MarkdownDescription: "The **`translate`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute that is used to specify whether an element's _translatable attribute_ values and its [`Text`](https://developer.mozilla.org/en-US/docs/Web/API/Text) node children should be translated when the page is localized, or whether to leave them unchanged. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/translate))",
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
	Children        types.List    `tfsdk:"children"`
	Attributionsrc  types.Dynamic `tfsdk:"attributionsrc"`
	Download        types.Dynamic `tfsdk:"download"`
	Href            types.Dynamic `tfsdk:"href"`
	Hreflang        types.Dynamic `tfsdk:"hreflang"`
	Ping            types.Dynamic `tfsdk:"ping"`
	Referrerpolicy  types.Dynamic `tfsdk:"referrerpolicy"`
	Rel             types.Dynamic `tfsdk:"rel"`
	Target          types.Dynamic `tfsdk:"target"`
	Type            types.Dynamic `tfsdk:"type"`
	Accesskey       types.Dynamic `tfsdk:"accesskey"`
	Autocapitalize  types.Dynamic `tfsdk:"autocapitalize"`
	Autofocus       types.Dynamic `tfsdk:"autofocus"`
	Class           types.Dynamic `tfsdk:"class"`
	Contenteditable types.Dynamic `tfsdk:"contenteditable"`
	Dir             types.Dynamic `tfsdk:"dir"`
	Draggable       types.Dynamic `tfsdk:"draggable"`
	Enterkeyhint    types.Dynamic `tfsdk:"enterkeyhint"`
	Exportparts     types.Dynamic `tfsdk:"exportparts"`
	Hidden          types.Dynamic `tfsdk:"hidden"`
	Id              types.Dynamic `tfsdk:"id"`
	Inert           types.Dynamic `tfsdk:"inert"`
	Inputmode       types.Dynamic `tfsdk:"inputmode"`
	Is              types.Dynamic `tfsdk:"is"`
	Itemid          types.Dynamic `tfsdk:"itemid"`
	Itemprop        types.Dynamic `tfsdk:"itemprop"`
	Itemref         types.Dynamic `tfsdk:"itemref"`
	Itemscope       types.Dynamic `tfsdk:"itemscope"`
	Itemtype        types.Dynamic `tfsdk:"itemtype"`
	Lang            types.Dynamic `tfsdk:"lang"`
	Nonce           types.Dynamic `tfsdk:"nonce"`
	Part            types.Dynamic `tfsdk:"part"`
	Popover         types.Dynamic `tfsdk:"popover"`
	Slot            types.Dynamic `tfsdk:"slot"`
	Spellcheck      types.Dynamic `tfsdk:"spellcheck"`
	Style           types.Dynamic `tfsdk:"style"`
	Tabindex        types.Dynamic `tfsdk:"tabindex"`
	Title           types.Dynamic `tfsdk:"title"`
	Translate       types.Dynamic `tfsdk:"translate"`

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
