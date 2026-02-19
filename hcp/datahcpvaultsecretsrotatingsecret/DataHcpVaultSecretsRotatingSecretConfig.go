// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datahcpvaultsecretsrotatingsecret

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataHcpVaultSecretsRotatingSecretConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the Vault Secrets application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/vault_secrets_rotating_secret#app_name DataHcpVaultSecretsRotatingSecret#app_name}
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// The name of the Vault Secrets secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/vault_secrets_rotating_secret#secret_name DataHcpVaultSecretsRotatingSecret#secret_name}
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

