// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package packerbucketiambinding

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PackerBucketIamBindingConfig struct {
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
	// The principal to bind to the given role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/packer_bucket_iam_binding#principal_id PackerBucketIamBinding#principal_id}
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// The bucket's resource name in the format packer/project/<project ID>/bucket/<bucket name>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/packer_bucket_iam_binding#resource_name PackerBucketIamBinding#resource_name}
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// The role name to bind to the given principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/packer_bucket_iam_binding#role PackerBucketIamBinding#role}
	Role *string `field:"required" json:"role" yaml:"role"`
}

