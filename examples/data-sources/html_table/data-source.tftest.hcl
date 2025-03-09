run "simple" {
  assert {
    condition     = data.html_table.simple.html == "<table><thead><tr><th>head1</th><th>head2</th><th>head3</th></tr></thead><tbody><tr><td>data1</td><td>data2</td><td>data3</td></tr></tbody></table>"
    error_message = ""
  }
}
