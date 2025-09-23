// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointapplication


type WaypointApplicationTemplateInputVariables struct {
	// Variable name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/waypoint_application#name WaypointApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Variable value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/waypoint_application#value WaypointApplication#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Variable type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/waypoint_application#variable_type WaypointApplication#variable_type}
	VariableType *string `field:"optional" json:"variableType" yaml:"variableType"`
}

