// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointaddon

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WaypointAddOnConfig struct {
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
	// The ID of the Application that this Add-on is created for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/waypoint_add_on#application_id WaypointAddOn#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The ID of the Add-on Definition that this Add-on is created from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/waypoint_add_on#definition_id WaypointAddOn#definition_id}
	DefinitionId *string `field:"required" json:"definitionId" yaml:"definitionId"`
	// The name of the Add-on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/waypoint_add_on#name WaypointAddOn#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Input variables set for the add-on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/waypoint_add_on#add_on_input_variables WaypointAddOn#add_on_input_variables}
	AddOnInputVariables interface{} `field:"optional" json:"addOnInputVariables" yaml:"addOnInputVariables"`
	// The ID of the HCP project where the Waypoint AddOn is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/waypoint_add_on#project_id WaypointAddOn#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

