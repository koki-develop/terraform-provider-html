package datasources

import "github.com/hashicorp/terraform-plugin-framework/resource"

var Resources []func() resource.Resource

func register(f func() resource.Resource) {
	Resources = append(Resources, f)
}
