// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultradarresourceiampolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VaultRadarResourceIamPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The policy to apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_radar_resource_iam_policy#policy_data VaultRadarResourceIamPolicy#policy_data}
	PolicyData *string `field:"required" json:"policyData" yaml:"policyData"`
	// The HCP resource name associated with the Radar resource.
	//
	// This is the name of the resource in the format `vault-radar/project/<project_id>/scan-target/<scan_target_id>`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_radar_resource_iam_policy#resource_name VaultRadarResourceIamPolicy#resource_name}
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
}

