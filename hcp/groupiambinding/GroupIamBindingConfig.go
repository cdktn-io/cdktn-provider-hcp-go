// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groupiambinding

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GroupIamBindingConfig struct {
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
	// The group's resource name in format `iam/organization/<organization_id>/group/<group_name>`. The shortened `<group_name>` version can be used for input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/group_iam_binding#name GroupIamBinding#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The principal to bind to the given role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/group_iam_binding#principal_id GroupIamBinding#principal_id}
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// The role name to bind to the given principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/group_iam_binding#role GroupIamBinding#role}
	Role *string `field:"required" json:"role" yaml:"role"`
}

