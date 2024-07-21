resource "html_a" "simple" {
  children = ["Click me!"]
  href     = "https://example.com"
}
# => <a href="https://example.com">Click me!</a>

resource "html_a" "download" {
  children = ["Download"]
  download = true
  href     = "https://example.com/file"
}
# => <a download href="https://example.com/file">Download</a>

resource "html_a" "nolink" {
  children = ["No link"]
}
# => <a>No link</a>
