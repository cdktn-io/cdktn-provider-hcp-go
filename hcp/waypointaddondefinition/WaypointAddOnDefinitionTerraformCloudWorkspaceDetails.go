// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package waypointaddondefinition


type WaypointAddOnDefinitionTerraformCloudWorkspaceDetails struct {
	// Name of the Terraform Cloud Project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_add_on_definition#name WaypointAddOnDefinition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Terraform Cloud Project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_add_on_definition#terraform_project_id WaypointAddOnDefinition#terraform_project_id}
	TerraformProjectId *string `field:"required" json:"terraformProjectId" yaml:"terraformProjectId"`
}

