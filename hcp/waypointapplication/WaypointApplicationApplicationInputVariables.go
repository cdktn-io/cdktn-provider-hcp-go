// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointapplication


type WaypointApplicationApplicationInputVariables struct {
	// Variable name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/waypoint_application#name WaypointApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Variable value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/waypoint_application#value WaypointApplication#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Variable type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/waypoint_application#variable_type WaypointApplication#variable_type}
	VariableType *string `field:"required" json:"variableType" yaml:"variableType"`
}

