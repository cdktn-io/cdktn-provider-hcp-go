// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultradarsourcegithubenterprise

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VaultRadarSourceGithubEnterpriseConfig struct {
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
	// Fully qualified domain name of the server. (Example: myserver.acme.com).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_source_github_enterprise#domain_name VaultRadarSourceGithubEnterprise#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// GitHub organization Vault Radar will monitor. Example: "octocat" for the org https://yourcodeserver.com/octocat.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_source_github_enterprise#github_organization VaultRadarSourceGithubEnterprise#github_organization}
	GithubOrganization *string `field:"required" json:"githubOrganization" yaml:"githubOrganization"`
	// GitHub personal access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_source_github_enterprise#token VaultRadarSourceGithubEnterprise#token}
	Token *string `field:"required" json:"token" yaml:"token"`
	// The detector type which will monitor this resource. The default is HCP if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_source_github_enterprise#detector_type VaultRadarSourceGithubEnterprise#detector_type}
	DetectorType *string `field:"optional" json:"detectorType" yaml:"detectorType"`
	// The ID of the HCP project where Vault Radar is located.
	//
	// If not specified, the project specified in the HCP Provider config block will be used, if configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_source_github_enterprise#project_id VaultRadarSourceGithubEnterprise#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

