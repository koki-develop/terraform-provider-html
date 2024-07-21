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
	_ resource.Resource = &resource_img{}
)

func newResource_img() resource.Resource {
	return &resource_img{}
}

type resource_img struct{}

func (r *resource_img) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_img"
}

func (r *resource_img) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The **`<img>`** [HTML](https://developer.mozilla.org/en-US/docs/Web/HTML) element embeds an image into the document.\n\nFor more information, see the [documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img).",
		Attributes: map[string]schema.Attribute{
			"children": schema.ListAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "The children of the element.",
				Optional:            true,
			},
			"alt": schema.DynamicAttribute{
				MarkdownDescription: "Defines text that can replace the image in the page.",
				Optional:            true,
			},
			"crossorigin": schema.DynamicAttribute{
				MarkdownDescription: "Indicates if the fetching of the image must be done using a [CORS](https://developer.mozilla.org/en-US/docs/Glossary/CORS) request.",
				Optional:            true,
			},
			"decoding": schema.DynamicAttribute{
				MarkdownDescription: "This attribute provides a hint to the browser as to whether it should perform image decoding along with rendering the other DOM content in a single presentation step that looks more \"correct\" (`sync`), or render and present the other DOM content first and then decode the image and present it later (`async`).",
				Optional:            true,
			},
			"elementtiming": schema.DynamicAttribute{
				MarkdownDescription: "Marks the image for observation by the [`PerformanceElementTiming`](https://developer.mozilla.org/en-US/docs/Web/API/PerformanceElementTiming) API.",
				Optional:            true,
			},
			"fetchpriority": schema.DynamicAttribute{
				MarkdownDescription: "Provides a hint of the relative priority to use when fetching the image.",
				Optional:            true,
			},
			"height": schema.DynamicAttribute{
				MarkdownDescription: "The intrinsic height of the image, in pixels.",
				Optional:            true,
			},
			"ismap": schema.DynamicAttribute{
				MarkdownDescription: "This Boolean attribute indicates that the image is part of a [server-side map](https://en.wikipedia.org/wiki/Image_map#Server-side).",
				Optional:            true,
			},
			"loading": schema.DynamicAttribute{
				MarkdownDescription: "Indicates how the browser should load the image.",
				Optional:            true,
			},
			"referrerpolicy": schema.DynamicAttribute{
				MarkdownDescription: "A string indicating which referrer to use when fetching the resource.",
				Optional:            true,
			},
			"sizes": schema.DynamicAttribute{
				MarkdownDescription: "One or more strings separated by commas, indicating a set of source sizes.",
				Optional:            true,
			},
			"src": schema.DynamicAttribute{
				MarkdownDescription: "The image [URL](https://developer.mozilla.org/en-US/docs/Glossary/URL).",
				Optional:            true,
			},
			"srcset": schema.DynamicAttribute{
				MarkdownDescription: "One or more strings separated by commas, indicating possible image sources for the [user agent](https://developer.mozilla.org/en-US/docs/Glossary/User_agent) to use.",
				Optional:            true,
			},
			"width": schema.DynamicAttribute{
				MarkdownDescription: "The intrinsic width of the image in pixels. Must be an integer without a unit.",
				Optional:            true,
			},
			"usemap": schema.DynamicAttribute{
				MarkdownDescription: "The partial [URL](https://developer.mozilla.org/en-US/docs/Glossary/URL) (starting with `#`) of an [image map](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map) associated with the element.",
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

type resource_imgModel struct {
	Children        types.List    `tfsdk:"children"`
	Alt             types.Dynamic `tfsdk:"alt"`
	Crossorigin     types.Dynamic `tfsdk:"crossorigin"`
	Decoding        types.Dynamic `tfsdk:"decoding"`
	Elementtiming   types.Dynamic `tfsdk:"elementtiming"`
	Fetchpriority   types.Dynamic `tfsdk:"fetchpriority"`
	Height          types.Dynamic `tfsdk:"height"`
	Ismap           types.Dynamic `tfsdk:"ismap"`
	Loading         types.Dynamic `tfsdk:"loading"`
	Referrerpolicy  types.Dynamic `tfsdk:"referrerpolicy"`
	Sizes           types.Dynamic `tfsdk:"sizes"`
	Src             types.Dynamic `tfsdk:"src"`
	Srcset          types.Dynamic `tfsdk:"srcset"`
	Width           types.Dynamic `tfsdk:"width"`
	Usemap          types.Dynamic `tfsdk:"usemap"`
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

func (r *resource_img) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resource_img) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resource_img) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resource_img) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resource_img) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resource_imgModel{},
		g,
		s,
		diags,
		func(m *resource_imgModel) bool {
			html := new(strings.Builder)
			html.WriteString("<img")

			attrs := []string{}
			if !m.Alt.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "alt", m.Alt)
				if err != nil {
					diags.AddError("invalid alt attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Crossorigin.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "crossorigin", m.Crossorigin)
				if err != nil {
					diags.AddError("invalid crossorigin attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Decoding.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "decoding", m.Decoding)
				if err != nil {
					diags.AddError("invalid decoding attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Elementtiming.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "elementtiming", m.Elementtiming)
				if err != nil {
					diags.AddError("invalid elementtiming attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Fetchpriority.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "fetchpriority", m.Fetchpriority)
				if err != nil {
					diags.AddError("invalid fetchpriority attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Height.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "height", m.Height)
				if err != nil {
					diags.AddError("invalid height attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Ismap.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "ismap", m.Ismap)
				if err != nil {
					diags.AddError("invalid ismap attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Loading.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "loading", m.Loading)
				if err != nil {
					diags.AddError("invalid loading attribute", err.Error())
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
			if !m.Sizes.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "sizes", m.Sizes)
				if err != nil {
					diags.AddError("invalid sizes attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Src.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "src", m.Src)
				if err != nil {
					diags.AddError("invalid src attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Srcset.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "srcset", m.Srcset)
				if err != nil {
					diags.AddError("invalid srcset attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Width.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "width", m.Width)
				if err != nil {
					diags.AddError("invalid width attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Usemap.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "usemap", m.Usemap)
				if err != nil {
					diags.AddError("invalid usemap attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Accesskey.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "accesskey", m.Accesskey)
				if err != nil {
					diags.AddError("invalid accesskey attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Autocapitalize.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "autocapitalize", m.Autocapitalize)
				if err != nil {
					diags.AddError("invalid autocapitalize attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Autofocus.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "autofocus", m.Autofocus)
				if err != nil {
					diags.AddError("invalid autofocus attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Class.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "class", m.Class)
				if err != nil {
					diags.AddError("invalid class attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Contenteditable.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "contenteditable", m.Contenteditable)
				if err != nil {
					diags.AddError("invalid contenteditable attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Dir.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "dir", m.Dir)
				if err != nil {
					diags.AddError("invalid dir attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Draggable.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "draggable", m.Draggable)
				if err != nil {
					diags.AddError("invalid draggable attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Enterkeyhint.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "enterkeyhint", m.Enterkeyhint)
				if err != nil {
					diags.AddError("invalid enterkeyhint attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Exportparts.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "exportparts", m.Exportparts)
				if err != nil {
					diags.AddError("invalid exportparts attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Hidden.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "hidden", m.Hidden)
				if err != nil {
					diags.AddError("invalid hidden attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Id.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "id", m.Id)
				if err != nil {
					diags.AddError("invalid id attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Inert.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "inert", m.Inert)
				if err != nil {
					diags.AddError("invalid inert attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Inputmode.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "inputmode", m.Inputmode)
				if err != nil {
					diags.AddError("invalid inputmode attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Is.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "is", m.Is)
				if err != nil {
					diags.AddError("invalid is attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Itemid.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "itemid", m.Itemid)
				if err != nil {
					diags.AddError("invalid itemid attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Itemprop.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "itemprop", m.Itemprop)
				if err != nil {
					diags.AddError("invalid itemprop attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Itemref.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "itemref", m.Itemref)
				if err != nil {
					diags.AddError("invalid itemref attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Itemscope.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "itemscope", m.Itemscope)
				if err != nil {
					diags.AddError("invalid itemscope attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Itemtype.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "itemtype", m.Itemtype)
				if err != nil {
					diags.AddError("invalid itemtype attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Lang.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "lang", m.Lang)
				if err != nil {
					diags.AddError("invalid lang attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Nonce.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "nonce", m.Nonce)
				if err != nil {
					diags.AddError("invalid nonce attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Part.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "part", m.Part)
				if err != nil {
					diags.AddError("invalid part attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Popover.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "popover", m.Popover)
				if err != nil {
					diags.AddError("invalid popover attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Slot.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "slot", m.Slot)
				if err != nil {
					diags.AddError("invalid slot attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Spellcheck.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "spellcheck", m.Spellcheck)
				if err != nil {
					diags.AddError("invalid spellcheck attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Style.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "style", m.Style)
				if err != nil {
					diags.AddError("invalid style attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Tabindex.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "tabindex", m.Tabindex)
				if err != nil {
					diags.AddError("invalid tabindex attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Title.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "title", m.Title)
				if err != nil {
					diags.AddError("invalid title attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Translate.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "translate", m.Translate)
				if err != nil {
					diags.AddError("invalid translate attribute", err.Error())
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
				html.WriteString("</img>")
			}

			m.HTML = types.StringValue(html.String())
			return true
		},
	)
}

func init() {
	register(newResource_img)
}
