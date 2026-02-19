// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultradarresourceiambinding

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VaultRadarResourceIamBindingConfig struct {
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
	// The principal to bind to the given role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_resource_iam_binding#principal_id VaultRadarResourceIamBinding#principal_id}
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// The HCP resource name associated with the Radar resource.
	//
	// This is the name of the resource in the format `vault-radar/project/<project_id>/scan-target/<scan_target_id>`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_resource_iam_binding#resource_name VaultRadarResourceIamBinding#resource_name}
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// The role name to bind to the given principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_resource_iam_binding#role VaultRadarResourceIamBinding#role}
	Role *string `field:"required" json:"role" yaml:"role"`
}

