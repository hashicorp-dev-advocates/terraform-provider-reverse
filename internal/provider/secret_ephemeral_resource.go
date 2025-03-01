// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp-sandbox/go-reverse/reverse"
	"github.com/hashicorp-sandbox/go-reverse/secret"

	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ ephemeral.EphemeralResource = &SecretEphemeralResource{}

func NewSecretEphemeralResource() ephemeral.EphemeralResource {
	return &SecretEphemeralResource{}
}

// SecretEphemeralResource defines the ephemeral resource implementation.
type SecretEphemeralResource struct {
	// client *http.Client // If applicable, a client can be initialized here.
}

// SecretEphemeralResourceModel describes the ephemeral resource data model.
type SecretEphemeralResourceModel struct {
	SecretID types.String `tfsdk:"secret_id"`
	Value    types.String `tfsdk:"value"`
}

func (r *SecretEphemeralResource) Metadata(_ context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_secret"
}

func (r *SecretEphemeralResource) Schema(ctx context.Context, _ ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example ephemeral resource",

		Attributes: map[string]schema.Attribute{
			"secret_id": schema.StringAttribute{
				MarkdownDescription: "Identifier used to retrieve the secret value",
				Required:            true, // Ephemeral resources expect their dependencies to already exist.
			},
			"value": schema.StringAttribute{
				Computed: true,
				// Sensitive:           true, // If applicable, mark the attribute as sensitive.
				MarkdownDescription: "Reversed secret value",
			},
		},
	}
}

func (r *SecretEphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var data SecretEphemeralResourceModel

	// Read Terraform config data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	secret := secret.GetByID("some/id")
	data.Value = types.StringValue(reverse.String(secret))

	// Save data into ephemeral result data
	resp.Diagnostics.Append(resp.Result.Set(ctx, &data)...)
}
