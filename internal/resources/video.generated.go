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
	_ resource.Resource = &resource_video{}
)

func newResource_video() resource.Resource {
	return &resource_video{}
}

type resource_video struct{}

func (r *resource_video) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_video"
}

func (r *resource_video) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The **`<video>`** [HTML](https://developer.mozilla.org/en-US/docs/Web/HTML) element embeds a media player which supports video playback into the document.\n\nFor more information, see the [documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video).",
		Attributes: map[string]schema.Attribute{
			"children": schema.ListAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "The children of the element.",
				Optional:            true,
			},
			"autoplay": schema.DynamicAttribute{
				MarkdownDescription: "A Boolean attribute; if specified, the video automatically begins to play back as soon as it can without stopping to finish loading the data.",
				Optional:            true,
			},
			"controls": schema.DynamicAttribute{
				MarkdownDescription: "If this attribute is present, the browser will offer controls to allow the user to control video playback, including volume, seeking, and pause/resume playback.",
				Optional:            true,
			},
			"controlslist": schema.DynamicAttribute{
				MarkdownDescription: "The [`controlslist`](https://wicg.github.io/controls-list/explainer.html) attribute, when specified, helps the browser select what controls to show for the `video` element whenever the browser shows its own set of controls (that is, when the `controls` attribute is specified).",
				Optional:            true,
			},
			"crossorigin": schema.DynamicAttribute{
				MarkdownDescription: "This [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute indicates whether to use CORS to fetch the related video.",
				Optional:            true,
			},
			"disablepictureinpicture": schema.DynamicAttribute{
				MarkdownDescription: "Prevents the browser from suggesting a Picture-in-Picture context menu or to request Picture-in-Picture automatically in some cases.",
				Optional:            true,
			},
			"disableremoteplayback": schema.DynamicAttribute{
				MarkdownDescription: "A Boolean attribute used to disable the capability of remote playback in devices that are attached using wired (HDMI, DVI, etc.) and wireless technologies (Miracast, Chromecast, DLNA, AirPlay, etc.).",
				Optional:            true,
			},
			"height": schema.DynamicAttribute{
				MarkdownDescription: "The height of the video's display area, in [CSS pixels](https://drafts.csswg.org/css-values/#px) (absolute values only; [no percentages](https://html.spec.whatwg.org/multipage/embedded-content.html#dimension-attributes)).",
				Optional:            true,
			},
			"loop": schema.DynamicAttribute{
				MarkdownDescription: "A Boolean attribute; if specified, the browser will automatically seek back to the start upon reaching the end of the video.",
				Optional:            true,
			},
			"muted": schema.DynamicAttribute{
				MarkdownDescription: "A Boolean attribute that indicates the default audio mute setting contained in the video.",
				Optional:            true,
			},
			"playsinline": schema.DynamicAttribute{
				MarkdownDescription: "A Boolean attribute indicating that the video is to be played \"inline\", that is, within the element's playback area.",
				Optional:            true,
			},
			"poster": schema.DynamicAttribute{
				MarkdownDescription: "A URL for an image to be shown while the video is downloading.",
				Optional:            true,
			},
			"preload": schema.DynamicAttribute{
				MarkdownDescription: "This [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute is intended to provide a hint to the browser about what the author thinks will lead to the best user experience regarding what content is loaded before the video is played.",
				Optional:            true,
			},
			"src": schema.DynamicAttribute{
				MarkdownDescription: "The URL of the video to embed.",
				Optional:            true,
			},
			"width": schema.DynamicAttribute{
				MarkdownDescription: "The width of the video's display area, in [CSS pixels](https://drafts.csswg.org/css-values/#px) (absolute values only; [no percentages](https://html.spec.whatwg.org/multipage/embedded-content.html#dimension-attributes)).",
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

type resource_videoModel struct {
	Children                types.List    `tfsdk:"children"`
	Autoplay                types.Dynamic `tfsdk:"autoplay"`
	Controls                types.Dynamic `tfsdk:"controls"`
	Controlslist            types.Dynamic `tfsdk:"controlslist"`
	Crossorigin             types.Dynamic `tfsdk:"crossorigin"`
	Disablepictureinpicture types.Dynamic `tfsdk:"disablepictureinpicture"`
	Disableremoteplayback   types.Dynamic `tfsdk:"disableremoteplayback"`
	Height                  types.Dynamic `tfsdk:"height"`
	Loop                    types.Dynamic `tfsdk:"loop"`
	Muted                   types.Dynamic `tfsdk:"muted"`
	Playsinline             types.Dynamic `tfsdk:"playsinline"`
	Poster                  types.Dynamic `tfsdk:"poster"`
	Preload                 types.Dynamic `tfsdk:"preload"`
	Src                     types.Dynamic `tfsdk:"src"`
	Width                   types.Dynamic `tfsdk:"width"`
	Accesskey               types.Dynamic `tfsdk:"accesskey"`
	Autocapitalize          types.Dynamic `tfsdk:"autocapitalize"`
	Autofocus               types.Dynamic `tfsdk:"autofocus"`
	Class                   types.Dynamic `tfsdk:"class"`
	Contenteditable         types.Dynamic `tfsdk:"contenteditable"`
	Dir                     types.Dynamic `tfsdk:"dir"`
	Draggable               types.Dynamic `tfsdk:"draggable"`
	Enterkeyhint            types.Dynamic `tfsdk:"enterkeyhint"`
	Exportparts             types.Dynamic `tfsdk:"exportparts"`
	Hidden                  types.Dynamic `tfsdk:"hidden"`
	Id                      types.Dynamic `tfsdk:"id"`
	Inert                   types.Dynamic `tfsdk:"inert"`
	Inputmode               types.Dynamic `tfsdk:"inputmode"`
	Is                      types.Dynamic `tfsdk:"is"`
	Itemid                  types.Dynamic `tfsdk:"itemid"`
	Itemprop                types.Dynamic `tfsdk:"itemprop"`
	Itemref                 types.Dynamic `tfsdk:"itemref"`
	Itemscope               types.Dynamic `tfsdk:"itemscope"`
	Itemtype                types.Dynamic `tfsdk:"itemtype"`
	Lang                    types.Dynamic `tfsdk:"lang"`
	Nonce                   types.Dynamic `tfsdk:"nonce"`
	Part                    types.Dynamic `tfsdk:"part"`
	Popover                 types.Dynamic `tfsdk:"popover"`
	Slot                    types.Dynamic `tfsdk:"slot"`
	Spellcheck              types.Dynamic `tfsdk:"spellcheck"`
	Style                   types.Dynamic `tfsdk:"style"`
	Tabindex                types.Dynamic `tfsdk:"tabindex"`
	Title                   types.Dynamic `tfsdk:"title"`
	Translate               types.Dynamic `tfsdk:"translate"`

	HTML types.String `tfsdk:"html"`
}

func (r *resource_video) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resource_video) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.handleRequest(ctx, &req.State, &resp.State, &resp.Diagnostics)
}

func (r *resource_video) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.handleRequest(ctx, &req.Plan, &resp.State, &resp.Diagnostics)
}

func (r *resource_video) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *resource_video) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&resource_videoModel{},
		g,
		s,
		diags,
		func(m *resource_videoModel) bool {
			html := new(strings.Builder)
			html.WriteString("<video")

			attrs := []string{}
			if !m.Autoplay.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "autoplay", m.Autoplay)
				if err != nil {
					diags.AddError("invalid autoplay attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Controls.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "controls", m.Controls)
				if err != nil {
					diags.AddError("invalid controls attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Controlslist.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "controlslist", m.Controlslist)
				if err != nil {
					diags.AddError("invalid controlslist attribute", err.Error())
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
			if !m.Disablepictureinpicture.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "disablepictureinpicture", m.Disablepictureinpicture)
				if err != nil {
					diags.AddError("invalid disablepictureinpicture attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Disableremoteplayback.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "disableremoteplayback", m.Disableremoteplayback)
				if err != nil {
					diags.AddError("invalid disableremoteplayback attribute", err.Error())
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
			if !m.Loop.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "loop", m.Loop)
				if err != nil {
					diags.AddError("invalid loop attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Muted.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "muted", m.Muted)
				if err != nil {
					diags.AddError("invalid muted attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Playsinline.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "playsinline", m.Playsinline)
				if err != nil {
					diags.AddError("invalid playsinline attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Poster.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "poster", m.Poster)
				if err != nil {
					diags.AddError("invalid poster attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Preload.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "preload", m.Preload)
				if err != nil {
					diags.AddError("invalid preload attribute", err.Error())
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
			if !m.Width.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "width", m.Width)
				if err != nil {
					diags.AddError("invalid width attribute", err.Error())
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
				html.WriteString("</video>")
			}

			m.HTML = types.StringValue(html.String())
			return true
		},
	)
}

func init() {
	register(newResource_video)
}
