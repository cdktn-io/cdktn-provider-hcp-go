// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointaddon


type WaypointAddOnAddOnDefinitionInputVariables struct {
	// Variable name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/waypoint_add_on#name WaypointAddOn#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Variable value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/waypoint_add_on#value WaypointAddOn#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Variable type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/waypoint_add_on#variable_type WaypointAddOn#variable_type}
	VariableType *string `field:"optional" json:"variableType" yaml:"variableType"`
}

