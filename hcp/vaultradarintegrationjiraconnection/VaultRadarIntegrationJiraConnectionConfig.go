// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultradarintegrationjiraconnection

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VaultRadarIntegrationJiraConnectionConfig struct {
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
	// The Jira base URL. Example: https://acme.atlassian.net.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_integration_jira_connection#base_url VaultRadarIntegrationJiraConnection#base_url}
	BaseUrl *string `field:"required" json:"baseUrl" yaml:"baseUrl"`
	// Jira user's email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_integration_jira_connection#email VaultRadarIntegrationJiraConnection#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Name of connection. Name must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_integration_jira_connection#name VaultRadarIntegrationJiraConnection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A Jira API token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_integration_jira_connection#token VaultRadarIntegrationJiraConnection#token}
	Token *string `field:"required" json:"token" yaml:"token"`
	// The ID of the HCP project where Vault Radar is located.
	//
	// If not specified, the project specified in the HCP Provider config block will be used, if configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_integration_jira_connection#project_id VaultRadarIntegrationJiraConnection#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

