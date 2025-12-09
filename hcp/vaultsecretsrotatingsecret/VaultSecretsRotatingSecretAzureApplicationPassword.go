// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsrotatingsecret


type VaultSecretsRotatingSecretAzureApplicationPassword struct {
	// Application client ID to rotate the application password for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_rotating_secret#app_client_id VaultSecretsRotatingSecret#app_client_id}
	AppClientId *string `field:"required" json:"appClientId" yaml:"appClientId"`
	// Application object ID to rotate the application password for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_rotating_secret#app_object_id VaultSecretsRotatingSecret#app_object_id}
	AppObjectId *string `field:"required" json:"appObjectId" yaml:"appObjectId"`
}

