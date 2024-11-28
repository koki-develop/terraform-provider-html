package datasources

import "github.com/hashicorp/terraform-plugin-framework/datasource"

var DataSources []func() datasource.DataSource

func register(f func() datasource.DataSource) {
	DataSources = append(DataSources, f)
}
