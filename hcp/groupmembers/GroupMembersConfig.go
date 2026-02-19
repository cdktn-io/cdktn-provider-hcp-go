// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groupmembers

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GroupMembersConfig struct {
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
	// The group's resource name in the format `iam/organization/<organization_id>/group/<name>`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/group_members#group GroupMembers#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// A list of user principal IDs to add to the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/group_members#members GroupMembers#members}
	Members *[]*string `field:"required" json:"members" yaml:"members"`
}

