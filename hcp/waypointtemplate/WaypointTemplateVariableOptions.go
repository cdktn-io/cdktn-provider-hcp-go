// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointtemplate


type WaypointTemplateVariableOptions struct {
	// Variable name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#name WaypointTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Variable type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#variable_type WaypointTemplate#variable_type}
	VariableType *string `field:"required" json:"variableType" yaml:"variableType"`
	// List of options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#options WaypointTemplate#options}
	Options *[]*string `field:"optional" json:"options" yaml:"options"`
	// Whether the variable is editable by the user creating an application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#user_editable WaypointTemplate#user_editable}
	UserEditable interface{} `field:"optional" json:"userEditable" yaml:"userEditable"`
}

