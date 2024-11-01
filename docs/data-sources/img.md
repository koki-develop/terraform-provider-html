---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "html_img Data Source - terraform-provider-html"
subcategory: ""
description: |-
  The <img> HTML https://developer.mozilla.org/en-US/docs/Web/HTML element embeds an image into the document.
  For more information, see the documentation https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img.
---

# html_img (Data Source)

The **`<img>`** [HTML](https://developer.mozilla.org/en-US/docs/Web/HTML) element embeds an image into the document.

For more information, see the [documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img).



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `accesskey` (Dynamic) The **`accesskey`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) provides a hint for generating a keyboard shortcut for the current element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/accesskey))
- `alt` (Dynamic) Defines text that can replace the image in the page.
- `autocapitalize` (Dynamic) The **`autocapitalize`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute that controls whether inputted text is automatically capitalized and, if so, in what manner. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/autocapitalize))
- `autofocus` (Dynamic) The **`autofocus`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is a Boolean attribute indicating that an element should be focused on page load, or when the [`<dialog>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog) that it is part of is displayed. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/autofocus))
- `children` (List of String) The children of the element.
- `class` (Dynamic) The **`class`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is a space-separated list of the case-sensitive classes of the element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/class))
- `contenteditable` (Dynamic) The **`contenteditable`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an enumerated attribute indicating if the element should be editable by the user. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/contenteditable))
- `crossorigin` (Dynamic) Indicates if the fetching of the image must be done using a [CORS](https://developer.mozilla.org/en-US/docs/Glossary/CORS) request.
- `decoding` (Dynamic) This attribute provides a hint to the browser as to whether it should perform image decoding along with rendering the other DOM content in a single presentation step that looks more "correct" (`sync`), or render and present the other DOM content first and then decode the image and present it later (`async`).
- `dir` (Dynamic) The **`dir`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute that indicates the directionality of the element's text. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/dir))
- `draggable` (Dynamic) The **`draggable`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute that indicates whether the element can be dragged, either with native browser behavior or the [HTML Drag and Drop API](https://developer.mozilla.org/en-US/docs/Web/API/HTML_Drag_and_Drop_API). ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/draggable))
- `elementtiming` (Dynamic) Marks the image for observation by the [`PerformanceElementTiming`](https://developer.mozilla.org/en-US/docs/Web/API/PerformanceElementTiming) API.
- `enterkeyhint` (Dynamic) The **`enterkeyhint`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute defining what action label (or icon) to present for the enter key on virtual keyboards. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/enterkeyhint))
- `exportparts` (Dynamic) The **`exportparts`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) allows you to select and style elements existing in nested [shadow trees](https://developer.mozilla.org/en-US/docs/Glossary/Shadow_tree), by exporting their `part` names. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/exportparts))
- `fetchpriority` (Dynamic) Provides a hint of the relative priority to use when fetching the image.
- `height` (Dynamic) The intrinsic height of the image, in pixels.
- `hidden` (Dynamic) The **`hidden`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute indicating that the browser should not render the contents of the element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/hidden))
- `id` (Dynamic) The **`id`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) defines an identifier (ID) which must be unique in the whole document. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/id))
- `inert` (Dynamic) The **`inert`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is a Boolean attribute indicating that the browser will ignore the element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/inert))
- `inputmode` (Dynamic) The **`inputmode`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute that hints at the type of data that might be entered by the user while editing the element or its contents. This allows a browser to display an appropriate virtual keyboard. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/inputmode))
- `is` (Dynamic) The **`is`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) allows you to specify that a standard HTML element should behave like a defined custom built-in element (see [Using custom elements](https://developer.mozilla.org/en-US/docs/Web/API/Web_components/Using_custom_elements) for more details). ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/is))
- `ismap` (Dynamic) This Boolean attribute indicates that the image is part of a [server-side map](https://en.wikipedia.org/wiki/Image_map#Server-side).
- `itemid` (Dynamic) The **`itemid`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) provides microdata in the form of a unique, global identifier of an item. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemid))
- `itemprop` (Dynamic) The **`itemprop`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is used to add properties to an item. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemprop))
- `itemref` (Dynamic) Properties that are not descendants of an element with the [`itemscope`](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemscope) attribute can be associated with an item using the [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) **`itemref`**. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemref))
- `itemscope` (Dynamic) **`itemscope`** is a boolean [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) that defines the scope of associated metadata. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemscope))
- `itemtype` (Dynamic) The [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) **`itemtype`** specifies the URL of the vocabulary that will be used to define `itemprop`'s (item properties) in the data structure. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemtype))
- `lang` (Dynamic) The **`lang`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) helps define the language of an element: the language that non-editable elements are written in, or the language that the editable elements should be written in by the user. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/lang))
- `loading` (Dynamic) Indicates how the browser should load the image.
- `nonce` (Dynamic) The **`nonce`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is a content attribute defining a cryptographic nonce ("number used once") which can be used by [Content Security Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP) to determine whether or not a given fetch will be allowed to proceed for a given element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/nonce))
- `part` (Dynamic) The **`part`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) contains a space-separated list of the part names of the element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/part))
- `popover` (Dynamic) The **`popover`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is used to designate an element as a popover element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/popover))
- `referrerpolicy` (Dynamic) A string indicating which referrer to use when fetching the resource.
- `sizes` (Dynamic) One or more strings separated by commas, indicating a set of source sizes.
- `slot` (Dynamic) The **`slot`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) assigns a slot in a [shadow DOM](https://developer.mozilla.org/en-US/docs/Web/API/Web_components/Using_shadow_DOM) shadow tree to an element: An element with a `slot` attribute is assigned to the slot created by the [`<slot>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/slot) element whose [`name`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/slot#name) attribute's value matches that `slot` attribute's value. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/slot))
- `spellcheck` (Dynamic) The **`spellcheck`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute that defines whether the element may be checked for spelling errors. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/spellcheck))
- `src` (Dynamic) The image [URL](https://developer.mozilla.org/en-US/docs/Glossary/URL).
- `srcset` (Dynamic) One or more strings separated by commas, indicating possible image sources for the [user agent](https://developer.mozilla.org/en-US/docs/Glossary/User_agent) to use.
- `style` (Dynamic) The **`style`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) contains [CSS](https://developer.mozilla.org/en-US/docs/Web/CSS) styling declarations to be applied to the element. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/style))
- `tabindex` (Dynamic) The **`tabindex`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) allows developers to make HTML elements focusable, allow or prevent them from being sequentially focusable (usually with the <kbd>Tab</kbd> key, hence the name) and determine their relative ordering for sequential focus navigation. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/tabindex))
- `title` (Dynamic) The **`title`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) contains text representing advisory information related to the element it belongs to. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/title))
- `translate` (Dynamic) The **`translate`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) is an [enumerated](https://developer.mozilla.org/en-US/docs/Glossary/Enumerated) attribute that is used to specify whether an element's _translatable attribute_ values and its [`Text`](https://developer.mozilla.org/en-US/docs/Web/API/Text) node children should be translated when the page is localized, or whether to leave them unchanged. ([Documentation](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/translate))
- `usemap` (Dynamic) The partial [URL](https://developer.mozilla.org/en-US/docs/Glossary/URL) (starting with `#`) of an [image map](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map) associated with the element.
- `width` (Dynamic) The intrinsic width of the image in pixels. Must be an integer without a unit.

### Read-Only

- `html` (String) The html of the element.
