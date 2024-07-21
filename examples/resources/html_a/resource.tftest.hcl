run "simple" {
  assert {
    condition     = html_a.simple.html == "<a href=\"https://example.com\">Click me!</a>"
    error_message = ""
  }
}
