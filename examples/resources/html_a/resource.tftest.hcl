run "simple" {
  assert {
    condition     = html_a.simple.html == "<a href=\"https://example.com\">Click me!</a>"
    error_message = ""
  }
}

run "download" {
  assert {
    condition     = html_a.download.html == "<a download href=\"https://example.com/file\">Download</a>"
    error_message = ""
  }
}

run "nolink" {
  assert {
    condition     = html_a.nolink.html == "<a>No link</a>"
    error_message = ""
  }
}
