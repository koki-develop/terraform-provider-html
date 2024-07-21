---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "html_a Resource - terraform-provider-html"
subcategory: ""
description: |-
  The <a> HTML https://developer.mozilla.org/en-US/docs/Web/HTML element (or anchor element), with its href attribute https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a#href, creates a hyperlink to web pages, files, email addresses, locations in the same page, or anything else a URL can address.
---

# html_a (Resource)

The **`<a>`** [HTML](https://developer.mozilla.org/en-US/docs/Web/HTML) element (or _anchor_ element), with [its `href` attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a#href), creates a hyperlink to web pages, files, email addresses, locations in the same page, or anything else a URL can address.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `accesskey` (Dynamic) The **`accesskey`** [global attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes) provides a hint for generating a keyboard shortcut for the current element.
- `attributionsrc` (Dynamic) Specifies that you want the browser to send an [`Attribution-Reporting-Eligible`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Attribution-Reporting-Eligible) header.
- `children` (List of String) The children of the element.
- `download` (Dynamic) Causes the browser to treat the linked URL as a download.
- `href` (Dynamic) The URL that the hyperlink points to.
- `hreflang` (Dynamic) Hints at the human language of the linked URL.
- `ping` (Dynamic) A space-separated list of URLs.
- `referrerpolicy` (Dynamic) How much of the [referrer](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referer) to send when following the link.
- `rel` (Dynamic) The relationship of the linked URL as space-separated link types.
- `target` (Dynamic) Where to display the linked URL, as the name for a _browsing context_ (a tab, window, or [`<iframe>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe)).
- `type` (Dynamic) Hints at the linked URL's format with a [MIME type](https://developer.mozilla.org/en-US/docs/Glossary/MIME_type).

### Read-Only

- `html` (String) The html of the element.