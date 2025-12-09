// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegrationgcp


type VaultSecretsIntegrationGcpServiceAccountKeyA struct {
	// JSON or base64 encoded service account key received from GCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration_gcp#credentials VaultSecretsIntegrationGcp#credentials}
	Credentials *string `field:"required" json:"credentials" yaml:"credentials"`
}

