// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegration


type VaultSecretsIntegrationTwilioStaticCredentials struct {
	// Account SID for the target Twilio account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#account_sid VaultSecretsIntegration#account_sid}
	AccountSid *string `field:"required" json:"accountSid" yaml:"accountSid"`
	// Api key secret used with the api key SID to authenticate against the target Twilio account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#api_key_secret VaultSecretsIntegration#api_key_secret}
	ApiKeySecret *string `field:"required" json:"apiKeySecret" yaml:"apiKeySecret"`
	// Api key SID to authenticate against the target Twilio account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#api_key_sid VaultSecretsIntegration#api_key_sid}
	ApiKeySid *string `field:"required" json:"apiKeySid" yaml:"apiKeySid"`
}

