// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.Provider = &ReverseProvider{}
var _ provider.ProviderWithFunctions = &ReverseProvider{}
var _ provider.ProviderWithEphemeralResources = &ReverseProvider{}

// ReverseProvider defines the provider implementation.
type ReverseProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// ReverseProviderModel describes the provider data model.
type ReverseProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *ReverseProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "reverse"
	resp.Version = p.version
}

func (p *ReverseProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Example provider attribute",
				Optional:            true,
			},
		},
	}
}

func (p *ReverseProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data ReverseProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client

	ctx = tflog.SetField(ctx, "endpoint", data.Endpoint.ValueString())
	tflog.Debug(ctx, "Successfully configured the provider")
}

func (p *ReverseProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewStringResource,
	}
}

func (p *ReverseProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{
		NewSecretEphemeralResource,
	}
}

func (p *ReverseProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *ReverseProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ReverseProvider{
			version: version,
		}
	}
}
