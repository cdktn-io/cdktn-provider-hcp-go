// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsrotatingsecret


type VaultSecretsRotatingSecretGcpServiceAccountKey struct {
	// GCP service account email to impersonate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_rotating_secret#service_account_email VaultSecretsRotatingSecret#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

