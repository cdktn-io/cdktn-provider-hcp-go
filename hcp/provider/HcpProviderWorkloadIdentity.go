// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type HcpProviderWorkloadIdentity struct {
	// The resource_name of the Workload Identity Provider to exchange the token with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs#resource_name HcpProvider#resource_name}
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// The JWT token retrieved from an OpenID Connect (OIDC) or OAuth2 provider.
	//
	// At least one of `token_file` or `token` must be set, if both are set then `token` takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs#token HcpProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// The path to a file containing a JWT token retrieved from an OpenID Connect (OIDC) or OAuth2 provider.
	//
	// At least one of `token_file` or `token` must be set, if both are set then `token` takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs#token_file HcpProvider#token_file}
	TokenFile *string `field:"optional" json:"tokenFile" yaml:"tokenFile"`
}

