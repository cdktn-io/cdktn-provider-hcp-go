// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahcpwaypointaction


type DataHcpWaypointActionRequest struct {
	// Agent mode allows users to define the agent to use for the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/waypoint_action#agent DataHcpWaypointAction#agent}
	Agent *DataHcpWaypointActionRequestAgent `field:"optional" json:"agent" yaml:"agent"`
}

