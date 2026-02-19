// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsdynamicsecret


type VaultSecretsDynamicSecretAwsAssumeRole struct {
	// AWS IAM role ARN to assume when generating credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_dynamic_secret#iam_role_arn VaultSecretsDynamicSecret#iam_role_arn}
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
}

