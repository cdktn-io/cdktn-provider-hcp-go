// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package waypointaction

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WaypointActionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the Action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_action#name WaypointAction#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The kind of HTTP request this should trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_action#request WaypointAction#request}
	Request *WaypointActionRequest `field:"required" json:"request" yaml:"request"`
	// A description of the Action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_action#description WaypointAction#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the HCP project where the Action is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_action#project_id WaypointAction#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

