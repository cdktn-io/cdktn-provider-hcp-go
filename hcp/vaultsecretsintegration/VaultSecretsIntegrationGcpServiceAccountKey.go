// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegration


type VaultSecretsIntegrationGcpServiceAccountKey struct {
	// JSON or base64 encoded service account key received from GCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_secrets_integration#credentials VaultSecretsIntegration#credentials}
	Credentials *string `field:"required" json:"credentials" yaml:"credentials"`
}

