// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package waypointagentgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WaypointAgentGroupConfig struct {
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
	// The name of the Agent Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_agent_group#name WaypointAgentGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the Agent Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_agent_group#description WaypointAgentGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the Waypoint organization to which the Agent Group belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_agent_group#organization_id WaypointAgentGroup#organization_id}
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
	// The ID of the Waypoint project to which the Agent Group belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_agent_group#project_id WaypointAgentGroup#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

