// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsrotatingsecret


type VaultSecretsRotatingSecretAwsAccessKeys struct {
	// AWS IAM username to rotate the access keys for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/vault_secrets_rotating_secret#iam_username VaultSecretsRotatingSecret#iam_username}
	IamUsername *string `field:"required" json:"iamUsername" yaml:"iamUsername"`
}

