resource "html_a" "simple" {
  children = ["Click me!"]
  href     = "https://example.com"
}
# => <a href="https://example.com">Click me!</a>
