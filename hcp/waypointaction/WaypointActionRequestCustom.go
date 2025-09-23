// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointaction


type WaypointActionRequestCustom struct {
	// The HTTP method to use for the request. Must be one of: 'GET', 'POST', 'PUT', 'DELETE', or 'PATCH'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/waypoint_action#method WaypointAction#method}
	Method *string `field:"required" json:"method" yaml:"method"`
	// The body to be submitted with the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/waypoint_action#body WaypointAction#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Key value headers to send with the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/waypoint_action#headers WaypointAction#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// The full URL this request should make when invoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/waypoint_action#url WaypointAction#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

