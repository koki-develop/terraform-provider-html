data "html_table" "simple" {
  children = [
    data.html_thead.simple.html,
    data.html_tbody.simple.html,
  ]
}
# => <table>
#      <thead>
#        <tr>
#          <th>head1</th>
#          <th>head2</th>
#          <th>head3</th>
#        </tr>
#      </thead>
#      <tbody>
#        <tr>
#          <td>data1</td>
#          <td>data2</td>
#          <td>data3</td>
#        </tr>
#      </tbody>
#    </table>

data "html_thead" "simple" {
  children = [data.html_tr.simple_header.html]
}

data "html_tr" "simple_header" {
  children = [for k, v in data.html_th.simple : v.html]
}

data "html_th" "simple" {
  for_each = toset(["head1", "head2", "head3"])
  children = [each.value]
}

data "html_tbody" "simple" {
  children = [data.html_tr.simple_body.html]
}

data "html_tr" "simple_body" {
  children = [for k, v in data.html_td.simple : v.html]
}

data "html_td" "simple" {
  for_each = toset(["data1", "data2", "data3"])
  children = [each.value]
}
