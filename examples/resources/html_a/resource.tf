resource "html_a" "main" {
  children = ["Click me!"]
  href     = "https://example.com"
}
# => <a href="https://example.com">Click me!</a>
