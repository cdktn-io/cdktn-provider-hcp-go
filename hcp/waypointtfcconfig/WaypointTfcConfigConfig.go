// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointtfcconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WaypointTfcConfigConfig struct {
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
	// The Terraform Cloud Organization with which the token is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_tfc_config#tfc_org_name WaypointTfcConfig#tfc_org_name}
	TfcOrgName *string `field:"required" json:"tfcOrgName" yaml:"tfcOrgName"`
	// Terraform Cloud team token. The token must include permissions to manage workspaces and applications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_tfc_config#token WaypointTfcConfig#token}
	Token *string `field:"required" json:"token" yaml:"token"`
	// Waypoint Project ID to associate with the TFC config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_tfc_config#project_id WaypointTfcConfig#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

