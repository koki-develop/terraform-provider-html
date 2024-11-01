package datasources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/koki-develop/terraform-provider-html/internal/util"
)

var (
	_ datasource.DataSource = &datasource_button{}
)

func newDatasource_button() datasource.DataSource {
	return &datasource_button{}
}

type datasource_button struct{}

func (d *datasource_button) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_button"
}

func (d *datasource_button) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "**`<button>`** [HTML](https://developer.mozilla.org/en-US/docs/Web/HTML) element is an interactive element activated by a user with a mouse, keyboard, finger, voice command, or other assistive technology.\n\nFor more information, see the [documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button).",
		Attributes: map[string]schema.Attribute{
			"children": schema.ListAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "The children of the element.",
				Optional:            true,
			},
			"autofocus": schema.DynamicAttribute{
				MarkdownDescription: "This Boolean attribute specifies that the button should have input [focus](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/focus) when the page loads.",
				Optional:            true,
			},
			"disabled": schema.DynamicAttribute{
				MarkdownDescription: "This Boolean attribute prevents the user from interacting with the button: it cannot be pressed or focused.",
				Optional:            true,
			},
			"form": schema.DynamicAttribute{
				MarkdownDescription: "The [`<form>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form) element to associate the button with (its _form owner_).",
				Optional:            true,
			},
			"formaction": schema.DynamicAttribute{
				MarkdownDescription: "The URL that processes the information submitted by the button.",
				Optional:            true,
			},
			"formenctype": schema.DynamicAttribute{
				MarkdownDescription: "If the button is a submit button (it's inside/associated with a `<form>` and doesn't have `type=\"button\"`), specifies how to encode the form data that is submitted.",
				Optional:            true,
			},
			"formmethod": schema.DynamicAttribute{
				MarkdownDescription: "If the button is a submit button (it's inside/associated with a `<form>` and doesn't have `type=\"button\"`), this attribute specifies the [HTTP method](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods) used to submit the form.",
				Optional:            true,
			},
			"formnovalidate": schema.DynamicAttribute{
				MarkdownDescription: "If the button is a submit button, this Boolean attribute specifies that the form is not to be [validated](https://developer.mozilla.org/en-US/docs/Learn/Forms/Form_validation) when it is submitted.",
				Optional:            true,
			},
			"formtarget": schema.DynamicAttribute{
				MarkdownDescription: "If the button is a submit button, this attribute is an author-defined name or standardized, underscore-prefixed keyword indicating where to display the response from submitting the form.",
				Optional:            true,
			},
			"name": schema.DynamicAttribute{
				MarkdownDescription: "The name of the button, submitted as a pair with the button's `value` as part of the form data, when that button is used to submit the form.",
				Optional:            true,
			},
			"popovertarget": schema.DynamicAttribute{
				MarkdownDescription: "Turns a `<button>` element into a popover control button; takes the ID of the popover element to control as its value.",
				Optional:            true,
			},
			"popovertargetaction": schema.DynamicAttribute{
				MarkdownDescription: "Specifies the action to be performed on a popover element being controlled by a control `<button>`.",
				Optional:            true,
			},
			"type": schema.DynamicAttribute{
				MarkdownDescription: "The default behavior of the button.",
				Optional:            true,
			},
			"value": schema.DynamicAttribute{
				MarkdownDescription: "Defines the value associated with the button's name when it's submitted with the form data.",
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

type datasource_buttonModel struct {
	Children            types.List    `tfsdk:"children"`
	Autofocus           types.Dynamic `tfsdk:"autofocus"`
	Disabled            types.Dynamic `tfsdk:"disabled"`
	Form                types.Dynamic `tfsdk:"form"`
	Formaction          types.Dynamic `tfsdk:"formaction"`
	Formenctype         types.Dynamic `tfsdk:"formenctype"`
	Formmethod          types.Dynamic `tfsdk:"formmethod"`
	Formnovalidate      types.Dynamic `tfsdk:"formnovalidate"`
	Formtarget          types.Dynamic `tfsdk:"formtarget"`
	Name                types.Dynamic `tfsdk:"name"`
	Popovertarget       types.Dynamic `tfsdk:"popovertarget"`
	Popovertargetaction types.Dynamic `tfsdk:"popovertargetaction"`
	Type                types.Dynamic `tfsdk:"type"`
	Value               types.Dynamic `tfsdk:"value"`
	Accesskey           types.Dynamic `tfsdk:"accesskey"`
	Autocapitalize      types.Dynamic `tfsdk:"autocapitalize"`
	Class               types.Dynamic `tfsdk:"class"`
	Contenteditable     types.Dynamic `tfsdk:"contenteditable"`
	Dir                 types.Dynamic `tfsdk:"dir"`
	Draggable           types.Dynamic `tfsdk:"draggable"`
	Enterkeyhint        types.Dynamic `tfsdk:"enterkeyhint"`
	Exportparts         types.Dynamic `tfsdk:"exportparts"`
	Hidden              types.Dynamic `tfsdk:"hidden"`
	Id                  types.Dynamic `tfsdk:"id"`
	Inert               types.Dynamic `tfsdk:"inert"`
	Inputmode           types.Dynamic `tfsdk:"inputmode"`
	Is                  types.Dynamic `tfsdk:"is"`
	Itemid              types.Dynamic `tfsdk:"itemid"`
	Itemprop            types.Dynamic `tfsdk:"itemprop"`
	Itemref             types.Dynamic `tfsdk:"itemref"`
	Itemscope           types.Dynamic `tfsdk:"itemscope"`
	Itemtype            types.Dynamic `tfsdk:"itemtype"`
	Lang                types.Dynamic `tfsdk:"lang"`
	Nonce               types.Dynamic `tfsdk:"nonce"`
	Part                types.Dynamic `tfsdk:"part"`
	Popover             types.Dynamic `tfsdk:"popover"`
	Slot                types.Dynamic `tfsdk:"slot"`
	Spellcheck          types.Dynamic `tfsdk:"spellcheck"`
	Style               types.Dynamic `tfsdk:"style"`
	Tabindex            types.Dynamic `tfsdk:"tabindex"`
	Title               types.Dynamic `tfsdk:"title"`
	Translate           types.Dynamic `tfsdk:"translate"`

	HTML types.String `tfsdk:"html"`
}

func (d *datasource_button) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.handleRequest(ctx, &req.Config, &resp.State, &resp.Diagnostics)
}

func (d *datasource_button) handleRequest(ctx context.Context, g util.ModelGetter, s util.ModelSetter, diags *diag.Diagnostics) {
	util.HandleRequest(
		ctx,
		&datasource_buttonModel{},
		g,
		s,
		diags,
		func(m *datasource_buttonModel) bool {
			html := new(strings.Builder)
			html.WriteString("<button")

			attrs := []string{}
			if !m.Autofocus.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "autofocus", m.Autofocus)
				if err != nil {
					diags.AddError("invalid autofocus attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Disabled.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "disabled", m.Disabled)
				if err != nil {
					diags.AddError("invalid disabled attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Form.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "form", m.Form)
				if err != nil {
					diags.AddError("invalid form attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Formaction.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "formaction", m.Formaction)
				if err != nil {
					diags.AddError("invalid formaction attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Formenctype.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "formenctype", m.Formenctype)
				if err != nil {
					diags.AddError("invalid formenctype attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Formmethod.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "formmethod", m.Formmethod)
				if err != nil {
					diags.AddError("invalid formmethod attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Formnovalidate.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "formnovalidate", m.Formnovalidate)
				if err != nil {
					diags.AddError("invalid formnovalidate attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Formtarget.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "formtarget", m.Formtarget)
				if err != nil {
					diags.AddError("invalid formtarget attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Name.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "name", m.Name)
				if err != nil {
					diags.AddError("invalid name attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Popovertarget.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "popovertarget", m.Popovertarget)
				if err != nil {
					diags.AddError("invalid popovertarget attribute", err.Error())
					return false
				}
				attrs = append(attrs, attr)
			}
			if !m.Popovertargetaction.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "popovertargetaction", m.Popovertargetaction)
				if err != nil {
					diags.AddError("invalid popovertargetaction attribute", err.Error())
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
			if !m.Value.IsNull() {
				attr, err := util.StringifyAttribute(ctx, "value", m.Value)
				if err != nil {
					diags.AddError("invalid value attribute", err.Error())
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
				html.WriteString("</button>")
			}

			m.HTML = types.StringValue(html.String())
			return true
		},
	)
}

func init() {
	register(newDatasource_button)
}
