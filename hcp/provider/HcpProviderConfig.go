// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type HcpProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs#alias HcpProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The OAuth2 Client ID for API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs#client_id HcpProvider#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The OAuth2 Client Secret for API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs#client_secret HcpProvider#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The path to an HCP credential file to use to authenticate the provider to HCP.
	//
	// You can alternatively set the HCP_CRED_FILE environment variable to point at a credential file as well. Using a credential file allows you to authenticate the provider as a service principal via client credentials or dynamically based on Workload Identity Federation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs#credential_file HcpProvider#credential_file}
	CredentialFile *string `field:"optional" json:"credentialFile" yaml:"credentialFile"`
	// The geography in which HCP resources should be created. Default is `us`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs#geography HcpProvider#geography}
	Geography *string `field:"optional" json:"geography" yaml:"geography"`
	// The default project in which resources should be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs#project_id HcpProvider#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// When set to true, the provider will skip checking the HCP status page for service outages or returning warnings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs#skip_status_check HcpProvider#skip_status_check}
	SkipStatusCheck interface{} `field:"optional" json:"skipStatusCheck" yaml:"skipStatusCheck"`
	// workload_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs#workload_identity HcpProvider#workload_identity}
	WorkloadIdentity interface{} `field:"optional" json:"workloadIdentity" yaml:"workloadIdentity"`
}

