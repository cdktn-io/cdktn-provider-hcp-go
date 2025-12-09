// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointaction


type WaypointActionRequestAgent struct {
	// The name of the group that the operation is in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_action#group WaypointAction#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// The identifying name of the operation in the agent config file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_action#operation_id WaypointAction#operation_id}
	OperationId *string `field:"required" json:"operationId" yaml:"operationId"`
	// An optional action run id. If specified the agent will interact with the actions subsystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_action#action_run_id WaypointAction#action_run_id}
	ActionRunId *string `field:"optional" json:"actionRunId" yaml:"actionRunId"`
	// Arguments to the operation, specified as JSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_action#body WaypointAction#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
}

