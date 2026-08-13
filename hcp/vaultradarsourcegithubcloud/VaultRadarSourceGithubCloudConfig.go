// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultradarsourcegithubcloud

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VaultRadarSourceGithubCloudConfig struct {
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
	// GitHub organization Vault Radar will monitor. Example: type "octocat" for the org https://github.com/octocat.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.114.0/docs/resources/vault_radar_source_github_cloud#github_organization VaultRadarSourceGithubCloud#github_organization}
	GithubOrganization *string `field:"required" json:"githubOrganization" yaml:"githubOrganization"`
	// The detector type to use for monitoring this source.
	//
	// Valid values are 'hcp' (managed by HCP) or 'agent' (self-hosted agent). Defaults to 'hcp'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.114.0/docs/resources/vault_radar_source_github_cloud#detector_type VaultRadarSourceGithubCloud#detector_type}
	DetectorType *string `field:"optional" json:"detectorType" yaml:"detectorType"`
	// The ID of the HCP project where Vault Radar is located.
	//
	// If not specified, the project specified in the HCP Provider config block will be used, if configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.114.0/docs/resources/vault_radar_source_github_cloud#project_id VaultRadarSourceGithubCloud#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// GitHub personal access token.
	//
	// Required when detector_type is 'hcp' or not specified (defaults to 'hcp'). Cannot be used when detector_type is 'agent'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.114.0/docs/resources/vault_radar_source_github_cloud#token VaultRadarSourceGithubCloud#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Environment variable name containing the GitHub personal access token.
	//
	// When detector_type is 'agent', this is required. When detector_type is 'hcp' or not specified (defaults to 'hcp'), this is optional and can be set to enable optional secret copying via the Vault Radar Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.114.0/docs/resources/vault_radar_source_github_cloud#token_env_var VaultRadarSourceGithubCloud#token_env_var}
	TokenEnvVar *string `field:"optional" json:"tokenEnvVar" yaml:"tokenEnvVar"`
}

