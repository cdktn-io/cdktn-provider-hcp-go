// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointtemplate


type WaypointTemplateTerraformCloudWorkspaceDetails struct {
	// Name of the Terraform Cloud Project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/waypoint_template#name WaypointTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Terraform Cloud Project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/waypoint_template#terraform_project_id WaypointTemplate#terraform_project_id}
	TerraformProjectId *string `field:"required" json:"terraformProjectId" yaml:"terraformProjectId"`
}

