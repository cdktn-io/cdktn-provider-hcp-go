// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointaction


type WaypointActionRequest struct {
	// Agent mode allows users to define the agent to use for the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/waypoint_action#agent WaypointAction#agent}
	Agent *WaypointActionRequestAgent `field:"optional" json:"agent" yaml:"agent"`
	// Custom mode allows users to define the HTTP method, the request body, etc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/waypoint_action#custom WaypointAction#custom}
	Custom *WaypointActionRequestCustom `field:"optional" json:"custom" yaml:"custom"`
}

