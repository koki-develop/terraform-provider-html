package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
	_ provider.Provider = (*htmlProvider)(nil)
)

type htmlProvider struct {
	version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &htmlProvider{
			version: version,
		}
	}
}

func (p *htmlProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "html"
	resp.Version = p.version
}

func (p *htmlProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The Next Generation HTML Builder.",
	}
}

func (p *htmlProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
}

func (p *htmlProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *htmlProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
